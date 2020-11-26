package main

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
)

/// 代理语雀API
const (
	baseUrl = "https://www.yuque.com/api/v2/"
)

var (
	client = gorequest.New()
)

func init() {
	client.Set("User-Agent", "gorequest/1.0")
	client.Set("X-Auth-Token", config.Token)
}

type proxyResp struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func proxyGet(relateUrl string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", baseUrl, relateUrl)
	var out proxyResp
	resp, _, errs := client.Clone().Get(url).EndStruct(&out)
	if errs != nil {
		return nil, fmt.Errorf("proxy get %s: %v", relateUrl, errs)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("proxy get %s: %v", relateUrl, fmt.Errorf("status %d: %s", out.Status, out.Message))
	}
	return out.Data, nil
}

func ProxyArticleList() ([]Article, error) {
	data, err := proxyGet(fmt.Sprintf("repos/%s/docs", config.Namespace))
	if err != nil {
		return nil, err
	}
	var out []Article
	err = json.Unmarshal(data, &out)
	return out, err
}

func ProxyArticleDetail(slug string) (*ArticleDetail, error) {
	data, err := proxyGet(fmt.Sprintf("repos/%s/docs/%s", config.Namespace, slug))
	if err != nil {
		return nil, err
	}
	var out ArticleDetail
	err = json.Unmarshal(data, &out)
	return &out, err
}
