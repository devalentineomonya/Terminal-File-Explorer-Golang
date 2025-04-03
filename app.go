package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/gdamore/tcell/v2"
)

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

	currentDir, err :=os.Getwd()
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
	a.screen.HideCursor()

	for {
		a.handleResize()
		a.draw()

		ev := a.screen.PollEvent()

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
				switch ev.Rune() {
				case 'q', 'e':
					return
				case 'j':
					a.selected++
					if a.selected >= len(a.contents) {
						a.selected = len(a.contents) - 1
					}
				case 'k':
					a.selected--
					if a.selected < 0 {
						a.selected = 0
					}
				case 'h':
					a.currentDir = filepath.Dir(a.currentDir)
					a.selected = 0
					a.shouldRedraw = true
				case 'l':
					a.handleEnter()
				case 'g':
					a.selected = 0
				case 'G':
					a.selected = len(a.contents) - 1
				case '/':

				case 'c':
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
			a.screen.Fini()

			openFile(path)

			time.Sleep(100 * time.Millisecond)

			a.screen.Init()
			a.shouldRedraw = true
		}
	}
}

func (a *App) draw() {
	if a.shouldRedraw {
		a.screen.Clear()

		navBorderStyle := tcell.StyleDefault.Foreground(tcell.ColorTeal).Bold(true)
		a.drawBorder(0, 0, a.leftWidth+2, a.termHeight, navBorderStyle, "Navigation")

		previewBorderStyle := tcell.StyleDefault.Foreground(tcell.ColorTeal).Bold(true)
		a.drawBorder(a.rightStart, 0, a.rightWidth+2, a.termHeight, previewBorderStyle, "Preview")

		a.getContents()

		a.shouldRedraw = false
	}

	a.drawNavigation()

	a.updatePreview()

	a.screen.Show()
}
