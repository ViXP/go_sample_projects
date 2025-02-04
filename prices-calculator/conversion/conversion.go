package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	parsed := make([]float64, len(strings))

	for index, value := range strings {
		parsedValue, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return nil, errors.New("unparsable value in prices.txt")
		}

		parsed[index] = parsedValue
	}
	return parsed, nil
}
