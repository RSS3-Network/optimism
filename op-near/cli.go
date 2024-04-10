package opnear

import (
	"errors"

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

var (
	ErrMissingDaAccount  = errors.New("near da account cannot be empty")
	ErrMissingDaContract = errors.New("near da contract cannot be empty")
	ErrMissingDaKey      = errors.New("near da key cannot be empty")
	ErrInvalidDaNetwork  = errors.New("near da network invalid")
	ErrDaNamespaceIdZero = errors.New("near da namespace id cannot be 0")
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
			EnvVars:  opservice.PrefixEnvVar(envPrefix, "NEAR_DA_KEY"),
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
	DaAccount     string
	DaContract    string
	DaKey         string
	DaNetwork     string
	DaNamespaceId uint32
}

func (c CLIConfig) Check() error {
	if c.DaAccount == "" {
		return ErrMissingDaAccount
	}
	if c.DaContract == "" {
		return ErrMissingDaContract
	}
	if c.DaKey == "" {
		return ErrMissingDaKey
	}
	if c.DaNetwork != "Mainnet" && c.DaNetwork != "Testnet" {
		return ErrInvalidDaNetwork
	}
	if c.DaNamespaceId == 0 {
		return ErrDaNamespaceIdZero
	}

	return nil
}

func ReadCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		DaAccount:     ctx.String(NearDaAccountFlagName),
		DaContract:    ctx.String(NearDaContractFlagName),
		DaKey:         ctx.String(NearDaKeyFlagName),
		DaNetwork:     ctx.String(NearDaNetworkFlagName),
		DaNamespaceId: uint32(ctx.Uint64(NearDaNamespaceIdFlagName)),
	}
}
