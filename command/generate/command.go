package generate

import (
	"time"
	"github.com/continuul/random-names/pkg/namesgenerator"
	"math/rand"
	"github.com/spf13/cobra"
	"fmt"
	"github.com/continuul/random-names/command"
)

func New(cli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generates a Docker-style random name",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runThis(cli)
		},
	}
	return cmd
}

func runThis(cli command.Cli) error {
	rand.Seed(time.Now().UnixNano())
	fmt.Fprintln(cli.Out(), namesgenerator.GetRandomName(0))
	return nil
}
