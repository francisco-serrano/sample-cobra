package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

const (
	Major  = "0"
	Minor  = "1"
	Fix    = "0"
	Verbal = "Initial Version"
)

func versionCmd(w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Describes version",
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = fmt.Fprintf(w, "version: %s.%s.%s-beta %s", Major, Minor, Fix, Verbal)
		},
	}
}
