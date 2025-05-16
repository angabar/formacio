package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
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

	// defer nos permite ejecutar un bloque de codigo cuando la ejecucion de la
	// funcion que lo engloba termine
	defer file.Close()

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
		return nil, errors.New("an error ocurred reading the file content")
	}

	return lines, nil
}

func (fm *FileManager) WriteResult(data any) error {
	file, error := os.Create(fm.OutputFilePath)

	if error != nil {
		return errors.New("failed to create file")
	}

	defer file.Close()

	// Si hacemos que esepere 3 segundos en cada archivo el total seran 12
	// porque pasa por 4 archivos diferentes, un punto perfecto para usar
	// rutinas
	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	error = encoder.Encode(data)

	if error != nil {
		return errors.New("failed to encode json")
	}

	return nil
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
