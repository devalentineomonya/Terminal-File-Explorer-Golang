package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Application state
type App struct {
	screen       tcell.Screen
	currentDir   string
	contents     []string
	selected     int
	termWidth    int
	termHeight   int
	leftWidth    int
	rightStart   int
	rightWidth   int
	shouldRedraw bool
}

func newApp() (*App, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err := screen.Init(); err != nil {
		return nil, err
	}

	width, height := screen.Size()

	// Get current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	app := &App{
		screen:       screen,
		currentDir:   currentDir,
		contents:     []string{},
		selected:     0,
		termWidth:    width,
		termHeight:   height,
		leftWidth:    width * 40 / 100,
		rightStart:   (width * 40 / 100) + 2,
		rightWidth:   width - (width * 40 / 100) - 3,
		shouldRedraw: true,
	}

	return app, nil
}

func (a *App) handleResize() {
	width, height := a.screen.Size()
	if width != a.termWidth || height != a.termHeight {
		a.termWidth = width
		a.termHeight = height
		a.leftWidth = width * 40 / 100
		a.rightStart = (width * 40 / 100) + 2
		a.rightWidth = width - (width * 40 / 100) - 3
		a.shouldRedraw = true
	}
}

func (a *App) run() {
	// Hide cursor
	a.screen.HideCursor()

	// Event loop
	for {
		a.handleResize()
		a.draw()

		// Poll event
		ev := a.screen.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			a.screen.Sync()
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			case tcell.KeyUp, tcell.KeyCtrlP:
				a.selected--
				if a.selected < 0 {
					a.selected = 0
				}
			case tcell.KeyDown, tcell.KeyCtrlN:
				a.selected++
				if a.selected >= len(a.contents) {
					a.selected = len(a.contents) - 1
				}
			case tcell.KeyEnter:
				a.handleEnter()
			case tcell.KeyRune:
				// Vim-like key bindings
				switch ev.Rune() {
				case 'q', 'e':
					return
				case 'j': // Down
					a.selected++
					if a.selected >= len(a.contents) {
						a.selected = len(a.contents) - 1
					}
				case 'k': // Up
					a.selected--
					if a.selected < 0 {
						a.selected = 0
					}
				case 'h': // Left/parent directory
					a.currentDir = filepath.Dir(a.currentDir)
					a.selected = 0
					a.shouldRedraw = true
				case 'l': // Right/enter directory or open file
					a.handleEnter()
				case 'g': // Go to top
					a.selected = 0
				case 'G': // Go to bottom
					a.selected = len(a.contents) - 1
				case '/': // Search functionality could be implemented here
					// TODO: Implement search
				case 'c': // Command mode
					// Command mode functionality could be implemented here
				}
			}
		}
	}
}

func (a *App) handleEnter() {
	if len(a.contents) > 0 && a.selected < len(a.contents) {
		item := a.contents[a.selected]
		path := filepath.Join(a.currentDir, item)

		if item == ".." {
			a.currentDir = filepath.Dir(a.currentDir)
			a.selected = 0
			a.shouldRedraw = true
		} else if isDirectory(path) {
			a.currentDir = path
			a.selected = 0
			a.shouldRedraw = true
		} else {
			// Temporarily restore terminal
			a.screen.Fini()

			// Open file with default application
			openFile(path)

			// Small delay to let the application start
			time.Sleep(100 * time.Millisecond)

			// Reinit screen
			a.screen.Init()
			a.shouldRedraw = true
		}
	}
}

func (a *App) draw() {
	if a.shouldRedraw {
		a.screen.Clear()

		// Draw borders
		navBorderStyle := tcell.StyleDefault.Foreground(tcell.ColorTeal).Bold(true)
		a.drawBorder(0, 0, a.leftWidth+2, a.termHeight, navBorderStyle, "Navigation")

		previewBorderStyle := tcell.StyleDefault.Foreground(tcell.ColorTeal).Bold(true)
		a.drawBorder(a.rightStart, 0, a.rightWidth+2, a.termHeight, previewBorderStyle, "Preview")

		// Get directory contents
		a.getContents()

		a.shouldRedraw = false
	}

	// Draw navigation panel
	a.drawNavigation()

	// Update preview
	a.updatePreview()

	// Show the result
	a.screen.Show()
}
