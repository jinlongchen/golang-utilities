package http

import (
	"io/ioutil"
	"net/http"
)

func GetRequestBody(r *http.Request) []byte {
	if r.Method == "POST" || r.Method == "PUT" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			println("ioutil.ReadAll(r.Body) err:", err.Error())
			return nil
		}
		return data
	} else {
		println("method:", r.Method)
	}
	return nil
}
