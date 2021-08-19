package pa

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var (
	re              = regexp.MustCompile(`href="(?s:(.*?))"`)
	excludeSuffixes = []string{"png", "ico"}
)

// Go 开始爬取，传入一个匿名函数，入参是每个页面的数据
func Go(url string, f func(url, body string)) error {
	if f == nil {
		panic("callback function must be not null！")
	}

	body, err := get(url)
	if err != nil {
		return err
	}

	do(body, f)

	return nil
}

func do(body string, f func(url, body string)) {
	matched := re.FindAllString(body, -1)
	for _, match := range matched {
		url := processRawUrl(match)

		if url == "" || isExcludeUrl(url, excludeSuffixes) {
			continue
		}

		url = addPrefix(url)
		go func() {
			body, err := get(url)
			if err != nil {
				log.Println(err)
				return
			}
			f(url, body)
			do(body, f)
		}()
	}
}

// processRawUrl 返回不带https前缀的url
func processRawUrl(match string) string {
	var url string
	split := strings.Split(match, "//")

	for i, s := range split {
		if i == 0 {
			continue
		}
		url = url + s
	}

	if strings.HasSuffix(url, `"`) {
		url = url[0:strings.LastIndex(url, `"`)]
	}

	return url
}

func addPrefix(url string) string {
	return "https://" + url
}

// get 发送get请求，返回网页的body
func get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	return string(bytes), err
}

// isExcludeUrl 判断url的后缀是否有我们要排除的
func isExcludeUrl(url string, excludeSuffixes []string) bool {
	if len(excludeSuffixes) == 0 || cap(excludeSuffixes) == 0 {
		return false
	}
	for _, suffix := range excludeSuffixes {
		if strings.HasSuffix(url, suffix) {
			return true
		}
	}
	return false
}
