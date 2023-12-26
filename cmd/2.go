package cmd

import (
	"aoc-2023/day02"

	"github.com/spf13/cobra"
)

var cmd2 = &cobra.Command{
	Use: "2",
	Run: func(cmd *cobra.Command, args []string) {
		err := day02.Solve(inputFileName, debug)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	solveCmd.AddCommand(cmd2)
}
