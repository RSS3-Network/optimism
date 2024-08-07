package opnear

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-service/retry"
	near "github.com/near/rollup-data-availability/gopkg/da-rpc"
)

const (
	defaultSubmitTimeout = 90 * time.Second
	defaultGetTimeout    = 90 * time.Second

	defaultSubmitAttempts = 200
	defaultGetAttempts    = 20
)

type DAClient struct {
	Client near.Config
}

func NewDAClient(accountN, contractN, keyN, networkN string, nameSpace uint32) (*DAClient, error) {
	client, err := near.NewConfig(accountN, contractN, keyN, networkN, nameSpace)
	if err != nil {
		return nil, fmt.Errorf("failed to create near da client: %w", err)
	}
	return &DAClient{
		Client: *client,
	}, nil
}

func (c *DAClient) FreeDAClient() {
	log.Info("DAClient: free NEAR client")
	c.Client.FreeClient()
}

func (c *DAClient) Submit(data []byte) ([]byte, error) {
	log.Info("DAClient: start submitting blob with retry")
	startTime := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), defaultSubmitTimeout)
	defer cancel()

	bOff := retry.ExponentialDA()
	frameRef, err := retry.Do(ctx, defaultSubmitAttempts, bOff, func() ([]byte, error) {
		result, er := c.Client.ForceSubmit(data)
		if er != nil {
			log.Warn("DAClient: submit blob to near da", "er", er)
			return nil, er
		}

		return result, nil
	})
	if err != nil {
		log.Error("DAClient: failed to submit blob to near da", "err", err)
		return nil, err
	}

	log.Info("DAClient: end submitting blob with retry", "id", hex.EncodeToString(frameRef), "elapsed time", time.Since(startTime))
	return frameRef, nil
}

func (c *DAClient) Get(frameRefBytes []byte, txIndex uint32) ([]byte, error) {
	log.Info("DAClient: start getting blob with retry")
	startTime := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), defaultGetTimeout)
	defer cancel()

	bOff := retry.ExponentialDA()
	blobData, err := retry.Do(ctx, defaultGetAttempts, bOff, func() ([]byte, error) {
		result, er := c.Client.Get(frameRefBytes, txIndex)
		if er != nil {
			log.Warn("DAClient: get blob from near da", "id", hex.EncodeToString(frameRefBytes), "er", er)
			return nil, er
		}

		return result, nil
	})
	if err != nil {
		log.Error("DAClient:  failed to get blob from near da", "id", hex.EncodeToString(frameRefBytes), "err", err)
		return nil, err
	}

	log.Info("DAClient: end getting blob with retry", "id", hex.EncodeToString(frameRefBytes), "data size", len(blobData), "elapsed time", time.Since(startTime))
	return blobData, nil
}
