package server

import (
	"simple_db/file"
)

type SimpleDB struct {
	fileManager *file.Manager
}

func NewSimpleDB(dirName string, blockSize int32, buffSize int32) *SimpleDB {
	fileManager, _ := file.NewManager(dirName, blockSize)

	return &SimpleDB{
		fileManager: fileManager,
	}
}
