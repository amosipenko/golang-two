//Напишите программу, которая запускает n потоков и дожидается завершения их всех

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		wg      = sync.WaitGroup{}
		n       = 1000
		counter int64
	)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()

	fmt.Printf("Number of passed threads: %d", counter)
}
