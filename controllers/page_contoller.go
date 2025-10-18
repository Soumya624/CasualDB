package controllers

import (
	"errors"

	"casualdb.com/m/models"
)

type PageController struct {
	*models.Page
}

func NewPage(size int) *models.Page {
	return &models.Page{
		Bytes: make([]byte, size),
	}
}

func (pc *PageController) Read(offset int, temporaryDestination []byte) int {
	n := copy(temporaryDestination, pc.Bytes[offset:])

	return n
}

func (pc *PageController) Write(offset int, data []byte) (int, error) {
	if (offset + len(data)) > len(pc.Bytes) {
		return 0, errors.New("Unable to write in the page")
	}

	n := copy(pc.Bytes[offset:], data)

	return n, nil
}
