package chart

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
func (fu fileUtil) ReadByLines(filePath string, handler func(line string)) error {
	if f, err := os.Open(filePath); err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			handler(line)
		}
	} else {
		return err
	}
	return nil
}

// ReadByChunks reads a file in `chunkSize` pieces, dispatched to the handler.
func (fu fileUtil) ReadByChunks(filePath string, chunkSize int, handler func(line []byte)) error {
	if f, err := os.Open(filePath); err == nil {
		defer f.Close()

		chunk := make([]byte, chunkSize)
		for {
			readBytes, err := f.Read(chunk)
			if err == io.EOF {
				break
			}
			readData := chunk[:readBytes]
			handler(readData)
		}
	} else {
		return err
	}
	return nil
}
