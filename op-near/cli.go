package opnear

import (
	"github.com/urfave/cli/v2"

	opservice "github.com/ethereum-optimism/optimism/op-service"
)

const (
	NearDaAccountFlagName     = "near-da-account"
	NearDaContractFlagName    = "near-da-contract"
	NearDaKeyFlagName         = "near-da-key"
	NearDaNetworkFlagName     = "near-da-network"
	NearDaNamespaceIdFlagName = "near-da-namespace-id"
)

func CLIFlags(envPrefix string, category string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     NearDaAccountFlagName,
			Usage:    "Near DA Signer",
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "NEAR_DA_ACCOUNT"),
			Category: category,
		},
		&cli.StringFlag{
			Name:     NearDaContractFlagName,
			Usage:    "Near DA Contract",
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "NEAR_DA_CONTRACT"),
			Category: category,
		},
		&cli.StringFlag{
			Name:     NearDaKeyFlagName,
			Usage:    "Key for sending messages to the Near DA node",
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "NEAR_DA_ACCOUNT"),
			Category: category,
		},
		&cli.StringFlag{
			Name:     NearDaNetworkFlagName,
			Usage:    "Network for Near DA node (Testnet or Mainnet)",
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "NEAR_DA_NETWORK"),
			Category: category,
		},
		&cli.StringFlag{
			Name:     NearDaNamespaceIdFlagName,
			Usage:    "Namespace ID for Near DA node",
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "NEAR_DA_NAMESPACE_ID"),
			Category: category,
		},
	}
}

type CLIConfig struct {
	NearDaAccount     string
	NearDaContract    string
	NearDaKey         string
	NearDaNetwork     string
	NearDaNamespaceId uint32
}

func (c CLIConfig) Check() error {
	return nil
}

func ReadCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		NearDaAccount:     ctx.String(NearDaAccountFlagName),
		NearDaContract:    ctx.String(NearDaContractFlagName),
		NearDaKey:         ctx.String(NearDaKeyFlagName),
		NearDaNetwork:     ctx.String(NearDaNetworkFlagName),
		NearDaNamespaceId: uint32(ctx.Uint64(NearDaNamespaceIdFlagName)),
	}
}
