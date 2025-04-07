# Contribution Guidelines

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](https://github.com/devalentineomonya/Terminal-File-Explorer-Golang/pulls)
[![Open Issues](https://img.shields.io/github/issues-raw/devalentineomonya/Terminal-File-Explorer-Golang)](https://github.com/devalentineomonya/Terminal-File-Explorer-Golang/issues)

We welcome contributions from developers of all skill levels! Please take a moment to review these guidelines before submitting your contribution.

## Table of Contents
1. [Ways to Contribute](#ways-to-contribute)
2. [Development Setup](#development-setup)
3. [Pull Request Guidelines](#pull-request-guidelines)
4. [Code Style Guide](#code-style-guide)
5. [Documentation Standards](#documentation-standards)
6. [Code of Conduct](#code-of-conduct)

## Ways to Contribute

### 🐛 Report Bugs
- Check existing issues before creating new ones
- Use the bug report template
- Include:
  - OS and terminal environment
  - Steps to reproduce
  - Expected vs actual behavior
  - Relevant screenshots

### 💡 Suggest Features
- Describe the use case clearly
- Explain potential alternatives
- Include mockups if proposing UI changes

### 👩💻 Code Contributions
- Start with "good first issue" labeled tickets
- Keep PRs focused on single objectives
- Add tests for new functionality

### 📖 Improve Documentation
- Fix typos and update outdated information
- Add usage examples
- Translate documentation

### 🤝 Community Help
- Answer questions in Discussions
- Review open PRs
- Share project with others

## Development Setup

### Prerequisites
- Go 1.21+
- GNU Make 4.3+
- Docker 24.0+ (for containerized testing)

### Local Environment

```bash
# Fork and clone repository
git clone https://github.com/<your-username>/Terminal-File-Explorer-Golang.git
cd Terminal-File-Explorer-Golang

# Install dependencies
make deps

# Build project
make build

# Run tests
make test

# Start development mode
make run
```

### Docker Development

```bash
# Build development container
make docker-dev

# Enter container shell
make docker-shell

# Run tests in container
make docker-test
```

## Pull Request Guidelines

1   - `feat/description` for new features
   - `fix/description` for bug fixes
   - `docs/description` for documentation

2. Keep commits atomic and squashed logically

3. Reference related issues using #<issue-number>

4. Add/update tests in `_test.go` files

5. Update documentation and README if needed

6. Include before/after screenshots for UI changes

## Code Style Guide

### Go Formatting
- Follow standard Go formatting (`gofmt`)
- Maximum line length: 120 characters
- Use `camelCase` for variables
- Document exported functions with GoDoc

### Commit Messages
Follow [Conventional Commits](https://www.conventionalcommits.org) format:
```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

Valid types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Formatting changes
- `refactor`: Code refactoring
- `test`: Test-related changes
- `chore`: Maintenance tasks

### TUI Specific
- Keep UI components in `internal/ui/`
- Use tcell's style constants for colors
- Separate business logic from presentation

## Documentation Standards
- Use Markdown format
- Add code examples with syntax highlighting
- Update both user and developer docs
- Keep comments focused on "why" rather than "what"

## Code of Conduct
All contributors must adhere to our [Code of Conduct](CODE_OF_CONDUCT.md). Please report unacceptable behavior to [email protected].

## Recognition
- Contributors will be listed in the project README
- Significant contributions earn commit access
- Top contributors receive special mentions in release notes

---

**Need Help?**  
Join our [Discussions](https://github.com/devalentineomonya/Terminal-File-Explorer-Golang/discussions) or contact [@devalentineomonya](https://github.com/devalentineomonya)