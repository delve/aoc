package main

import (
	"os"

	"aocgen/gopherholes/aoc"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "AOC is a tool to support completing Advent of Code puzzles",
	Long:  "AOC supports generating puzzle data, including inputs directly from the website, and benchmarking answers",
}

func Execute() {
	aoc.AttachCmd(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func main() {
	Execute()
}
