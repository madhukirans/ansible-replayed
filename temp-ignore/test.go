package main

import (
	"net/http"
	"strings"
	"fmt"
	"time"
	"sync"
)

var mutex sync.RWMutex


func main(){
	wg := sync.WaitGroup{}

	for i :=0; i< 2000; i++ {
		go func(){
			mutex.Lock()
			wg.Add(1)
			mutex.Unlock()
			body := strings.NewReader("Hi how are you")
			var netClient = &http.Client{
				Timeout: time.Second * 10,
			}
			res, err := netClient.Post("http://localhost:8080", "application/json", body)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println(res)
			}
			res, err = netClient.Get("http://localhost:8080")
			wg.Done()

		}()
	}
	mutex.Lock()
	wg.Wait()
	mutex.Unlock()
}
