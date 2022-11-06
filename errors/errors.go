package main

import (
	"errors"
	"strconv"
)

func errorGenerator() error {
	return errors.New("error")
}

func convertFromStrToFloat(str string) (float64, error) { // Erros are returned as the last value
	return strconv.ParseFloat(str, 64)
}

func main() {
	err := errorGenerator()
	if err != nil {
		panic(err)
	}
}
