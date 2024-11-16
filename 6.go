package main

import (
	"fmt"
	"sync"
	"time"
)

var M = 2000
var N = 10
var i = 0
var mutex sync.Mutex
var wg sync.WaitGroup
var shi = 0

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go move()
	}
	wg.Wait()
	time.Sleep(time.Second * 5)
}
func move() {

	for {
		mutex.Lock()
		i++
		if M > 0 {
			M -= 200
			shi += 200
			if M > 0 {
				fmt.Printf("\r%d", shi)
				time.Sleep(time.Millisecond * 100)

			} else {
				fmt.Printf("\r%s\n", "ALL MOVED")
			}
			fmt.Println(" the car --", i, "Moved", 200)

			wg.Done()
		} else {
			break
		}
		mutex.Unlock()
	}
}
