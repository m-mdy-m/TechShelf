package common

import (
	"errors"
	"os"
	"path/filepath"
)

func CatalogPath(override string) (string, error) {
	if override != "" {
		return override, nil
	}
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cwd, "catalog.json"), nil
}

func EnsureDir(path string) error {
	if path == "" {
		return errors.New("empty dir path")
	}
	return os.MkdirAll(path, 0o755)
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func ListCategories(shelvesDir string) ([]string, error) {
	entries, err := os.ReadDir(shelvesDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var cats []string
	for _, e := range entries {
		if e.IsDir() {
			cats = append(cats, e.Name())
		}
	}
	return cats, nil
}
