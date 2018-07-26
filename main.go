package main

import (
	"os"
	"github.com/docker/docker/pkg/term"
	"github.com/continuul/random-names/command"
	"github.com/continuul/random-names/pkg/stream"
	"github.com/spf13/cobra"
	"fmt"
	"github.com/continuul/random-names/command/commands"
	"math/rand"
	"time"
)

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func newRootCommand(cli *command.CliInstance) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "random-names",
		Short:   "random-names provides Docker-style random names",
		Long:    `A name generator for Docker style names.`,
		Run:     runHelp,
		Version: fmt.Sprintf("%s, build %s, describe %s", command.Version, command.GitCommit, command.GitDescribe),
	}
	flags := cmd.PersistentFlags()
	flags.BoolP("version", "v", false, "Print version information and quit")
	commands.AddCommands(cmd, cli)
	return cmd
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	args := os.Args[1:]
	for _, arg := range args {
		if arg == "--" {
			break
		}

		if arg == "-v" || arg == "--version" {
			args = []string{"version"}
			break
		}
	}

	stdin, stdout, stderr := term.StdStreams()
	cli := command.NewCliInstance(stream.NewInStream(stdin), stream.NewOutStream(stdout), stderr)
	cmd := newRootCommand(cli)

	if err := cmd.Execute(); err != nil {
		if sterr, ok := err.(command.StatusError); ok {
			if sterr.Status != "" {
				fmt.Fprintln(stderr, sterr.Status)
			}
			// StatusError should only be used for errors, and all errors should
			// have a non-zero exit status, so never exit with 0
			if sterr.StatusCode == 0 {
				os.Exit(1)
			}
			os.Exit(sterr.StatusCode)
		}
		fmt.Fprintln(stderr, err)
		os.Exit(1)
	}
}
