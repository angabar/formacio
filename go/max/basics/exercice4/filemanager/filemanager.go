package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm *FileManager) ReadLines() ([]string, error) {
	file, error := os.Open(fm.InputFilePath)

	if error != nil {
		return nil, errors.New("an error ocurred opening the file")
	}

	// NewScanner nos permite leer el documento
	scanner := bufio.NewScanner(file)

	var lines []string

	// Scan devuelve linea a linea lo que va leyendo del archivo hasta terminar
	// o hasta que suceda un error
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Para captar el error de un scanner tenemos que invocar el metodo Err, de
	// lo contrario no podremos detectarlo y dar un feedback
	error = scanner.Err()

	if error != nil {
		file.Close()
		return nil, errors.New("an error ocurred reading the file content")
	}

	file.Close()
	return lines, nil
}

func (fm *FileManager) WriteResult(data any) error {
	file, error := os.Create(fm.OutputFilePath)

	if error != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	error = encoder.Encode(data)

	if error != nil {
		file.Close()
		return errors.New("failed to encode json")
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
