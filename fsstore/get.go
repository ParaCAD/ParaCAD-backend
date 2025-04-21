package fsstore

import (
	"os"
	"path/filepath"
)

func (fs *FSStore) GetFile(fileName string) ([]byte, error) {
	filePath := filepath.Join(fs.storageDirectory, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, nil
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}
