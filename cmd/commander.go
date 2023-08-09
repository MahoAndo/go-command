package cmd

import "github.com/spf13/cobra"

type Builder interface {
	Build() *cobra.Command
}
