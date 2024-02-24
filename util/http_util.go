package util

import (
	"caiyu-go/logging"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGet(url string, params map[string]string) ([]byte, error) {
	startTime := time.Now()
	resp, _ := httpGet(url, params)

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	httpLog := NewHttpLog(startTime, "GET", url, params, string(respByte))
	if resp.StatusCode == 200 {
		logging.Info(httpLog.String(), resp)
	} else {
		logging.Error(httpLog.String(), resp)
	}

	return respByte, err
}

type HttpLog struct {
	Duration string
	Method   string
	Url      string
	Params   map[string]string
	Resp     string
}

func NewHttpLog(startTime time.Time, method string, url string, params map[string]string, resp string) *HttpLog {
	duration := time.Now().Sub(startTime).String()
	return &HttpLog{Duration: duration, Method: method, Url: url, Params: params, Resp: resp}
}

func (h HttpLog) String() string {

	return fmt.Sprintf("%s|%s|%s|...", h.Method, h.Url, h.Duration)
}

func httpGet(url string, params map[string]string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	query := request.URL.Query()

	for paramKey, value := range params {
		query.Add(paramKey, value)
	}
	request.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("not 200 @%s", url))
	}
	return resp, err
}
