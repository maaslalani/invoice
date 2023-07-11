package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "invoice",
	Short: "Invoice generates invoices from the command line.",
	Long:  `Invoice generates invoices from the command line.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
