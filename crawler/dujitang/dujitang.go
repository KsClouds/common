package dujitang

import (
	"io/ioutil"
	"net/http"
)

func Get() (rsp string) {
	response, err := http.Get("https://du.shadiao.app/api.php")
	if err == nil {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err == nil {
			rsp = string(body)
		}
	}
	return
}
