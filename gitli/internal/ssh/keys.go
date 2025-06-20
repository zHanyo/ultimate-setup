// internal/ssh/keys.go
package ssh

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
)

func GenerateKey(email ...string) {
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

	var userEmail string
	if len(email) > 0 {
		userEmail = email[0]
	} else {
		prompt := &survey.Input{
			Message: "Enter email for SSH key:",
		}
		if err := survey.AskOne(prompt, &userEmail); err != nil {
			fmt.Println("❌ Failed to get email input:", err)
			return
		}
	}

	cmd := exec.Command("ssh-keygen", "-t", "ed25519", "-C", userEmail, "-f", sshPath, "-N", "")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to generate SSH key:", err)
		return
	}

	fmt.Println("✅ SSH key generated at:", sshPath)
	pubKeyPath := sshPath + ".pub"
	pubKeyBytes, err := os.ReadFile(pubKeyPath)
	if err != nil {
		fmt.Println("🔴 Could not read public key:", err)
	} else {
		fmt.Println("\n🔑 Your public key:")
		fmt.Println(string(pubKeyBytes))
		fmt.Println("\n💡 Copy this to your Git provider:\n   🔗 GitHub → https://github.com/settings/keys\n   🔗 GitLab → https://gitlab.com/-/profile/keys\n   🔗 Bitbucket → https://bitbucket.org/account/settings/ssh-keys/")}	
	}
