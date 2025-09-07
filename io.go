package main

import (
	"os"
	"path/filepath"
)

func loadFile() ([]byte, error) {
	path := filepath.Join("skills", "skills.json")
	return os.ReadFile(path)
}
