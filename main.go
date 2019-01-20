package main

import (
	"strconv"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (v MyReader) Read([]byte) (int, error) {
	a, _ := strconv.Atoi("A")
	return a, nil
}

func main() {
	reader.Validate(MyReader{})
}
