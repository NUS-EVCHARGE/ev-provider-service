package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	headerContentType = "Content-Type"
	contentTypeJSON   = "application/json"
)

type RequestOption func(req *http.Request)

func WithBearerToken(token string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set("Authorization", "Bearer "+token)
	}
}

func WithBasicAuth(username, password string) RequestOption {
	return func(req *http.Request) {
		req.SetBasicAuth(username, password)
	}
}

func WithHeaders(headers map[string]string) RequestOption {
	return func(req *http.Request) {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
}

func LaunchHttpRequest(method, url string, params map[string]string, body interface{}, opts ...RequestOption) (response *http.Response, err error) {
	//logrus.Info("LaunchHttpRequest url: ", url)
	request, err := newRequest(url, method, params, body)
	if err != nil {
		return
	}

	for _, o := range opts {
		o(request)
	}

	return sendRequest(request)
}

func newRequest(url, method string, params map[string]string, body interface{}) (req *http.Request, err error) {
	var tempBody []byte
	if tempBody, err = json.Marshal(body); err != nil {
		return nil, err
	}

	req, err = http.NewRequest(method, url, bytes.NewBuffer(tempBody))
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	req.Header.Set(headerContentType, contentTypeJSON)

	return req, nil
}

func sendRequest(req *http.Request) (response *http.Response, err error) {
	return http.DefaultClient.Do(req)
}

func HttpResponseDefaultParser(resp *http.Response, value any) error {

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		_ = json.Unmarshal(b, &value)
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(b))
	}

	err = json.Unmarshal(b, &value)
	if err != nil {
		return fmt.Errorf("error: %v\n response in bytes: %v", err, b)
	}
	return err
}
