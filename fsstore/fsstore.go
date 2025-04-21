package fsstore

import "os"

type FSStore struct {
	storageDirectory string
}

func New(storageDirectory string) (*FSStore, error) {
	err := os.MkdirAll(storageDirectory, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return &FSStore{
		storageDirectory: storageDirectory,
	}, nil
}
