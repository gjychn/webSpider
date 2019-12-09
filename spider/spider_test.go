package spider

import (
	"fmt"
	"testing"
)

func TestSpider(t *testing.T) {
	fmt.Println(hrefGet(htmlGet("http://news.baidu.com/")))
}
