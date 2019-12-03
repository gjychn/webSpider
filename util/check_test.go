package util

import "testing"

func TestCheckHttpUrl(t *testing.T) {
	ret := CheckHttpUrl("www.baidu.com")
	if ret != false {
		t.Error("TestCheckHttpUrl err")
	}

	ret = CheckHttpUrl("http://www.baidu.com")
	if ret != true {
		t.Error("TestCheckHttpUrl err")
	}
}