package pa

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var (
	defaultExcludeSuffixes = []string{"png", "ico"}
	re                     = regexp.MustCompile(`href="(?s:(.*?))"`)
	defaultCallback        = func(url string, body string) {
		log.Println(url)
	}
)

// Pa 当前爬虫的主体
type Pa struct {
	mainUrl         string                 // 入口函数
	callback        func(url, body string) // 回调，每次爬取到一个页面就调用
	excludeSuffixes []string
}

func NewPa(mainUrl string) *Pa {
	return &Pa{
		mainUrl:         mainUrl,
		callback:        defaultCallback,
		excludeSuffixes: defaultExcludeSuffixes,
	}
}

// AddCallback 添加回调函数，每次爬取到一个页面就调用
func (p *Pa) AddCallback(f func(url, body string)) *Pa {
	p.callback = f
	return p
}

// AddExcludeSuffixes 添加网址后缀排除
func (p *Pa) AddExcludeSuffixes(excludeSuffixes []string) *Pa {
	p.excludeSuffixes = excludeSuffixes
	return p
}

// Go 开始爬取，传入一个匿名函数，入参是每个页面的数据
func (p *Pa) Go() error {

	body, err := get(p.mainUrl)
	if err != nil {
		return err
	}

	p.do(body)

	return nil
}

func (p *Pa) do(body string) {
	matched := re.FindAllString(body, -1)
	for _, match := range matched {
		url := processRawUrl(match)

		if url == "" || isExcludeUrl(url, p.excludeSuffixes) {
			continue
		}

		url = addPrefix(url)

		// new thread
		go func() {
			body, err := get(url)
			if err != nil {
				log.Println(err)
				return
			}

			// callback
			p.callback(url, body)
			// loop
			p.do(body)
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
	resp.Body.Close()
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
