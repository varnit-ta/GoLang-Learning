package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
WaitGroup is used to wait for the program to finish all goroutines
Mutex is used to lock and unlock the code
*/

var signals = []string{"test"}

var wg sync.WaitGroup //usually passed as a pointer
var mut sync.Mutex    //used to lock and unlock the code

func main() {
	//goroutine - a lightweight thread managed by the Go runtime
	//go keyword is used to create a goroutine
	weblist := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://flipkart.com",
		"http://youtube.com",
	}

	for _, web := range weblist {
		go getStatusCode(web)

		// wg.Add(1) is used to add a goroutine to the waitgroup
		wg.Add(1)
	}

	//wg.Wait() is used to wait for all goroutines to finish
	wg.Wait()
}

func getStatusCode(endpoint string) {
	//defer is used to decrement the waitgroup counter
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println(endpoint, "is down")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%s -> %d\n", endpoint, res.StatusCode)
	}
}
