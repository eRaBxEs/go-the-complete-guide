package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find file.")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balanve value.")
	}
	return value, nil
}

func WriteFloatToFile(balance float64, fileName string) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)

}
