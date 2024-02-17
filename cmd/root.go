/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocc",
	Short: "Collection of coding challenges in Go",
	Long: `This is a collection of solutions to coding challenges in Go. All of the implementations are tested and benchmarked. The goal is to provide a reference for Go developers to learn from and to compare different approaches to solving the same problem.
	
TDD is used to drive the development of the solutions. Learn more about the challenges on codingchallenges.fyi`,
}

// Execute executes the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
