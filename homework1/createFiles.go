package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numberOfFiles := 1000000
	err := createEmptyFiles(numberOfFiles)
	if err != nil {
		err = fmt.Errorf("Error in creating files. %s", err)
		fmt.Println(err)
	}
}

func createEmptyFiles(numberOfFiles int) error {
	for i := 1; i <= numberOfFiles; i++ {
		file, err := os.Create("C://EmptyFiles/File" + strconv.Itoa(i))
		if err != nil {
			return err
		}
		func() {
			defer file.Close()
		}()
	}

	return nil
}
