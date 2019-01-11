package http

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/jinlongchen/golang-utilities/errors"
	"github.com/jinlongchen/golang-utilities/log"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"time"
	gohttp "net/http"
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
func PostJSON(reqURL string, objToSend interface{}, out interface{}) error {
	jsonData, err := json.Marshal(objToSend)
	if err != nil {
		log.Errorf("marshal json err:%s", err.Error())
		return err
	}

	resp, err := gohttp.Post(reqURL, "application/json;charset=utf-8", bytes.NewReader(jsonData))
	if err != nil {
		log.Errorf("post json data to(%s)  err:%s", reqURL, err.Error())
		return err
	}
	return readJSON(resp, out)
}
func PostDataSsl(reqURL string, dataToSend, certPEMBlock, keyPEMBlock []byte) (respData []byte, err error) {
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return nil, err
	}

	tr := &gohttp.Transport{
		TLSClientConfig: &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true},
	}
	client := &gohttp.Client{Transport: tr}

	ret, err := client.Post(reqURL, "application/x-www-form-urlencoded", bytes.NewReader(dataToSend))
	if err != nil {
		return nil, err
	}
	defer func() {
		err = ret.Body.Close()
		log.Errorf(err.Error())
	}()

	data, err := ioutil.ReadAll(ret.Body)

	if err != nil {
		return nil, err
	}

	return data, err
}

func PostFiles(reqURL string, values map[string][]string, progressReporter func(r int64)) (ret []byte, err error) {
	var b ProgressReader
	b.Reporter = progressReporter

	w := multipart.NewWriter(&b)
	for key, files := range values {
		for _, file := range files {
			var fw io.Writer
			fileName := filepath.Base(file)
			if fw, err = w.CreateFormFile(key, fileName); err != nil {
				return
			}
			f, err := os.Open(file)
			if err != nil {
				return nil, err
			}

			if _, err = io.Copy(fw, f); err != nil {
				return nil, err
			}
		}
	}

	w.Close()

	client := &gohttp.Client{}
	request, _ := gohttp.NewRequest("POST", reqURL, &b)

	request.Header.Set("Content-Type", w.FormDataContentType())

	response, err := client.Do(request)
	if err != nil {
		return
	}

	defer response.Body.Close()

	var body []byte

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		var reader *gzip.Reader
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return
		}
		defer reader.Close()

		body, err = ioutil.ReadAll(reader)
		if err != nil {
			return
		}
	default:
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return
		}
	}

	if response.StatusCode == 200 {
		return body, nil
	} else {
		return body, errors.WithCode(nil, fmt.Sprintf("HTTP_%d", response.StatusCode), response.Status)
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
