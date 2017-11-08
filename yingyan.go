package yingyan

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

type Client struct {
	ak         string
	sk         string
	httpClient *fasthttp.Client
	serviceID  string
	s          bool
}

// visit http://lbsyun.baidu.com/apiconsole/key/ 获取ak
// 如果设置的白名单则设置sk ""
func NewClient(ak, sk, serviceID string) *Client {
	return &Client{
		sk:         sk,
		ak:         ak,
		s:          true,
		serviceID:  serviceID,
		httpClient: &fasthttp.Client{},
	}
}

// SetHttpClient you can set your own http client
func (c *Client) SetHttpClient(httpClient *fasthttp.Client) {
	c.httpClient = httpClient
}

var snCache = make(map[string]string)

func (c *Client) Post(path string, param map[string]string) (body []byte, err error) {
	var sn string
	param["ak"] = c.ak
	param["service_id"] = c.serviceID

	snKey := fmt.Sprintf("%s%s", path, param)
	if sn = snCache[snKey]; sn == "" {
		sortQueryString := sortParam(param)
		sn = c.sign(path, sortQueryString)
		snCache[snKey] = sn
	} else {
		sn = snCache[snKey]
	}

	data := &fasthttp.Args{}

	for _, k := range sortParamKeys(param) {
		data.Add(k, param[k])
	}

	if sn != "" {
		data.Add("sn", sn)
	}

	_, body, err = c.httpClient.Post(nil, apiRootPath+path, data)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) Get(path string, param map[string]string) (body []byte, err error) {
	var sn, sortQueryString string
	param["ak"] = c.ak
	param["service_id"] = c.serviceID

	snKey := fmt.Sprintf("%s%s", path, param)
	if sn = snCache[snKey]; sn == "" {
		sortQueryString = sortParam(param)
		sn = c.sign(path, sortQueryString)
		snCache[snKey] = sn
	} else {
		sn = snCache[snKey]
	}

	if sn != "" {
		sortQueryString += sortQueryString + "&sn=" + sn
	}

	_, body, err = c.httpClient.GetTimeout(nil, apiRootPath+path+"?"+sortQueryString, 10*time.Second)

	if err != nil {
		return nil, err
	}

	return body, nil
}
