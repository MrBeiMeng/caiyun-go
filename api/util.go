package api

import (
	"caiyu-go/logging"
	"caiyu-go/util"
	"encoding/json"
	"fmt"
	url2 "net/url"
)

// CaiYunGet url 二级地址， 经纬度 "location": [39.2072, 101.6656],
func CaiYunGet(url string, location []float64, T interface{}, params map[string]string) {
	if len(location) != 2 {
		panic("错误的经纬度地址")
	}

	baseUrl := "https://api.caiyunapp.com/v2.6/"
	token := util.Get("Token")
	requestingUrl, err := url2.JoinPath(baseUrl, token, fmt.Sprintf("%.4f,%.4f", location[0], location[1]), url)
	if err != nil {
		panic(err.Error())
	}

	if params == nil {
		params = make(map[string]string)
	}
	params["alert_util"] = "true"

	respAllByte, err := util.HttpGet(requestingUrl, params)

	err = json.Unmarshal(respAllByte, T)
	if err != nil {
		panic(err)
	}

	logging.Info("注入成功", T)
}
