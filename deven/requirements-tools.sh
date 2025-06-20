#!/bin/bash
set -euxo pipefail

# Install kubectl using new packages repository
curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.30/deb/Release.key | gpg --dearmor -o /usr/share/keyrings/kubernetes-apt-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.30/deb/ /" >/etc/apt/sources.list.d/kubernetes.list
apt-get update && apt-get install -y kubectl

# Install helm
curl https://baltocdn.com/helm/signing.asc | gpg --dearmor -o /usr/share/keyrings/helm.gpg
echo "deb [signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main" >/etc/apt/sources.list.d/helm-stable-debian.list
apt-get update && apt-get install -y helm

# Install Terraform
curl -fsSL https://apt.releases.hashicorp.com/gpg | gpg --dearmor -o /usr/share/keyrings/hashicorp.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp.gpg] https://apt.releases.hashicorp.com $(. /etc/os-release && echo $VERSION_CODENAME) main" >/etc/apt/sources.list.d/hashicorp.list
apt-get update && apt-get install -y terraform

# Install Ansible
apt-get update && apt-get install -y ansible
