package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Get directory contents
func (a *App) getContents() {
	a.contents = []string{} // Clear contents

	// Add parent directory
	a.contents = append(a.contents, "..")

	// Read directory
	files, err := ReadDir(a.currentDir)
	if err != nil {
		return
	}

	// First add directories
	for _, file := range files {
		if file.IsDir() {
			a.contents = append(a.contents, file.Name())
		}
	}

	// Then add files
	for _, file := range files {
		if !file.IsDir() {
			a.contents = append(a.contents, file.Name())
		}
	}
}

// Check if path is a directory
func isDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// Check if file is a text file
func isTextFile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	// Read a small amount to check for binary content
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil {
		return false
	}

	// Use MIME type detection
	contentType := http.DetectContentType(buffer[:n])
	return strings.HasPrefix(contentType, "text/") ||
		strings.Contains(contentType, "javascript") ||
		strings.Contains(contentType, "json") ||
		strings.Contains(contentType, "xml")
}

// Open file with default application
func openFile(path string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", path)
	case "darwin":
		cmd = exec.Command("open", path)
	default: // Linux and others
		cmd = exec.Command("xdg-open", path)
	}

	return cmd.Start()
}

// ReadDir reads the directory named by dirname and returns
// a list of directory entries.
func ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

// ReadFile reads the file named by filename and returns the contents.
func ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// SplitLines splits a string into lines.
func SplitLines(s string) []string {
	return strings.Split(s, "\n")
}
