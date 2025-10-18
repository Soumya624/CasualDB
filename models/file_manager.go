package models

import (
	"os"
	"sync"
)

type FileManager struct {
	BlockSize     int
	FileDirectory string
	FileList      map[string]*os.File

	Mutex sync.Mutex
}
