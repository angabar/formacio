package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("Error reading file")
	}

	textValue := string(content)
	value, err := strconv.ParseFloat(textValue, 64)

	if err != nil {
		return 0, errors.New("Error converting the string")
	}

	return value, nil
}

func WriteBalanceToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(valueText), 0644)

	if err != nil {
		fmt.Println("Error creating the file")
	}
}
