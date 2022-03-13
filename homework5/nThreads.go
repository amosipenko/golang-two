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
		ch      = make(chan struct{}, n)
		counter int64
	)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			ch <- struct{}{}
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()

	close(ch)

	i := 0
	for range ch {
		i += 1
	}

	fmt.Printf("Number of passed threads (by channel): %d\n", i)
	fmt.Printf("Number of passed threads (by atomic): %d", counter)
}
