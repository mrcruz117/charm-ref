/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloworldCmd represents the helloworld command
var helloworldCmd = &cobra.Command{
	Use:   "helloworld",
	Short: "Prints 'hello, world' to the console",
	Long: `This command prints 'hello, world' to the console. For example:
	hello, world!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello, world!")
	},
}

func init() {
	rootCmd.AddCommand(helloworldCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloworldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloworldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
