package schemaregistry

import (
	"os"
	"path/filepath"
)

func CreateFile(content []byte, name string, path string) (string, error) {
	fullFilePath := filepath.Join(filepath.Clean(path), name)

	err := os.WriteFile(fullFilePath, content, 0600)

	if err != nil {
		return fullFilePath, err
	}

	return fullFilePath, nil
}

func RemoveFile(path string) error {
	err := os.Remove(path)

	return err
}
