# Terminal File Explorer

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go) 
![TUI](https://img.shields.io/badge/TUI-Advanced-4DD0E1?logo=terminal) 
![Docker](https://img.shields.io/badge/Docker-24.0+-2496ED?logo=docker) 
![Make](https://img.shields.io/badge/Make-4.3-6C4CAB?logo=gnu-bash) 
![Homebrew](https://img.shields.io/badge/Homebrew-4.0+-FBB040?logo=homebrew) 
![Community](https://img.shields.io/badge/Community-Active-2ECC71?logo=github) 
![License](https://img.shields.io/badge/License-MIT-blue)

![TUI Screenshot](https://raw.githubusercontent.com/devalentineomonya/Terminal-File-Explorer-Golang/main/.github/screenshot.png)  
*A Vim-inspired terminal file manager with modern UX/UI components*

[![GitHub stars](https://img.shields.io/github/stars/devalentineomonya/Terminal-File-Explorer-Golang?style=social)](https://github.com/devalentineomonya/Terminal-File-Explorer-Golang/stargazers)
[![GitHub release](https://img.shields.io/github/v/release/devalentineomonya/Terminal-File-Explorer-Golang)](https://github.com/devalentineomonya/Terminal-File-Explorer-Golang/releases)

## ✨ Key Features  

### Navigation & Interface  
🖥️ **Dual-Pane Layout**  
- Left panel: Hierarchical directory navigation  
- Right panel: File previews (text/images) & metadata  

🎮 **Vim-Style Controls**  
- Muscle-memory friendly keybindings (hjkl navigation)  
- Mode switching (normal/visual/command)  

### Deployment Options  
🐋 **Containerized**  
- Pre-built Docker images with Alpine base  
- Volume mounting support  

🍺 **Homebrew Ready**  
- Single-command installation for macOS/Linux  
- Automatic updates via brew  

## 🚀 Quick Start  

### Installation  

```bash  
# Via Docker
docker run -it --rm -v $PWD:/data ghcr.io/devalentineomonya/terminal-file-explorer:latest

# Homebrew 
brew tap devalentineomonya/tools
brew install terminal-file-explorer

# From Source
git clone https://github.com/devalentineomonya/Terminal-File-Explorer-Golang.git
cd Terminal-File-Explorer-Golang
make build && make install
```

### Make Commands  
```makefile
build    # Compile binary
install  # Install system-wide
docker   # Build Docker image
clean    # Remove build artifacts
run      # Start in dev mode
```

## 🐋 Docker Compose  

```yaml
version: '3.8'
services:
  file-explorer:
    image: ghcr.io/devalentineomonya/terminal-file-explorer:latest
    volumes:
      - ./:/workspace
    environment:
      - TUI_THEME=dark
    devices:
      - /dev/fuse:/dev/fuse
```

## 🛠 Development  

```bash
# Build with Make
make build

# Run tests
make test

# Create release builds
make release

# Build Docker image
make docker
```

## 🌐 Community  

[![GitHub Discussions](https://img.shields.io/badge/Join-Discussions-1DA1F2?logo=github)](https://github.com/devalentineomonya/Terminal-File-Explorer-Golang/discussions) 
[![Twitter Follow](https://img.shields.io/badge/Follow-@devalentineomonya-1DA1F2?logo=twitter)](https://twitter.com/devalentineomonya)

## 🛣 Roadmap  

[![GitHub Milestones](https://img.shields.io/badge/View-Milestones-6C4CAB)](https://github.com/devalentineomonya/Terminal-File-Explorer-Golang/milestones)

## 📜 License  

MIT License - See [LICENSE](https://github.com/devalentineomonya/Terminal-File-Explorer-Golang/blob/main/LICENSE)

---

**Maintainer**: [Valentine Omonya](https://github.com/devalentineomonya)  
**Contribute**: [Guidelines](https://github.com/devalentineomonya/Terminal-File-Explorer-Golang/blob/main/CONTRIBUTING.md)  
**Support**: [Open an Issue](https://github.com/devalentineomonya/Terminal-File-Explorer-Golang/issues/new/choose)