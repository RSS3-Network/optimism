package opnear

import (
	"fmt"
	near "github.com/near/rollup-data-availability/gopkg/da-rpc"
)

type DAClient struct {
	Client near.Config
}

func NewDAClient(NearDaAccount, NearDaContract, NearDaKey, NearDaNetwork string, NearDaNamespaceId uint32) (*DAClient, error) {
	nearDaConfig, err := near.NewConfig(NearDaAccount, NearDaContract, NearDaKey, NearDaNetwork, NearDaNamespaceId)
	if err != nil {
		return nil, fmt.Errorf("failed to create Near DA config: %w", err)
	}
	return &DAClient{
		Client: *nearDaConfig,
	}, nil
}
