package utils

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	//代理
	HTTP_PROXY_ENABLE = false                   //是否启用代理
	HTTP_PROXY_ADDR   = "http://127.0.0.1:9999" //代理地址

	//超时时间
	HTTP_TIMEOUT = 30

	//重试
	HTTP_RETRY_ENABLE   = true
	HTTP_RETRY_TIMES    = 5 //重试次数
	HTTP_RETRY_INTERVAL = 5 //重试间隔,单位秒

	//成功后延迟
	HTTP_SUCCESS_DELAY_ENABLE   = true
	HTTP_SUCCESS_DELAY_INTERVAL = 2 //成功后延迟时间
)

func getHttpClient() *http.Client {
	client := &http.Client{Timeout: time.Duration(HTTP_TIMEOUT) * time.Second}
	if HTTP_PROXY_ENABLE {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(HTTP_PROXY_ADDR)
		}
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				Proxy: proxy,
			},
		}
	}
	return client
}

func retryHttpRequest(req *http.Request) (*http.Response, error) {
	if HTTP_RETRY_ENABLE {
		// 读取req中的body
		var body []byte
		var err error
		if req.Body != nil {
			body, err = ioutil.ReadAll(req.Body)
		} else {
			body = nil
			err = nil
		}
		if err != nil {
			return nil, err
		}
		var resp *http.Response
		for i := 0; i < HTTP_RETRY_TIMES; i++ {
			//把刚刚读出来的body再写进去
			if body != nil {
				req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			} else {
				req.Body = nil
			}
			client := getHttpClient()
			resp, err = client.Do(req)
			if err == nil {
				//成功延迟
				if HTTP_SUCCESS_DELAY_ENABLE {
					time.Sleep(time.Duration(HTTP_SUCCESS_DELAY_INTERVAL) * time.Second)
				}
				return resp, err
			}
			time.Sleep(time.Duration(HTTP_RETRY_INTERVAL) * time.Second)
		}
		return resp, err
	} else {
		client := getHttpClient()
		resp, err := client.Do(req)
		if err == nil {
			//成功延迟
			if HTTP_SUCCESS_DELAY_ENABLE {
				time.Sleep(time.Duration(HTTP_SUCCESS_DELAY_INTERVAL) * time.Second)
			}
		}
		return resp, err
	}
}

func HttpPostJson(url string, headers map[string]string, jsonData []byte) (int, []*http.Cookie, []byte, error) {
	//构造请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	//发送请求
	resp, err := retryHttpRequest(req)
	if err != nil {
		return 0, nil, nil, err
	}
	defer resp.Body.Close()

	//处理请求
	statusCode := resp.StatusCode
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, nil, err
	}
	cookies := resp.Cookies()
	return statusCode, cookies, result, nil
}

func HttpPostForm(url string, headers map[string]string, formValues map[string]string) (int, []*http.Cookie, []byte, error) {
	//构造请求
	formStr := ""
	for k, v := range formValues {
		if formStr == "" {
			formStr = formStr + k + "=" + v
		} else {
			formStr = formStr + "&" + k + "=" + v
		}
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(formStr))
	if err != nil {
		return 0, nil, nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	//发送请求
	resp, err := retryHttpRequest(req)
	if err != nil {
		return 0, nil, nil, err
	}
	defer resp.Body.Close()

	//处理请求
	statusCode := resp.StatusCode
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, nil, err
	}
	cookies := resp.Cookies()
	return statusCode, cookies, result, nil
}

func HttpGet(url string, headers map[string]string, formValues map[string]string) (int, []*http.Cookie, []byte, error) {
	//构造请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, nil, nil, err
	}
	q := req.URL.Query()
	for k, v := range formValues {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	//发送请求
	resp, err := retryHttpRequest(req)
	if err != nil {
		return 0, nil, nil, err
	}
	defer resp.Body.Close()

	//处理请求
	statusCode := resp.StatusCode
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, nil, err
	}
	cookies := resp.Cookies()
	return statusCode, cookies, result, nil
}
