package concurrency

import "sync"

var (
	sc     sync.Mutex
	Result int
)

func NumInsert(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 1000; i++ {
		sc.Lock()
		Result = i
		sc.Unlock()
	}
}

func NumInsertChannel(ch chan int, done chan bool) {
	defer close(ch)
	for i := 0; i <= 1000; i++ {
		ch <- i
		if i == 99 {
			break
		}
	}

	done <- true

}
