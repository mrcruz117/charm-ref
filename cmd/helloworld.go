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
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		quiet, _ := cmd.Flags().GetBool("quiet")
		name := "world"
		if len(args) > 0 {
			name = args[0]
		}
		switch {
		case quiet:
			fmt.Printf("hello, %s\n", name)
		default:
			fmt.Printf("hello, %s!\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(helloworldCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	helloworldCmd.Flags().BoolP("quiet", "q", false, "Prints the message without an exclamation mark")
}
