package util

import (
	"bufio"
	"io"
	"os"
)

var (
	// File contains file utility functions
	File = fileUtil{}
)

type fileUtil struct{}

// ReadByLines reads a file and calls the handler for each line.
func (fu fileUtil) ReadByLines(filePath string, handler func(line string) error) error {
	var f *os.File
	var err error
	if f, err = os.Open(filePath); err == nil {
		defer f.Close()
		var line string
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line = scanner.Text()
			err = handler(line)
			if err != nil {
				return err
			}
		}
	}
	return err

}

// ReadByChunks reads a file in `chunkSize` pieces, dispatched to the handler.
func (fu fileUtil) ReadByChunks(filePath string, chunkSize int, handler func(line []byte) error) error {
	var f *os.File
	var err error
	if f, err = os.Open(filePath); err == nil {
		defer f.Close()

		chunk := make([]byte, chunkSize)
		for {
			readBytes, err := f.Read(chunk)
			if err == io.EOF {
				break
			}
			readData := chunk[:readBytes]
			err = handler(readData)
			if err != nil {
				return err
			}
		}
	}
	return err
}
