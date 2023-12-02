package cmd

import (
	"aoc-2023/day01"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use: "1",
	Run: func(cmd *cobra.Command, args []string) {
		err := day01.Solve(inputFileName, debug)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	solveCmd.AddCommand(cmd)
}
