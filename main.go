package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "gateway",
		Short:        "Bem-vindo",
		SilenceUsage: true,
	}

	cmd.AddCommand(printTimeCmd())
	cmd.AddCommand(printHelloCmd())

	cmd.AddCommand(mainCmd())
	cmd.AddCommand(jardimCmd())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
