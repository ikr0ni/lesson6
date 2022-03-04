package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	fmt.Println("123")
	trace.Start(os.Stderr)
	defer trace.Stop()

	var (
		counter int
		lock  sync.Mutex
		wg sync.WaitGroup
	)

	wg.Add(6)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			counter++
		}()
	}

	for i := 0; i < 4; i++ {
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			fmt.Println(counter)
		}()
	}

	wg.Wait()

}
