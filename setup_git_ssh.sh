#!/bin/bash

# Load variables from external file
if [ -f ./git_user.env ]; then
    source ./git_user.env
else
    echo "❌ git_user.env file not found!"
    exit 1
fi

# Update and install git
sudo apt update && sudo apt install -y git

# Git global configuration
git config --global user.name "$GIT_USER_NAME"
git config --global user.email "$GIT_USER_EMAIL"
git config --global init.defaultBranch main
git config --global pull.rebase true
git config --global color.ui auto
git config --global core.editor "nano"
git config --global credential.helper store

# Generate SSH key
if [ ! -f ~/.ssh/id_ed25519 ]; then
    ssh-keygen -t ed25519 -C "$GIT_USER_EMAIL" -f ~/.ssh/id_ed25519 -N ""
fi

# Show public key
echo "✅ SSH key generated. Add the following public key to GitHub:"
cat ~/.ssh/id_ed25519.pub

# Test GitHub SSH connection
echo "ℹ️ Run the following command after adding the key to GitHub:"
echo "ssh -T git@github.com"
