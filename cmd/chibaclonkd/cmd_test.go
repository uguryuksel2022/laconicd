package main_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"

	"github.com/tharsis/ethermint/app"
	chibaclonkd "github.com/tharsis/ethermint/cmd/chibaclonkd"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := chibaclonkd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",          // Test the init cmd
		"etherminttest", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, "ethermint_9000-1"),
	})

	err := svrcmd.Execute(rootCmd, chibaclonkd.EnvPrefix, app.DefaultNodeHome)
	require.NoError(t, err)
}
