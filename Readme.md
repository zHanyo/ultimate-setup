# Ultimate Setup 🚀

Ultimate Setup is a CLI tool designed to streamline the setup process for developers. It provides an easy way to configure Git, generate SSH keys, and manage the installation and uninstallation of essential CLI tools.

## Features

- **Git Configuration**: Quickly set up your Git username, email, and other global configurations.
- **SSH Key Management**: Generate a new SSH key or test your existing SSH connection to GitHub.
- **CLI Tools Management**: Install or uninstall essential CLI tools like:
  - `curl`
  - `ansible`
  - `terraform` (with a recommendation to use it in CI/CD pipelines)

## Prerequisites

- **Go**: Ensure you have Go installed (version 1.22.3 or later).
- **Linux/WSL**: This tool is designed to run in a Linux or WSL environment.

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd ultimate-setup
   ```

2. Build the project:
   ```bash
   go build -o ultimate-setup
   ```

3. Run the tool:
   ```bash
   ./ultimate-setup
   ```

## Usage

When you run the tool, you'll be presented with a menu of options:

1. **Configure Git**:
   - Set up your Git username, email, and other global configurations.
   - Example:
     ```
     Your first name: John
     Your last name: Doe
     Your email address: john.doe@example.com
     ```

2. **Generate SSH Key**:
   - Generate a new SSH key or test the connection if a key already exists.
   - Example:
     ```
     ⚠️ SSH key already exists at /home/user/.ssh/id_ed25519
     ✅ SSH connection successful.
     ```

3. **Install CLI Tools**:
   - Select tools to install from the following options:
     - `curl`
     - `ansible`
     - `terraform` (with a warning to use it in CI/CD pipelines)
   - Example:
     ```
     📦 Which CLI tools do you want to install? [curl, ansible]
     ```

4. **Uninstall CLI Tools**:
   - Select tools to uninstall from the installed options.

5. **Exit**:
   - Exit the program.

## Testing Git SSH Connection

To test your Git SSH connection manually, run the following command:

```bash
ssh -T git@github.com
```

### Expected Output

If successful, you should see:

```
Hi <username>! You've successfully authenticated, but GitHub does not provide shell access.
```

If it fails, ensure your SSH key is added to the SSH agent:

```bash
ssh-add ~/.ssh/id_ed25519
```

Also, verify that your public key is added to your GitHub account under **Settings > SSH and GPG keys**.

## Project Structure

```
ultimate-setup/
├── .gitignore
├── go.mod
├── go.sum
├── main.go
├── Readme.md
├── internal/
│   ├── git/
│   │   └── config.go
│   ├── ssh/
│   │   └── keys.go
│   ├── tools/
│   │   └── install.go
│   └── prompt/
│       └── menu.go
```

## Development

### Adding a New Tool
1. Add the tool's installation and uninstallation logic in `internal/tools/install.go`.
2. Update the `Install` and `Uninstall` functions to include the new tool.

### Testing
Run the project locally:
```bash
go run main.go
```

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

Happy coding! 🎉