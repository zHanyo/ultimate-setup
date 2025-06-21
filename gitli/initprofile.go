package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func initProfile() {
	baseDir := "configs"
	readmes := map[string]string{
		"ssh":       "# SSH Configuration\n\nPut your SSH keys and config here (e.g., GitHub keys).\nRecommended: `id_ed25519`, `id_ed25519.pub`, and `config`.\n\nDo not commit sensitive files.\n",
		"kube":      "# Kubernetes Configuration\n\nPlace your kubeconfig file here (default: `config`).\n\nTo test: `kubectl get pods --context=...`\n",
		"aws":       "# AWS CLI Configuration\n\nPut `credentials` and `config` files here.\n\nTo generate them: `aws configure`\n",
		"azure":     "# Azure CLI Configuration\n\nAzure stores its configuration in `~/.azure`.\n\nYou can copy the folder here after running `az login`.\n",
		"gcp":       "# Google Cloud CLI Configuration\n\nPut your `application_default_credentials.json` and `.config/gcloud` files here.\n\nCan be generated with: `gcloud auth application-default login`\n",
		"ansible":   "# Ansible Configuration\n\nPut your `ansible.cfg`, `inventory`, vault passwords or playbooks here.\n",
		"helm":      "# Helm Configuration\n\nStore your Helm plugins, repo configs, and environment-specific values here.\n",
		"terraform": "# Terraform Configuration\n\nUse this folder to manage `terraform.tfvars`, backend config files, and provider credentials.\n\nThis keeps your `.terraform` state folder outside source control.\n",
	}

	if err := os.MkdirAll(baseDir, 0755); err != nil {
		fmt.Println("❌ Failed to create", baseDir+":", err)
		return
	}

	for name, content := range readmes {
		dirPath := filepath.Join(baseDir, name)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			fmt.Println("❌ Failed to create", dirPath+":", err)
			continue
		}

		keepPath := filepath.Join(dirPath, ".keep")
		if _, err := os.Stat(keepPath); os.IsNotExist(err) {
			if err := os.WriteFile(keepPath, []byte{}, 0644); err != nil {
				fmt.Println("❌ Failed to create", keepPath+":", err)
			}
		}

		readmePath := filepath.Join(dirPath, "README.md")
		if _, err := os.Stat(readmePath); os.IsNotExist(err) {
			if err := os.WriteFile(readmePath, []byte(content), 0644); err != nil {
				fmt.Println("❌ Failed to create", readmePath+":", err)
			}
		}

		fmt.Printf("✔ Created %s/\n", dirPath)
	}
}
