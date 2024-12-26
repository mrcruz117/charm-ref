/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "Greet someone",
	Long:  `This subcommand greets someone with a personalized message.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := "world"
		if len(args) > 0 {
			name = args[0]
		}
		fmt.Printf("Greetings, %s!\n", name)
	},
}

func init() {
	helloworldCmd.AddCommand(greetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// greetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// greetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
