package yingyan

import (
	"sort"
	"strings"
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

func sortParam(param map[string]string) ([]string, string) {
	keySlice := make([]string, 0, 10)
	sortedKVSlice := make([]string, 0, 10)

	for k, _ := range param {
		keySlice = append(keySlice, k)
	}

	sort.Strings(keySlice)

	for i := 0; i < len(keySlice); i++ {
		sortedKVSlice = append(sortedKVSlice, keySlice[i]+"="+x[keySlice[i]])
	}

	return keySlice, strings.Join(sortedKVSlice, "&")
}

func (c *Client) Post(path string, param map[string]string) (body []byte, err error) {

	param["ak"] = c.ak
	param["service_id"] = c.serviceID

	sortKeys, sortQueryString := sortParam(param)

	data := &fasthttp.Args{}

	for _, k := range sortKeys {
		data.Add(sortKeys[k], param[sortKeys[k]])
	}

	sn := c.sign(path, sortQueryString)

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
	param["ak"] = c.ak
	param["service_id"] = c.serviceID

	sortKeys, sortQueryString := sortParam(param)

	data := &fasthttp.Args{}

	for _, k := range sortKeys {
		data.Add(sortKeys[k], param[sortKeys[k]])
	}

	sn := c.sign(path, sortQueryString)

	if sn != "" {
		data.Add("sn", sn)
	}

	sn := c.sign(path, sortQueryString)

	if sn != "" {
		data.Add("sn", sn)
	}

	_, body, err = c.httpClient.GetTimeout(nil, apiRootPath+path+"?"+sortQueryString, 10*time.Second)

	if err != nil {
		return nil, err
	}

	return body, nil
}
