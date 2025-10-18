package controllers

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"casualdb.com/m/models"
)

type FileController struct {
	*models.FileManager
}

func NewFileManager(blockSize int, fileDirectory string) *models.FileManager {
	return &models.FileManager{
		BlockSize:     blockSize,
		FileDirectory: fileDirectory,
		FileList:      make(map[string]*os.File),
	}
}

func (fc *FileController) Read(block *models.Block, page *models.Page) (int, error) {
	fc.Mutex.Lock()
	defer fc.Mutex.Unlock()

	file, err := fc.GetFile(block.FileName)
	if err != nil {
		return 0, errors.New("Unable to get the file name!")
	}

	n, err := file.ReadAt(page.Bytes, int64(block.Identity*fc.BlockSize))
	if err != nil && err != io.EOF {
		return 0, errors.New("Unable to read from the block and write in the blank page!")
	}

	return n, nil
}

func (fc *FileController) Write(block *models.Block, page *models.Page) (int, error) {
	fc.Mutex.Lock()
	defer fc.Mutex.Unlock()

	file, err := fc.GetFile(block.FileName)
	if err != nil {
		return 0, errors.New("Unable to get the file name!")
	}

	n, err := file.WriteAt(page.Bytes, int64(block.Identity*fc.BlockSize))
	if err != nil && err != io.EOF {
		return 0, errors.New("Unable to write from the page to the block in the given offset!")
	}

	return n, nil
}

func (fc *FileController) Close() error {
	fc.Mutex.Lock()
	defer fc.Mutex.Unlock()

	for _, file := range fc.FileList {
		err := file.Close()
		if err != nil {
			return errors.New("Unable to close the file!")
		}
	}

	return nil
}

func (fc *FileController) GetFile(filename string) (*os.File, error) {
	file, ok := fc.FileList[filename]

	var err error

	if !ok {

		// This opens the file at the specified path with read and write permissions (os.O_RDWR)
		// Creates the file if it does not exist (os.O_CREATE)
		// Ensures that writes are synchronized to stable storage (os.O_SYNC)
		// 0666 sets the file permissions to be readable and writable by all users
		file, err = os.OpenFile(filepath.Join(fc.FileDirectory, filename), os.O_RDWR|os.O_CREATE|os.O_SYNC, 0777)
		if err != nil {
			return nil, fmt.Errorf("Failed to open file: %w", err)
		}

		fc.FileList[filename] = file
	}

	return file, nil
}
