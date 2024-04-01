package driver

import (
	opnear "github.com/ethereum-optimism/optimism/op-near"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
)

func SetDAClient(cfg opnear.CLIConfig) error {
	client, err := opnear.NewDAClient(cfg.NearDaAccount, cfg.NearDaContract, cfg.NearDaKey, cfg.NearDaNetwork, cfg.NearDaNamespaceId)
	if err != nil {
		return err
	}
	return derive.SetDAClient(client)
}
