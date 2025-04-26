package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, strValue := range strings {
		floatVal, error := strconv.ParseFloat(strValue, 64)

		if error != nil {
			return nil, errors.New("failed to convert string to float")
		}

		floats = append(floats, floatVal)
	}

	return floats, nil
}
