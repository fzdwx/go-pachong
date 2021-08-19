package test

import (
	"log"
	"sync"
	"testing"

	"go-pachong/pa"
)

var wg sync.WaitGroup

func TestPa(t *testing.T) {

	url := "https://github.com/"

	// 一直阻塞，没有调用wg.Done()  keeps blocking
	wg.Add(1)

	_ = pa.Go(url, func(url, body string) {
		log.Println(url)
	})
	wg.Wait()
}
