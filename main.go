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
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--help":
			fmt.Println("Usage: ultimate-setup [options]")
			fmt.Println("Options:")
			fmt.Println("  --help                 Show this help message")
			fmt.Println("  configure-git          Configure Git settings")
			fmt.Println("  generate-ssh-key       Generate a new SSH key")
			fmt.Println("  install-cli-tools      Install common CLI tools")
			fmt.Println("  uninstall-cli-tools    Uninstall CLI tools")
			os.Exit(0)
		case "configure-git":
			git.SetupGit()
			os.Exit(0)
		case "generate-ssh-key":
			ssh.GenerateKey()
			os.Exit(0)
		case "install-cli-tools":
			tools.Install()
			os.Exit(0)
		case "uninstall-cli-tools":
			tools.Uninstall()
			os.Exit(0)
		default:
			fmt.Printf("Unknown option: %s\n", os.Args[1])
			os.Exit(1)
		}
	}

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
