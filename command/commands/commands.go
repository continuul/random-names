package commands

import (
	"github.com/spf13/cobra"
	"github.com/continuul/random-names/command/generate"
	"github.com/continuul/random-names/command"
	"github.com/continuul/random-names/command/server"
)

func AddCommands(cmd *cobra.Command, cli command.Cli) {
	cmd.AddCommand(
		generate.New(cli),
	)
	cmd.AddCommand(
		server.New(cli),
	)
}
