package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

func caseCmd(out io.Writer) *cobra.Command {
	caseCmd := &cobra.Command{
		Use:   "case",
		Short: "Modifies input case (upper/lower)",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	caseCmd.AddCommand(uppercaseCmd(out))
	caseCmd.AddCommand(lowercaseCmd(out))

	return caseCmd
}

func uppercaseCmd(out io.Writer) *cobra.Command {
	upper := &cobra.Command{
		Use:   "upper",
		Short: "Shows arg in uppercase",
		Run: func(cmd *cobra.Command, args []string) {
			enabled, err := cmd.Flags().GetString(flagEnabled)
			if err != nil {
				_, _ = fmt.Fprintln(out, err)
				os.Exit(1)
			}

			fmt.Println("enabled", enabled)

			_, _ = fmt.Fprintf(out, "Result: %s\n", strings.ToUpper(args[0]))
		},
	}

	if err := addDefaultRequiredFlags(upper); err != nil {
		fmt.Println(fmt.Errorf("error while adding required flags for upper: %w", err))
		os.Exit(1)
	}

	return upper
}

func lowercaseCmd(out io.Writer) *cobra.Command {
	lower := &cobra.Command{
		Use:   "lower",
		Short: "Shows arg in lowercase",
		Run: func(cmd *cobra.Command, args []string) {
			enabled, err := cmd.Flags().GetString(flagEnabled)
			if err != nil {
				_, _ = fmt.Fprintln(out, err)
				os.Exit(1)
			}

			fmt.Println("enabled", enabled)

			_, _ = fmt.Fprintf(out, "Result: %s\n", strings.ToUpper(args[0]))
		},
	}

	if err := addDefaultRequiredFlags(lower); err != nil {
		fmt.Println(fmt.Errorf("error while adding required flags for lower: %w", err))
		os.Exit(1)
	}

	return lower
}
