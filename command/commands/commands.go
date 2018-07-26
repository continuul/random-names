package commands

import (
	"github.com/spf13/cobra"
	"github.com/continuul/random-names/command/generate"
	"github.com/continuul/random-names/command"
)

func AddCommands(cmd *cobra.Command, cli command.Cli) {
	cmd.AddCommand(
		generate.New(cli),
	)
}
