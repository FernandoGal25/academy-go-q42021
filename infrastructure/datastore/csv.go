package datastore

import (
	"encoding/csv"
	"errors"
	"os"
)

type CSVHandler struct {
	Path   string
	File   *os.File
	Writer *csv.Writer
	Reader *csv.Reader
}

func NewCSVHandler(filePath string) *CSVHandler {
	return &CSVHandler{Path: filePath}
}

func (h *CSVHandler) BuildHandler() error {
	f, err := os.Open(h.Path)
	if err != nil {
		return errors.New("no se pudo abrir el archivo")
	}

	h.File = f
	h.Reader = csv.NewReader(f)
	h.Writer = csv.NewWriter(f)

	return nil
}

func (h *CSVHandler) Read() ([]string, error) {
	return h.Reader.Read()
}

func (h *CSVHandler) Write() ([]string, error) {
	return h.Reader.Read()
}
