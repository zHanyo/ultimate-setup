// main.go
package main

import (
	"fmt"
	"os"

	"ultimate-setup/internal/git"
	"ultimate-setup/internal/ssh"
	"ultimate-setup/internal/tools"
	"ultimate-setup/internal/prompt"
)

func main() {
	fmt.Println("🚀 Welcome to Ultimate Setup")

	// Define the menu options
	options := []string{
		"Configure Git",
		"Generate SSH key",
		"Install CLI tools",
		"Uninstall CLI tools",
		"Exit",
	}

	// Pass the options to ShowMenu
	choices, err := prompt.ShowMenu(options)
	if err != nil {
		fmt.Println("❌ Failed to display menu:", err)
		os.Exit(1)
	}

	for _, choice := range choices {
		switch choice {
		case "Configure Git":
			git.SetupGit()
		case "Generate SSH key":
			ssh.GenerateKey()
		case "Install CLI tools":
			tools.Install()
		case "Uninstall CLI tools":
			tools.Uninstall()
		case "Exit":
			fmt.Println("👋 Exiting...")
			os.Exit(0)
		}
	}
}
