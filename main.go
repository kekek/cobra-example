package main

import (
	"os"

	"github.com/spf13/cobra"

	"cobra-example/sub/get"
	"cobra-example/sub/list"
)

var Version string

var rootCmd = &cobra.Command{
	Use:   "cobra-example",
	Short: "练手项目，演示如何使用cobra的自动补全",
}

func init() {
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(get.GetCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
