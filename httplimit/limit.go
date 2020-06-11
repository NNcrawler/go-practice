package httplimit

import (
	"fmt"
	"net/http"
	"time"
)

func Run() {
	// c := make(chan int, 1)
	// go func() {
	// 	for {
	// 		time.Sleep(3 * time.Second)
	// 		<-c
	// 	}
	// }()
	handler := func(w http.ResponseWriter, r *http.Request) {
		// c <- 1
		time.Sleep(1 * time.Second)
		w.Write([]byte("hello world"))
	}
	http.HandleFunc("/sleep", handler)

	err := http.ListenAndServe(":1121", nil)
	fmt.Println("called", err)
}
