package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

const flagEnabled = "enabled"

func addDefaultRequiredFlags(cmd *cobra.Command) error {
	cmd.Flags().String(flagEnabled, "", "Flag to invalidate command action")

	if err := cmd.MarkFlagRequired(flagEnabled); err != nil {
		return fmt.Errorf("error while marking flag as required: %w", err)
	}

	return nil
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}

func Execute(w io.Writer) error {
	sampleCmd := &cobra.Command{
		Use:   "sample",
		Short: "Sample command for testing",
		Long:  "Sample command for testing cli apps using golang + cobra",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	sampleCmd.AddCommand(versionCmd(w))
	sampleCmd.AddCommand(caseCmd(w))

	return sampleCmd.Execute()
}

func main() {
	out := os.Stdout

	if err := Execute(out); err != nil {
		_, _ = fmt.Fprintln(out, err)
		os.Exit(1)
	}
}
