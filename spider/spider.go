package spider

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet(url string)  {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get err :",err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("body", body)

}
