package commands

import (
	"github.com/continuul/random-names/command"
	"github.com/continuul/random-names/command/generate"
	"github.com/continuul/random-names/command/server"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command, cli command.Cli) {
	cmd.AddCommand(
		generate.New(cli),
	)
	cmd.AddCommand(
		server.New(cli),
	)
}
