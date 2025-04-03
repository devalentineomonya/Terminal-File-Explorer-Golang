# File Browser

A terminal-based file browser with vim-style navigation and a preview panel.

## Features

- Two-panel interface with file navigation and preview
- Diamond icons for directories and files
- Alternating colors for directories
- Vim-style keyboard navigation
- Preview support for text files and directories

## Navigation

- Arrow keys: Move up/down
- Enter: Enter directory or open file
- Vim bindings:
  - `j`: Move down
  - `k`: Move up
  - `h`: Go to parent directory
  - `l`: Enter directory or open file
  - `g`: Go to top of list
  - `G`: Go to bottom of list
- `q` or `e`: Quit application

## Building and Running

### Using Make

```bash
# Build the application
make build

# Run the application
make run

# Clean build files
make clean
```

### Using Go commands

```bash
# Build the application
go build -o filebrowser *.go

# Run the application
./filebrowser
```

## Dependencies

- github.com/gdamore/tcell/v2: Terminal cell library
- github.com/mattn/go-runewidth: String width calculation
