package main

import (
	"os"
	"path/filepath"
)

func getEntriesFromPath(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func getRootPath() (string, error) {
	f, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fp := filepath.Dir(f)
	return fp, nil
}
