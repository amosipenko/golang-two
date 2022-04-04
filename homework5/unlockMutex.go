//Реализуйте функцию для разблокировки мьютекса с помощью defer

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	var (
		mx sync.Mutex
		n  = 100
		wg sync.WaitGroup
	)
	// Из горутин будем писать счетчик в файл
	file, err := os.Create("temp.file")
	if err != nil {
		err = fmt.Errorf("File not created", err)
		fmt.Println(err)
		return
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			mx.Lock()
			defer mx.Unlock()
			defer wg.Done()

			file.Seek(0, 0)
			data, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatal(err)
				return
			}

			lastCounter, err := strconv.Atoi(string(data))
			if len(data) != 0 && err != nil {
				log.Fatal(err)
				return
			}

			err = os.WriteFile(file.Name(), []byte(strconv.Itoa(lastCounter+1)), 0644)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()

	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("Couldn`t get gorutines result", err)
		fmt.Println(err)
		return
	}

	fmt.Printf("Обработано горутин: %s", string(content))

	err = file.Close()
	if err != nil {
		err = fmt.Errorf("Couldn`t close the file", err)
		fmt.Println(err)
		return
	}

	err = os.Remove(file.Name())
	if err != nil {
		err = fmt.Errorf("Couldn`t remove the file", err)
		fmt.Println(err)
	}
}
