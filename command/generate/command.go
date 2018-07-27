package generate

import (
	"fmt"
	"github.com/continuul/random-names/command"
	"github.com/continuul/random-names/pkg/namesgenerator"
	"github.com/spf13/cobra"
	"math/rand"
	"time"
)

// New creates a Cobra CLI command.
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
