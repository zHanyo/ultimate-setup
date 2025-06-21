# Agents.md – Guidance for OpenAI Codex

This `Agents.md` file provides detailed instructions for OpenAI Codex and other AI agents collaborating on this repository. It includes project structure, conventions, testing, and update policies.

---

## 🗂️ Project Structure for Codex Navigation

- `/gitli`: Go CLI tool source code  
  - `/internal/git`: Git configuration and logic modules  
  - `/internal/prompt`: Terminal menu interface  
  - `/internal/ssh`: SSH key logic  
  - `/internal/tools`: Tool installers (e.g. Git, Zsh, etc.)  
  - `main.go`: Entry point of the CLI  
  - `go.mod`, `go.sum`: Go module files  

- `/deven`: Docker-based development environment  
  - `Dockerfile`: Build file for dev container  
  - `requirements-apt.txt`: Required apt packages  
  - `requirements-pip.txt`: Python dependencies (if used)  
  - `requirements-tools.sh`: Custom tool installation script  

- `docker-build-run.sh`: Script to build and run the dev container  
- `.github/workflows`: GitHub Actions for CI  
  - `build-and-test.yml`: Build and test pipeline  
  - `dev-container.yml`: Dev environment pipeline  
  - `release.yml`: Release automation  

---

## ✅ Coding Standards and Best Practices

### General Guidelines

- **Language**: Go (Golang)  
- **Naming**: Use `camelCase` for variables, `PascalCase` for types, and `snake_case` for file names.  
- **No special characters or accents** should be used in file names, documentation, or code comments to avoid encoding/display issues, especially on GitHub.
- **Comment complex logic**, especially anything involving git operations, SSH, or Docker internals.

### Go CLI Conventions

- Keep functions small and composable.
- Store non-exported logic in the `internal/` directory.
- Interface-based design is preferred for future mocking and testability.
- User-facing commands must use clear prompts (TUI-friendly).

---

## 🧪 Testing Requirements for OpenAI Codex

Tests will be added in a later phase, but Codex should:

- Write unit tests in Go using the standard `testing` package.
- Ensure test files are named `*_test.go` and co-located with the tested code.
- Future test commands will include:

```bash
# Run all tests
go test ./...

# Run a specific package
go test ./gitli/internal/git
```

Codex should use `t.Run()` to structure subtests and ensure descriptive test names.

---

## 📘 Codex Responsibilities

OpenAI Codex must:

1. **Update `Readme.md`** whenever a change affects:
   - How to run the CLI
   - How to build or use the Docker container
   - Installation instructions for dependencies

2. **Never introduce accented characters or locale-specific symbols** in any code, config, or documentation.

3. **Ensure CLI remains interactive and intuitive**—CLI improvements should be incremental, tested, and clearly documented.

4. If a change affects usability or user commands (e.g., a new prompt option, altered flag, new setup step), Codex must reflect this in the README.

---

## 🚦 CI and Linting (Work in Progress)

GitHub Actions are used for CI via the following workflows:

- `build-and-test.yml` – Compiles and tests the CLI.
- `dev-container.yml` – Validates the Docker dev environment.
- `release.yml` – Automates tagged releases.

> A future linting pipeline will include `golangci-lint`.

Codex should prepare the codebase for it by:
- Avoiding unused variables/imports
- Using `err != nil` guards consistently
- Adding proper Go doc comments on exported symbols
