package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			fmt.Printf("Worker %d starting...\n", i)
			time.Sleep(2 * time.Second)
			fmt.Printf("Worker %d done\n", i)
		})

	}

	wg.Wait()
}
