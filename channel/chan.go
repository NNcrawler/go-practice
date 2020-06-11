package channel

import (
	"fmt"
	"os"
	// "os"
	"time"
)

func Run() {
	c := make(chan int, 2)
	go func() { c <- 1 }()
	go func() { c <- 1 }()
	go func() { c <- 1 }()
	go func() {
		time.Sleep(2 * time.Second)
		close(c)
	}()

	// time.Sleep(1 * time.Second)
	for {
		select {
		case numb, more := <-c:
			if more {
				fmt.Println(numb)
			} else {
				os.Exit(0)
			}
		case <- time.After(1 * time.Second):
			fmt.Println("TIMEOUT!!")
			os.Exit(1)
		default:
			time.Sleep(300 * time.Millisecond)
			fmt.Println("waiting")
		}
	}
	// fmt.Println(<-c)
}
