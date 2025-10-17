package server

import (
	"simple_db/file"
	"simple_db/log"
	"simple_db/buffer"
)

type SimpleDB struct {
	fileManager *file.Manager
	logManager  *log.Manager
	bufferManager *buffer.Manager
}

func NewSimpleDB(dirName string, blockSize int32, buffSize int32) *SimpleDB {
	fileManager, _ := file.NewManager(dirName, blockSize)
	logManager, _ := log.NewManager(fileManager, "simple_db.log")
	bufferManager := buffer.NewManager(fileManager, logManager, buffSize)


	return &SimpleDB{
		fileManager:  fileManager,
		logManager:   logManager,
		bufferManager: bufferManager,
	}
}
