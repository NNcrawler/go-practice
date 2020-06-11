package rangeoverroutine

import (
	"fmt"
	"sync"
)

func Run() {
	c := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(1)
	fwg:= sync.WaitGroup{}
	fwg.Add(2)
	go func() {
		wg.Wait()
		for ce := range c {
			fmt.Println("from 1-",ce)
		}
		fwg.Done()
	}()

	go func() {
		wg.Wait()
		for ce := range c {
			fmt.Println("from 2-",ce)
		}
		fwg.Done()
	}()

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	c <- 6
	c <- 7
	c <- 8
	close(c)
	wg.Done()
	fwg.Wait()
}
