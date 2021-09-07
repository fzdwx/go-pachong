# 流量粉碎机

* [流量粉碎机](#流量粉碎机)
    * [usage](#usage)
    * [todo List](#todo-List)
    * [last](#last)


_巨几把耗流量_

能根据入口的url爬取页面中的解析出来的url路径，然后再次对解析出来的url继续爬取...

can crawl the parsed URL path in the page according to the URL of the entrance, and then continue to crawl the parsed
URL again...

## usage

调用`pa.NewPa(string)`函数，第一个入参是入口url，然后调用`AddCallback(func(string,string))`传入的函数是每次爬取到的页面返回的页面数据，可以根据需要实现。最后调用`Go()`开始爬取

call `pa.NewPa(string)`, the first input parameter is the entry URL, and the call `AddCallback(func(string,string))`
function passed in is the page data returned by each crawled page, which can be implemented as needed,last call `Go()`

```go
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
```
## todo List

- [ ] todo...

## last

欢迎任何有助于项目的issue。

Any issues that help the project are welcome。
