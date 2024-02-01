package main

import (
	"fmt"
	"time"

	"github.com/ezequielbugnon/leetcode/concurrency"
)

func main() {
	//send := []int{7, 9, 15, 2, 6}
	//response := sort.QuickSort(send)
	//response := sort.MergeSort(send)

	//fmt.Println(response)

	channel := make(chan int)
	done := make(chan bool)
	/*var wg sync.WaitGroup

	wg.Add(3)

	go concurrency.NumInsert(&wg)
	go concurrency.NumInsert(&wg)
	go concurrency.NumInsert(&wg)

	wg.Wait()

	fmt.Println(concurrency.Result)*/

	go concurrency.NumInsertChannel(channel, done)

	go func() {
		for {
			select {
			case value := <-channel:
				fmt.Println(value)
			case <-done:
				fmt.Println("goroutine done")
				return
			}
		}
	}()

	time.Sleep(time.Second * 2)

}
