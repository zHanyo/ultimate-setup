// internal/git/config.go
package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func SetupGit() {
	fmt.Println("🔧 Configuring Git...")

	var firstName, lastName, email string

	// Prompts
	survey.AskOne(&survey.Input{
		Message: "Your first name:",
	}, &firstName)

	survey.AskOne(&survey.Input{
		Message: "Your last name:",
	}, &lastName)

	survey.AskOne(&survey.Input{
		Message: "Your email address:",
	}, &email)

	fullName := strings.TrimSpace(firstName + " " + lastName)

	// Git config
	if err := exec.Command("git", "config", "--global", "user.name", fullName).Run(); err != nil {
		fmt.Println("❌ Failed to set Git user name:", err)
		return
	}
	if err := exec.Command("git", "config", "--global", "user.email", email).Run(); err != nil {
		fmt.Println("❌ Failed to set Git user email:", err)
		return
	}
	if err := exec.Command("git", "config", "--global", "init.defaultBranch", "main").Run(); err != nil {
		fmt.Println("❌ Failed to set default branch:", err)
		return
	}
	if err := exec.Command("git", "config", "--global", "pull.rebase", "true").Run(); err != nil {
		fmt.Println("❌ Failed to set pull rebase:", err)
		return
	}
	if err := exec.Command("git", "config", "--global", "color.ui", "auto").Run(); err != nil {
		fmt.Println("❌ Failed to set color UI:", err)
		return
	}

	// Display configured Git settings
	fmt.Println("🔍 Verifying Git configuration...")
	displayGitConfig("user.name")
	displayGitConfig("user.email")
	displayGitConfig("init.defaultBranch")
	displayGitConfig("pull.rebase")
	displayGitConfig("color.ui")

	fmt.Printf("✅ Git configured: %s <%s>\n", fullName, email)
}

func displayGitConfig(key string) {
	cmd := exec.Command("git", "config", "--global", key)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("❌ Failed to retrieve %s: %v\n", key, err)
		return
	}
	fmt.Printf("✔ %s: %s\n", key, strings.TrimSpace(string(output)))
}
