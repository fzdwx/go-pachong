package test

import (
	"sync"
	"testing"

	"go-pachong/pa"
)

var wg sync.WaitGroup

func TestPa(t *testing.T) {

	url := "https://github.com/"

	// 一直阻塞，没有调用wg.Done()  keeps blocking
	wg.Add(1)

	_ = pa.NewPa(url).AddCallback(func(url, body string) {
		// nothing
	}).Go()
	wg.Wait()
}
