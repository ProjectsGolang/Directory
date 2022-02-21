package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	year := t.Year()
	month := t.Month()
	day := t.Day()
	directory := fmt.Sprintf("./assets/%d/%02d/%d", year, month, day)
	createDirectoryIfDoesNotExist(directory)
}

func createDirectoryIfDoesNotExist(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			panic(err)
		}
	}
}
