package main

import (
	"log"
	"sync"

	"go-pachong/pa"
)

var wg sync.WaitGroup

func main() {

	url := "https://github.com/"

	// 一直阻塞，没有调用wg.Done()  keeps blocking
	wg.Add(1)

	_ = pa.Go(url, func(url, body string) {
		log.Println(url)

		// if url== "xxx" {
		// 	wg.Done()
		// }
	})

	wg.Wait()
}
