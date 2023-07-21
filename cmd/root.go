package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "Scaffolder is a CLI tool for scaffolding directory structures",
}
