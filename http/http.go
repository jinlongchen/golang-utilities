package http

import (
	"crypto/tls"
	"github.com/jinlongchen/golang-utilities/log"
	gohttp "net/http"
	"net/http/cookiejar"
	"time"
	"compress/gzip"
	"io/ioutil"
	"github.com/jinlongchen/golang-utilities/errors"
	"fmt"
)

func GetData(reqURL string) ([]byte, error) {
	tr := &gohttp.Transport{
		TLSClientConfig: &tls.Config{},
	}
	cookieJar, _ := cookiejar.New(nil)

	client := &gohttp.Client{Transport: tr, Jar: cookieJar, Timeout: time.Duration(time.Second * 30)}
	request, _ := gohttp.NewRequest("GET", reqURL, nil)

	response, err := client.Do(request)

	if err != nil {
		log.Errorf("Get data error:%s", err.Error())
		return nil, err
	}

	defer response.Body.Close()

	var body []byte

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(response.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()

		body, err = ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
	default:
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
	}

	if response.StatusCode == 200 {
		return body, nil
	} else {
		return body, errors.WithCode(nil, fmt.Sprintf("HTTP_%d", response.StatusCode), response.Status)
	}
}
