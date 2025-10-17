package server

import (
	"simple_db/file"
	"simple_db/log"
)

type SimpleDB struct {
	fileManager *file.Manager
	logManager  *log.Manager
}

func NewSimpleDB(dirName string, blockSize int32, buffSize int32) *SimpleDB {
	fileManager, _ := file.NewManager(dirName, blockSize)
	logManager, _ := log.NewManager(fileManager, "simpledb.log")

	return &SimpleDB{
		fileManager: fileManager,
		logManager:      logManager,
	}
}
