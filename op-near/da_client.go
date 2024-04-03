package opnear

import (
	"fmt"

	near "github.com/near/rollup-data-availability/gopkg/da-rpc"
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

func FreeDAClient(c *DAClient) {
	c.Client.FreeClient()
}
