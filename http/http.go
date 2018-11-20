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
	"bytes"
	"io"
	"encoding/json"
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
func GetDataWithHeaders(reqURL string, reqHeader gohttp.Header) (gohttp.Header, []byte, error) {
	tr := &gohttp.Transport{
		TLSClientConfig: &tls.Config{},
	}
	cookieJar, _ := cookiejar.New(nil)

	client := &gohttp.Client{Transport: tr, Jar: cookieJar}
	request, _ := gohttp.NewRequest("GET", reqURL, nil)

	if reqHeader != nil {
		for key, value := range reqHeader {
			if len(value) > 0 {
				request.Header.Set(key, value[0])
			}
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}

	defer response.Body.Close()
	var body []byte
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(response.Body)
		if err != nil {
			return nil, nil, err
		}
		defer reader.Close()
		body, err = ioutil.ReadAll(reader)
		if err != nil {
			return nil, nil, err
		}
	default:
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, nil, err
		}
	}

	if response.StatusCode == 200 {
		return response.Header, body, nil
	} else {
		return response.Header, body, errors.WithCode(nil, fmt.Sprintf("HTTP_%d", response.StatusCode), response.Status)
	}
}
func GetJSON(url string, out interface{}) error {
	resp, err := gohttp.Get(url)
	if err != nil {
		return err
	}
	return readJSON(resp, out)
}

func PostData(reqURL string, bodyType string, data []byte) ([]byte, error) {
	timeout := time.Duration(15 * time.Second)
	client := &gohttp.Client{
		Timeout: timeout,
	}

	response, err := client.Post(reqURL, bodyType, bytes.NewReader(data))
	if err != nil {
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
func PostDataWithHeaders(reqURL string, reqHeader gohttp.Header, bodyType string, data []byte) (gohttp.Header, []byte, error) {
	tr := &gohttp.Transport{
		TLSClientConfig: &tls.Config{
		},
	}
	client := &gohttp.Client{
		Transport: tr,
		Timeout:   time.Duration(time.Second * 30),
	}

	request, _ := gohttp.NewRequest("POST", reqURL, bytes.NewReader(data))

	if reqHeader != nil {
		for key, value := range reqHeader {
			if len(value) > 0 {
				request.Header.Set(key, value[0])
			}
		}
	}
	request.Header.Set("Content-Type", bodyType)
	response, err := client.Do(request)

	if err != nil {
		return nil, nil, err
	}

	defer response.Body.Close()

	var body []byte

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(response.Body)
		if err != nil {
			return nil, nil, err
		}
		defer reader.Close()

		body, err = ioutil.ReadAll(reader)
		if err != nil {
			return nil, nil, err
		}
	default:
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, nil, err
		}
	}
	if response.StatusCode == 200 {
		return response.Header, body, nil
	} else {
		return response.Header, body, errors.WithCode(nil, fmt.Sprintf("HTTP_%d", response.StatusCode), response.Status)
	}
}

func readJSON(resp *gohttp.Response, out interface{}) (err error) {
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	if resp.StatusCode >= 400 {
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}

		return errors.WithCode(nil, fmt.Sprintf("HTTP_%d", resp.StatusCode), string(body))
	}

	if out == nil {
		io.Copy(ioutil.Discard, reader)
		return nil
	}

	decoder := json.NewDecoder(reader)
	return decoder.Decode(out)
}
