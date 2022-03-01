package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := 1000000
	for i := 1; i <= n; i++ {
		func() {
			file, err := os.Create("C://EmptyFiles/File" + strconv.Itoa(i))
			if err != nil {
				err = fmt.Errorf("File not created. %s", err)
				fmt.Println(err)
				return
			}
			defer file.Close()
		}()
	}
}
