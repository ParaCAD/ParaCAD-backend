package fsstore

import (
	"os"
	"path/filepath"
)

func (fs *FSStore) SaveFile(fileName string, fileContent []byte) error {
	filePath := filepath.Join(fs.storageDirectory, fileName)

	err := os.WriteFile(filePath, fileContent, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
