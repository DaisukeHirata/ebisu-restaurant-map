package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpGet(url string) ([]byte, error) {
	response, _ := http.Get(url)
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func Unescape(body string) string {
	parseBody, _ := url.Parse(body)
	return parseBody.Path
}
