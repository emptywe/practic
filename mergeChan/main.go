package main

import (
	"fmt"
	"time"
)

// joinChannels1 join info from channels if they will close
func joinChannels1(a, b, c chan int) chan int {
	num := make(chan int)
	go func() {
		for {
			aa, ok1 := <-a
			bb, ok2 := <-b
			cc, ok3 := <-c
			if !ok1 && !ok2 && !ok3 {
				close(num)
				return
			}
			num <- aa + bb + cc
		}
	}()
	return num
}

// joinChannels2 join info from channels until all of them stop responding for 20ms
func joinChannels2(a, b, c chan int) chan int {
	num := make(chan int)
	go func() {

		for {
			var timeOutA, timeOutB, timeOutC bool
			res := 0
			timeOut := time.NewTicker(time.Millisecond * 30)
			select {
			case vala := <-a:
				res += vala
				timeOut = time.NewTicker(time.Millisecond * 30)
			case <-timeOut.C:
				timeOutA = true
				break
			}

			select {
			case vala := <-b:
				res += vala
				timeOut = time.NewTicker(time.Millisecond * 30)
			case <-timeOut.C:
				timeOutB = true
				break
			}

			select {
			case vala := <-c:
				res += vala
				timeOut = time.NewTicker(time.Millisecond * 30)
			case <-timeOut.C:
				timeOutC = true
				break
			}
			if timeOutA && timeOutB && timeOutC {
				close(a)
				close(b)
				close(c)
				close(num)
				return
			}
			num <- res
		}
	}()
	return num
}

func main() {
	var (
		a = make(chan int)
		b = make(chan int)
		c = make(chan int)
	)

	go func() {
		for i := 1; i < 10; i++ {
			// some fast channel
			time.Sleep(time.Millisecond * 2)
			a <- i
		}
		// some last sent data
		a <- 1000
		// uncomment for joinChannels1
		//close(a)
	}()

	go func() {
		for i := 10; i < 100; i += 10 {
			// some normal channel
			time.Sleep(time.Millisecond * 5)
			b <- i
		}
		// uncomment for joinChannels1
		//close(b)
	}()

	go func() {
		for i := 100; i < 1000; i += 100 {
			// some slow channel
			time.Sleep(time.Millisecond * 10)
			c <- i
		}
		// uncomment for joinChannels1
		//close(c)
	}()

	//for num := range joinChannels1(a, b, c) {
	//	fmt.Println(num)
	//}

	for num := range joinChannels2(a, b, c) {
		fmt.Println(num)
	}

}
