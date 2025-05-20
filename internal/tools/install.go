// internal/tools/install.go
package tools

import (
	"fmt"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

func Install() {
	options := []string{
		"curl",
		"ansible",
		"terraform",
	}

	selectedTools := []string{}
	prompt := &survey.MultiSelect{
		Message: "📦 Which CLI tools do you want to install?",
		Options: options,
	}
	err := survey.AskOne(prompt, &selectedTools)
	if err != nil {
		fmt.Println("❌ Failed to display tool selection menu:", err)
		return
	}

	for _, tool := range selectedTools {
		switch tool {
		case "curl":
			installCurl()
		case "ansible":
			installAnsible()
		case "terraform":
			fmt.Println("⚠️ Terraform is better used from your CI/CD pipeline. Consider avoiding local installations. it wont work anyway so install it by yourself if really needed.")
		}
	}

	fmt.Println("\n🎉 CLI tools installation process complete.")
}

func Uninstall() {
	options := []string{
		"curl",
		"ansible",
		"terraform",
	}

	selectedTools := []string{}
	prompt := &survey.MultiSelect{
		Message: "🗑️ Which CLI tools do you want to uninstall?",
		Options: options,
	}
	err := survey.AskOne(prompt, &selectedTools)
	if err != nil {
		fmt.Println("❌ Failed to display tool selection menu:", err)
		return
	}

	for _, tool := range selectedTools {
		switch tool {
		case "curl":
			uninstallCurl()
		case "ansible":
			uninstallAnsible()
		}
	}

	fmt.Println("\n✅ CLI tools uninstallation process complete.")
}

func installCurl() {
	fmt.Println("Installing curl...")
	err := exec.Command("sh", "-c", "sudo apt install -y curl").Run()
	if err != nil {
		fmt.Println("❌ Failed to install curl:", err)
		return
	}
	fmt.Println("✔ curl installed successfully.")
}

func installAnsible() {
	fmt.Println("Installing ansible...")
	err := exec.Command("sh", "-c", "sudo apt update && sudo apt install -y software-properties-common && sudo add-apt-repository --yes --update ppa:ansible/ansible && sudo apt install -y ansible").Run()
	if err != nil {
		fmt.Println("❌ Failed to install ansible:", err)
		return
	}
	fmt.Println("✔ ansible installed successfully.")
}

func uninstallCurl() {
	fmt.Println("Uninstalling curl...")
	err := exec.Command("sh", "-c", "sudo apt remove -y curl").Run()
	if err != nil {
		fmt.Println("❌ Failed to uninstall curl:", err)
		return
	}
	fmt.Println("✔ curl uninstalled successfully.")
}

func uninstallAnsible() {
	fmt.Println("Uninstalling ansible...")
	err := exec.Command("sh", "-c", "sudo apt remove -y ansible").Run()
	if err != nil {
		fmt.Println("❌ Failed to uninstall ansible:", err)
		return
	}
	fmt.Println("✔ ansible uninstalled successfully.")
}

func InstallSpecific(tool string) {
	switch tool {
	case "curl":
		installCurl()
	case "ansible":
		installAnsible()
	case "terraform":
		fmt.Println("⚠️ Terraform is better used from your CI/CD pipeline. Consider avoiding local installations. It won't work anyway, so install it by yourself if really needed.")
	default:
		fmt.Printf("❌ Unknown tool: %s\n", tool)
	}
}

func UninstallSpecific(tool string) {
	switch tool {
	case "curl":
		uninstallCurl()
	case "ansible":
		uninstallAnsible()
	default:
		fmt.Printf("❌ Unknown tool: %s\n", tool)
	}
}
