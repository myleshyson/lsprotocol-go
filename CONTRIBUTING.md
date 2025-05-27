# Contributing to lsprotocol-golang

Thank you for your interest in contributing to lsprotocol-golang! This project generates Go bindings for the Language Server Protocol specification.

## Project Structure

- `protocol/` - Generated Go code with LSP types and structures
- `lspgenerator/` - Python generator that creates the Go code from LSP specification
- `magefile.go` - Automation scripts for common development tasks (using [mage](https://magefile.org/))

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.24.2+** - Required for building and testing the generated Go code
- **Python 3.12+** - Required for running the code generator
- **uv** - Python package manager for managing dependencies
- **mage** - Build tool for Go projects

### Installing Prerequisites

**uv (Python package manager):**
```bash
# macOS/Linux
curl -LsSf https://astral.sh/uv/install.sh | sh

# or via pip
pip install uv

# or via homebrew (macOS)
brew install uv

go install github.com/magefile/mage@latest
```

### Running Commands
```bash
# Generating types
mage generate

# running tests
mage test
```
