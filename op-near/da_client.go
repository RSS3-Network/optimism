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
	defaultGetTimeout = 60 * time.Second

	defaultGetAttempts = 5
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
	c.Client.FreeClient()
}

func (c *DAClient) Submit(data []byte) ([]byte, error) {
	return c.Client.ForceSubmit(data)
}

func (c *DAClient) Get(frameRefBytes []byte, txIndex uint32) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultGetTimeout)
	defer cancel()

	bOff := retry.Exponential()
	blobData, err := retry.Do(ctx, defaultGetAttempts, bOff, func() ([]byte, error) {
		data, err := c.Client.Get(frameRefBytes, txIndex)
		if err != nil {
			log.Warn("get blob from near", "id", hex.EncodeToString(frameRefBytes), "err", err)
			return nil, err
		}

		return data, nil
	})
	if err != nil {
		log.Error("failed to get data from near", "id", hex.EncodeToString(frameRefBytes), "err", err)
		return nil, err
	}
	return blobData, nil
}
