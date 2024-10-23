package wechatdomaincheck

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	wechatEndpoint = "https://cgi.urlsec.qq.com/index.php?m=url&a=validUrl&url="
)

type CheckResponse struct {
	Data   string `json:"data"`
	ReCode int    `json:"reCode"`
}

func Check(urlPath string) (*CheckResponse, error) {
	checkURL := wechatEndpoint + url.QueryEscape(urlPath)
	resp, err := http.Get(checkURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var wechatResp CheckResponse
	if err = json.NewDecoder(resp.Body).Decode(&wechatResp); err != nil {
		return nil, err
	}
	return &wechatResp, nil
}
