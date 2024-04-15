package driver

import (
	opnear "github.com/ethereum-optimism/optimism/op-near"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
)

func SetDAClient(cfg opnear.CLIConfig) error {
	client, err := opnear.NewDAClient(cfg.DaAccount, cfg.DaContract, cfg.DaKey, cfg.DaNetwork, cfg.DaNamespaceId)
	if err != nil {
		return err
	}
	return derive.SetDAClient(client)
}

func FreeDAClient() {
	derive.FreeDAClient()
}
