package io

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "account.txt"

func ReadFloatFromFile() (float64, error) {
	data, error := os.ReadFile(fileName)

	if error != nil {
		return 0, errors.New("failed to read the file")
	}

	value, error := strconv.ParseFloat(string(data), 64)

	if error != nil {
		return 0, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(value float64) {
	os.WriteFile(fileName, []byte(fmt.Sprint(value)), 0644)
}
