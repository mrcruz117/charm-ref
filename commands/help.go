package commands

import (
	"fmt"
)

func Help() {
	fmt.Println("Usage: go run . [command]")
	fmt.Println("Commands:")
	fmt.Println("  checkServer   Check the server status")
	fmt.Println("  --help, -h    Show this help message")
}
