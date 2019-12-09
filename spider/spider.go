package spider

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"webSpider/log"
)

func htmlGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get err :",err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error("")
	}
	//fmt.Println(string(body))
	return body
}

func hrefGet(data []byte) []string {
	reg := regexp.MustCompile(`href=\"[\/\w\s\d]+\"`)

	urls := reg.FindAllStringSubmatch(string(data), 1000)

	mp := make(map[string]bool)
	ret := make([]string, len(urls))
	for _,v := range urls  {
		for _, url := range v {
			if _, ok:= mp[url]; !ok {
				ret = append(ret, url)
				mp[url] = true
			}
		}
	}

	return ret
}
