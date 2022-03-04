package main

import (
	"fmt"
	"runtime"
	"sync"
)


///C race трассировкой что-то пошло не так, трассировка не собралась
////Получил ошибку 2022/03/04 17:32:24 Parsing trace...
//failed to parse trace: failed to read header: read 0, err EOF
func main(){
	var (
		wg sync.WaitGroup
		counter int
	)
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			counter++
		}()
		if i%100 == 0 {
			runtime.Gosched()
		}
	}

	wg.Wait()
	fmt.Println(counter)
}

///==================
//WARNING: DATA RACE
//Read at 0x00c000130010 by goroutine 8:
//  main.main.func1()
//      /Users/ikroni/Documents/study/lesson6/race/race.go:21 +0x58
//
//Previous write at 0x00c000130010 by goroutine 7:
//  main.main.func1()
//      /Users/ikroni/Documents/study/lesson6/race/race.go:21 +0x6c
//
//Goroutine 8 (running) created at:
//  main.main()
//      /Users/ikroni/Documents/study/lesson6/race/race.go:19 +0xc0
//
//Goroutine 7 (finished) created at:
//  main.main()
//      /Users/ikroni/Documents/study/lesson6/race/race.go:19 +0xc0
//==================
//1000
//Found 1 data race(s)
//exit status 66

///Принудительный вызов шедуллера работает:
//ikroni@Aleksandrs-MacBook-Pro race % GODEBUG=schedtrace=1000 go run race.go
//SCHED 0ms: gomaxprocs=8 idleprocs=7 threads=3 spinningthreads=0 idlethreads=0 runqueue=0 [0 0 0 0 0 0 0 0]
//# command-line-arguments
//SCHED 0ms: gomaxprocs=8 idleprocs=6 threads=3 spinningthreads=1 idlethreads=0 runqueue=0 [1 0 0 0 0 0 0 0]
//SCHED 0ms: gomaxprocs=8 idleprocs=7 threads=2 spinningthreads=0 idlethreads=0 runqueue=0 [0 0 0 0 0 0 0 0]
//932
