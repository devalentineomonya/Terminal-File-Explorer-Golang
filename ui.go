package main

import (
	"fmt"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

const (
	dirIcon  = "◆ "
	fileIcon = "◇ "
)

func (a *App) drawText(x, y int, style tcell.Style, text string) {
	for _, r := range text {
		a.screen.SetContent(x, y, r, nil, style)
		x += runewidth.RuneWidth(r)
	}
}

func (a *App) drawBorder(x, y, width, height int, style tcell.Style, title string) {
	a.screen.SetContent(x, y, tcell.RuneULCorner, nil, style)
	a.screen.SetContent(x+width-1, y, tcell.RuneURCorner, nil, style)
	a.screen.SetContent(x, y+height-1, tcell.RuneLLCorner, nil, style)
	a.screen.SetContent(x+width-1, y+height-1, tcell.RuneLRCorner, nil, style)

	for i := 1; i < width-1; i++ {
		a.screen.SetContent(x+i, y, tcell.RuneHLine, nil, style)
		a.screen.SetContent(x+i, y+height-1, tcell.RuneHLine, nil, style)
	}

	for i := 1; i < height-1; i++ {
		a.screen.SetContent(x, y+i, tcell.RuneVLine, nil, style)
		a.screen.SetContent(x+width-1, y+i, tcell.RuneVLine, nil, style)
	}

	if title != "" {
		titleX := x + (width-len(title)-4)/2
		a.drawText(titleX, y, style, "ΓöÇΓöÇΓöÇ "+title+" ΓöÇΓöÇΓöÇ")
	}
}

func (a *App) drawNavigation() {
	navStyle := tcell.StyleDefault
	selectedStyle := tcell.StyleDefault.Background(tcell.ColorGray).Bold(true)

	dirStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Bold(true)
	dirAltStyle := tcell.StyleDefault.Foreground(tcell.ColorDarkGray).Bold(true)
	fileStyle := tcell.StyleDefault.Foreground(tcell.ColorYellow)

	for i, item := range a.contents {
		if i >= a.termHeight-2 {
			break
		}

		style := navStyle
		prefix := ""

		if i == a.selected {
			style = selectedStyle
		} else if item == ".." || isDirectory(filepath.Join(a.currentDir, item)) {
			if i%2 == 0 {
				style = dirStyle
			} else {
				style = dirAltStyle
			}
			prefix = dirIcon
		} else {
			style = fileStyle
			prefix = fileIcon
		}

		for j := 1; j < a.leftWidth; j++ {
			a.screen.SetContent(j, i+1, ' ', nil, navStyle)
		}

		a.drawText(1, i+1, style, " "+prefix+item)
	}
}

func (a *App) updatePreview() {
	previewStyle := tcell.StyleDefault.Foreground(tcell.ColorLightGray)

	if len(a.contents) == 0 || a.selected >= len(a.contents) {
		return
	}

	item := a.contents[a.selected]
	path := filepath.Join(a.currentDir, item)
	var previewContent []string

	if item == ".." {
		previewContent = append(previewContent, "Parent Directory:")
		parentDir := filepath.Dir(a.currentDir)
		files, err := ReadDir(parentDir)
		if err == nil {
			for _, file := range files {
				prefix := fileIcon
				if file.IsDir() {
					prefix = dirIcon
				}
				info := fmt.Sprintf("%s %s %8d %s", prefix, file.Mode().String(), file.Size(), file.Name())
				previewContent = append(previewContent, info)
				if len(previewContent) >= a.termHeight-4 {
					break
				}
			}
		}
	} else if isDirectory(path) {
		previewContent = append(previewContent, "Directory Preview:")
		files, err := ReadDir(path)
		if err == nil {
			for _, file := range files {
				prefix := fileIcon
				if file.IsDir() {
					prefix = dirIcon
				}
				info := fmt.Sprintf("%s %s %8d %s", prefix, file.Mode().String(), file.Size(), file.Name())
				previewContent = append(previewContent, info)
				if len(previewContent) >= a.termHeight-4 {
					break
				}
			}
		}
	} else if isTextFile(path) {
		previewContent = append(previewContent, "File Preview:")
		data, err := ReadFile(path)
		if err == nil {
			lines := SplitLines(string(data))
			for i, line := range lines {
				if i >= a.termHeight-6 {
					break
				}
				previewContent = append(previewContent, line)
			}
		}
	} else {
		previewContent = append(previewContent, "File Preview:")
		previewContent = append(previewContent, "Binary file")
	}

	for i := 1; i < a.termHeight-1; i++ {
		for j := a.rightStart + 1; j < a.rightStart+a.rightWidth; j++ {
			a.screen.SetContent(j, i, ' ', nil, previewStyle)
		}
	}

	for i, line := range previewContent {
		if i >= a.termHeight-2 {
			break
		}

		displayLine := line
		if len(displayLine) > a.rightWidth-2 {
			displayLine = displayLine[:a.rightWidth-2]
		}

		a.drawText(a.rightStart+1, i+1, previewStyle, displayLine)
	}
}
