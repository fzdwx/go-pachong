# Go 爬虫
_巨几把耗流量_

能根据入口的url爬取页面中的解析出来的url路径，然后再次对解析出来的url继续爬取...

can crawl the parsed URL path in the page according to the URL of the entrance, and then continue to crawl the parsed
URL again...

## usage

调用pa.Go函数，第一个入参是入口url，传入的函数是每次爬取到的页面返回的页面数据，可以根据需要实现。

call `pa.Go(string,func(string,string))`, the first input parameter is the entry URL, and the function passed in is the page data returned by each
crawled page, which can be implemented as needed.

```go
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
```

## todo List

- [ ] todo...

## last
欢迎任何有助于项目的issue。

Any issues that help the project are welcome。
