package datastore

import (
	"encoding/csv"
	"log"
	"os"

	customErrors "github.com/FernandoGal25/academy-go-q42021/error"
)

// CSVHandler is a wrapper of the CSV file, combines the management of the
// os and enconding/csv libraries, reads and write from existing CSV files.
type CSVHandler struct {
	path      string
	Schema    []string
	readFile  *os.File
	writeFile *os.File
	writer    *csv.Writer
	reader    *csv.Reader
}

// NewCSVHandler returns an instance of the CSVHandler
func NewCSVHandler(filePath string) *CSVHandler {
	return &CSVHandler{path: filePath}
}

// BuildHandler is a method used to assign the handler with the required tools to manage
// the csv file, is separated of the initialization of the handler in
// order to avoid leaving the CSV file opened in case something fails.

func (h *CSVHandler) BuildHandler() error {
	wf, err := os.OpenFile(h.path, os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return customErrors.ErrCSVFormat{Message: "Cannot open CSV file to write", Err: err}
	}

	rf, err := os.Open(h.path)

	if err != nil {
		return customErrors.ErrCSVFormat{Message: "Cannot open CSV file to read", Err: err}
	}

	h.readFile = rf
	h.writeFile = wf
	h.reader = csv.NewReader(rf)
	h.writer = csv.NewWriter(wf)
	header, err := h.reader.Read()

	if err != nil {
		return customErrors.ErrCSVFormat{Message: "Failed to read CSV header", Err: err}
	}

	h.Schema = header

	return nil
}

// Read is a method that reads one line of the CSV file.
func (h *CSVHandler) Read() ([]string, error) {
	record, err := h.reader.Read()
	if err != nil {
		return nil, customErrors.ErrCSVFormat{Message: "Cannot read CSV", Err: err}
	}

	return record, nil
}

// ReadAll is a method that reads all the lines of the CSV file.
func (h *CSVHandler) ReadAll() ([][]string, error) {
	records, err := h.reader.ReadAll()
	if err != nil {
		return nil, customErrors.ErrCSVFormat{Message: "Cannot read CSV", Err: err}
	}

	return records, nil
}

// Write is a method that writes one row of the csv.
func (h *CSVHandler) Write(r []string) error {
	if err := h.writer.Write(r); err != nil {
		return customErrors.ErrCSVFormat{Message: "Cannot write in CSV", Err: err}
	}

	return nil
}

// Close is a method that wraps the os.File method Close, closes the stream
// of the opened CSV file and flushes the write stream in order to
// persist the changes.
func (h *CSVHandler) Close() {
	h.writer.Flush()

	if err := h.writeFile.Close(); err != nil {
		log.Fatalln(customErrors.ErrCSVFormat{Message: "Cannot close write file", Err: err})
	}

	if err := h.readFile.Close(); err != nil {
		log.Fatalln(customErrors.ErrCSVFormat{Message: "Cannot close read file", Err: err})
	}
}
