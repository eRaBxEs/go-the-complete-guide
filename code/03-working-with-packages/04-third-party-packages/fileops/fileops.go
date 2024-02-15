package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) { // adjusted to make the file function generic
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000.0, errors.New("failed to read balance file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000.0, errors.New("failed to parse stored value")
	}
	return value, nil
}

func WriteFloatToFile(value float64, fileName string) { // adjusted to make the file function generic

	valueText := fmt.Sprint(value) // to be used to create the needed byte collections for the WriteFile function
	// os gives us a bunch of functions to interact with a file system
	os.WriteFile(fileName, []byte(valueText), 0644)
}
