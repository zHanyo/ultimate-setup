// internal/ssh/keys.go
package ssh

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
)

func GenerateKey() {
	sshPath := filepath.Join(os.Getenv("HOME"), ".ssh", "id_ed25519")

	// Vérifie si une clé existe déjà
	if _, err := os.Stat(sshPath); err == nil {
		fmt.Println("⚠️ SSH key already exists at", sshPath)
		// Test the SSH connection
		cmd := exec.Command("ssh", "-T", "git@github.com")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("❌ Failed to test SSH connection:", err)
		} else {
			fmt.Println("✅ SSH connection successful.")
		}
		return
	}

	var email string
	prompt := &survey.Input{
		Message: "Enter email for SSH key:",
	}
	if err := survey.AskOne(prompt, &email); err != nil {
		fmt.Println("❌ Failed to get email input:", err)
		return
	}

	cmd := exec.Command("ssh-keygen", "-t", "ed25519", "-C", email, "-f", sshPath, "-N", "")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to generate SSH key:", err)
		return
	}

	fmt.Println("✅ SSH key generated at:", sshPath)
}
