package input

import (
    "bufio"
    "os"
)

func GetFile() (*os.File, error) {
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func ReadInput(file *os.File) (*bufio.Scanner, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()
	scanner := bufio.NewScanner(file)

	buffer := make([]byte, fileSize)
	scanner.Buffer(buffer, int(fileSize))

	return bufio.NewScanner(file), nil
}
