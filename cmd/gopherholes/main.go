package main

import (
	"os"

	"gopherholes/pkg/gen"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gopherholes",
	Short: "AOC is a tool to support completing Advent of Code puzzles",
	Long:  "AOC supports generating puzzle data, including inputs directly from the website, and benchmarking answers",
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "AOC is a tool to support completing Advent of Code puzzles",
	Long:  "AOC supports generating puzzle data, including inputs directly from the website, and benchmarking answers",
	Run: func(cmd *cobra.Command, args []string) {
		gen.Init()
	},
}

func Execute() {
	rootCmd.AddCommand(setupCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func main() {
	Execute()
}
