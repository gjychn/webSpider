package download

import (
	"io/ioutil"
	"net/http"
	"webSpider/log"
)

type Download struct {
	fileName string
	sourceUrl string
}

func (d* Download)DownloadByHttp(url string) error {
	client := http.Client{}

	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		log.Error("http NewRequest err : %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error("http resp err : ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("io read body err: ", err)
	}

	d.writeFile(body)

	return nil
}


func (d * Download)writeFile(body []byte) {
	err := ioutil.WriteFile(d.fileName, body, 0666)

	if err != nil {
		log.Error("ioutil.WriteFile err : ", err)
	}

}