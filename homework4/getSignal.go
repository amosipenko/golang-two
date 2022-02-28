//Написать программу, которая при получении в канал сигнала SIGTERM
//останавливается не позднее, чем за одну секунду (установить таймаут).

package main

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	var ch = make(chan os.Signal, 1)

	go func(ctx context.Context) {
		for {
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)

			signal := <-ch
			fmt.Println("Got signal: ", signal)
			if signal == syscall.SIGTERM {
				_, cancelFunc := context.WithTimeout(ctx, 1*time.Second)
				func() {
					defer cancelFunc()
				}()
			}

		}
	}(ctx)

	ch <- syscall.SIGTERM
	time.Sleep(3 * time.Second)
}
