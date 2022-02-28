//С помощью пула воркеров написать программу, которая запускает 1000 горутин,
//каждая из которых увеличивает число на 1. Дождаться завершения всех горутин и убедиться,
//что при каждом запуске программы итоговое число равно 1000.

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var (
		jobs    = make(chan struct{}, 10)
		counter int64
	)

	for i := 1; i <= 1000; i++ {
		jobs <- struct{}{}

		go func() {
			defer func() {
				<-jobs
			}()
			atomic.AddInt64(&counter, 1)
		}()
	}

	// Ждем завершения горутин, проверяя длину канала.
	// Вместо WaitGroup, которая будет на следующих уроках.
	for len(jobs) > 0 {
		time.Sleep(1 * time.Second)
	}

	fmt.Println(counter)
}
