package datastore

import (
	"encoding/csv"
	"errors"
	"os"
)

/*
	Wrapper of the CSV file, combines the management
	of the os and enconding/csv libraries.
*/
type CSVHandler struct {
	path   string
	schema []string
	file   *os.File
	writer *csv.Writer
	reader *csv.Reader
}

/*
	Returns an instance of the CSVHandler
*/
func NewCSVHandler(filePath string) *CSVHandler {
	return &CSVHandler{path: filePath}
}

/*
	Method used to assign the handler with the required
	tools to manage the csv file, is separated of the
	initialization of the handler in order to avoid leaving
	the CSV file opened in case something fails.
*/
func (h *CSVHandler) BuildHandler() error {
	f, err := os.Open(h.path)

	if err != nil {
		return errors.New("CANNOT OPEN CSV FILE")
	}

	h.file = f
	h.reader = csv.NewReader(f)
	h.writer = csv.NewWriter(f)
	header, err := h.Read()
	if err != nil {
		return errors.New("CANNOT READ CSV FILE")
	}

	h.schema = header
	return nil
}

/*
	Method that reads one line of the CSV file.
*/
func (h *CSVHandler) Read() ([]string, error) {
	return h.reader.Read()
}

/*
	Method that reads all the lines of the CSV file.
*/
func (h *CSVHandler) ReadAll() ([][]string, error) {
	return h.reader.ReadAll()
}

/*
	Method that wraps the os.File method Close,
	closes the stream of the opened CSV file.
*/
func (h *CSVHandler) Close() error {
	return h.file.Close()
}
