package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func (a *App) getContents() {
	a.contents = []string{}

	a.contents = append(a.contents, "..")

	files, err := ReadDir(a.currentDir)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			a.contents = append(a.contents, file.Name())
		}
	}

	for _, file := range files {
		if !file.IsDir() {
			a.contents = append(a.contents, file.Name())
		}
	}
}

func isDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func isTextFile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil {
		return false
	}

	contentType := http.DetectContentType(buffer[:n])
	return strings.HasPrefix(contentType, "text/") ||
		strings.Contains(contentType, "javascript") ||
		strings.Contains(contentType, "json") ||
		strings.Contains(contentType, "xml")
}

func openFile(path string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", path)
	case "darwin":
		cmd = exec.Command("open", path)
	default:
		cmd = exec.Command("xdg-open", path)
	}

	return cmd.Start()
}

func ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

func ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func SplitLines(s string) []string {
	return strings.Split(s, "\n")
}
