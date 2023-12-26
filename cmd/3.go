package cmd

import (
	"aoc-2023/day03"

	"github.com/spf13/cobra"
)

var cmd3 = &cobra.Command{
	Use: "3",
	Run: func(cmd *cobra.Command, args []string) {
		err := day03.Solve(inputFileName, debug)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	solveCmd.AddCommand(cmd3)
}
