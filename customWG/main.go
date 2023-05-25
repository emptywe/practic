package main

import (
	"fmt"
	"sync/atomic"
)

type cWaitGroup struct {
	numWait atomic.Int32
}

func (wg *cWaitGroup) Add(delta int) {
	wg.numWait.Add(int32(delta))
}

func (wg *cWaitGroup) Done() {
	wg.numWait.Add(-1)
}

func (wg *cWaitGroup) Wait() {
	for {
		if wg.numWait.Load() == 0 {
			return
		}
	}
}

func main() {
	var wg cWaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
		wg.Wait()
	}
	fmt.Println("Done")

}
