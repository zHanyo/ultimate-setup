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

	if len(email) == 0 || email[0] == "" {
		fmt.Println("❌ Email is required for non-interactive SSH key generation.")
		return
	}

	userEmail := email[0]

	cmd := exec.Command("ssh-keygen", "-t", "ed25519", "-C", userEmail, "-f", sshPath, "-N", "")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to generate SSH key:", err)
		return
	}

	fmt.Println("✅ SSH key generated at:", sshPath)
}
