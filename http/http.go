package http

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/proxy"
	"io"
	"mime/multipart"
	"net"
	netHttp "net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jinlongchen/golang-utilities/errors"
)

func GetData(reqURL string, reqHeader netHttp.Header, proxy string, timeoutSeconds int) ([]byte, error) {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{})
	if len(proxy) > 0 {
		if tr, err := setupProxyTransport(proxy); err == nil {
			client.SetTransport(tr)
		}
	}
	if timeoutSeconds > 0 {
		client.SetTimeout(time.Duration(timeoutSeconds) * time.Second)
	}

	r := client.R()
	if len(reqHeader) > 0 {
		r.SetHeaders(restyHeaderFromNetHeader(reqHeader))
	}

	resp, err := r.Get(reqURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == 200 {
		return resp.Body(), nil
	} else {
		return resp.Body(), errors.WithCode(nil, fmt.Sprintf("HTTP_%d", resp.StatusCode()), resp.Status())
	}
}

func GetJSON(url string, out interface{}) error {
	client := resty.New()
	client.SetTimeout(30 * time.Second)

	resp, err := client.R().Get(url)
	if err != nil {
		return err
	}
	return readJSON(resp, out)
}

func PutDataWithHeaders(reqURL string, reqHeader netHttp.Header, contentType string, data []byte) (netHttp.Header, []byte, error) {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{})
	client.SetTimeout(30 * time.Second)

	resp, err := client.R().
		SetHeaders(restyHeaderFromNetHeader(reqHeader)).
		SetHeader("Content-Type", contentType).
		SetBody(data).
		Put(reqURL)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode() == 200 {
		return resp.Header(), resp.Body(), nil
	} else {
		return resp.Header(), resp.Body(), errors.WithCode(nil, fmt.Sprintf("HTTP_%d", resp.StatusCode()), resp.Status())
	}
}

func PostData(reqURL string, reqHeader netHttp.Header, contentType string, proxy string, timeoutSeconds int, data []byte) ([]byte, error) {
	client := resty.New()

	if len(proxy) > 0 {
		if tr, err := setupProxyTransport(proxy); err == nil {
			client.SetTransport(tr)
		}
	}
	if timeoutSeconds > 0 {
		client.SetTimeout(time.Duration(timeoutSeconds) * time.Second)
	}

	r := client.R()
	if len(reqHeader) > 0 {
		r.SetHeaders(restyHeaderFromNetHeader(reqHeader))
	}

	resp, err := r.
		SetHeader("Content-Type", contentType).
		SetBody(data).
		Post(reqURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == 200 {
		return resp.Body(), nil
	} else {
		return resp.Body(), errors.WithCode(nil, fmt.Sprintf("HTTP_%d", resp.StatusCode()), resp.Status())
	}
}

func PostDataWithHeaders(reqURL string, reqHeader netHttp.Header, contentType string, data []byte, timeout time.Duration) (netHttp.Header, []byte, error) {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{})
	client.SetTimeout(timeout)

	resp, err := client.R().
		SetHeaders(restyHeaderFromNetHeader(reqHeader)).
		SetHeader("Content-Type", contentType).
		SetBody(data).
		Post(reqURL)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode() == 200 {
		return resp.Header(), resp.Body(), nil
	} else {
		return resp.Header(), resp.Body(), errors.WithCode(nil, fmt.Sprintf("HTTP_%d", resp.StatusCode()), resp.Status())
	}
}

func PostJSON(reqURL string, objToSend interface{}, out interface{}) error {
	client := resty.New()
	client.SetTimeout(30 * time.Second)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetBody(objToSend).
		Post(reqURL)
	if err != nil {
		return err
	}
	return readJSON(resp, out)
}

func PostDataSsl(reqURL string, dataToSend, certPEMBlock, keyPEMBlock []byte) (respData []byte, err error) {
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return nil, err
	}

	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true})

	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody(dataToSend).
		Post(reqURL)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}

func PostXml(reqURL string, xmlToSend string, objReceived interface{}) (respData []byte, err error) {
	client := resty.New()
	client.SetTimeout(30 * time.Second)

	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf-8").
		SetBody(xmlToSend).
		Post(reqURL)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(resp.Body(), objReceived)
	return resp.Body(), err
}

func PostFiles(reqURL string, filesVal map[string][]string, form map[string]string, progressReporter func(r int64)) (ret []byte, err error) {
	var b ProgressReader
	b.Reporter = progressReporter

	w := multipart.NewWriter(&b)
	for key, files := range filesVal {
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
	for key, val := range form {
		if fw, err := w.CreateFormField(key); err == nil {
			if _, err = fw.Write([]byte(val)); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	w.Close()

	client := resty.New()
	client.SetTimeout(30 * time.Second)

	resp, err := client.R().
		SetHeader("Content-Type", w.FormDataContentType()).
		SetBody(&b).
		Post(reqURL)
	if err != nil {
		return
	}

	if resp.StatusCode() == 200 {
		return resp.Body(), nil
	} else {
		return resp.Body(), errors.WithCode(nil, fmt.Sprintf("HTTP_%d", resp.StatusCode()), resp.Status())
	}
}

func DownloadFile(reqURL string, filePath string) error {
	client := resty.New()
	client.SetTimeout(30 * time.Second)

	resp, err := client.R().Get(reqURL)
	if err != nil {
		return err
	}

	dir := path.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	switch resp.Header().Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(bytes.NewReader(resp.Body()))
		if err != nil {
			return err
		}
		defer reader.Close()

		writer, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		defer writer.Close()

		_, err = io.Copy(writer, reader)
		if err != nil {
			return err
		}
	default:
		writer, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		defer writer.Close()

		_, err = io.Copy(writer, bytes.NewReader(resp.Body()))
		if err != nil {
			return err
		}
	}
	if resp.StatusCode() == 200 {
		return nil
	} else {
		return errors.WithCode(nil, fmt.Sprintf("HTTP_%d", resp.StatusCode()), resp.Status())
	}
}

func DeleteDataWithHeaders(reqURL string, reqHeader netHttp.Header) (netHttp.Header, []byte, error) {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{})
	client.SetTimeout(30 * time.Second)

	resp, err := client.R().SetHeaders(restyHeaderFromNetHeader(reqHeader)).Delete(reqURL)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode() == 200 {
		return resp.Header(), resp.Body(), nil
	} else {
		return resp.Header(), resp.Body(), errors.WithCode(nil, fmt.Sprintf("HTTP_%d", resp.StatusCode()), resp.Status())
	}
}

func readJSON(resp *resty.Response, out interface{}) (err error) {
	var reader io.ReadCloser
	switch resp.Header().Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(bytes.NewReader(resp.Body()))
		defer reader.Close()
	default:
		reader = io.NopCloser(bytes.NewReader(resp.Body()))
	}

	if resp.StatusCode() >= 400 {
		body, err := io.ReadAll(reader)
		if err != nil {
			return err
		}

		return errors.WithCode(nil, fmt.Sprintf("HTTP_%d", resp.StatusCode()), string(body))
	}

	if out == nil {
		io.Copy(io.Discard, reader)
		return nil
	}

	decoder := json.NewDecoder(reader)
	return decoder.Decode(out)
}

func restyHeaderFromNetHeader(reqHeader netHttp.Header) map[string]string {
	ret := make(map[string]string)
	for key, values := range reqHeader {
		if len(values) > 0 {
			ret[key] = values[0]
		}
	}
	return ret
}
func setupProxyTransport(proxyAddr string) (*netHttp.Transport, error) {
	if strings.HasPrefix(proxyAddr, "socks5://") {
		socks5Addr := strings.TrimPrefix(proxyAddr, "socks5://")

		dialer, err := proxy.SOCKS5("tcp", socks5Addr, nil, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("failed to create SOCKS5 dialer: %v", err)
		}

		return &netHttp.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.Dial(network, addr)
			},
		}, nil

	} else if strings.HasPrefix(proxyAddr, "http://") || strings.HasPrefix(proxyAddr, "https://") {
		proxyURL, err := url.Parse(proxyAddr)
		if err != nil {
			return nil, fmt.Errorf("invalid HTTP proxy URL: %v", err)
		}

		return &netHttp.Transport{
			Proxy: netHttp.ProxyURL(proxyURL),
		}, nil
	}

	return nil, fmt.Errorf("unsupported proxy type: %s", proxyAddr)
}
