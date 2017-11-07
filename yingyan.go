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

func (c *Client) Post(path string, param map[string]string) (body []byte, err error) {
	data := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(data)
	//data := &fasthttp.Args{}
	data.Add("ak", c.ak)
	data.Add("service_id", c.serviceID)
	for k, v := range param {
		data.Add(k, v)
	}
	sn := c.sign(path, data)
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

	data := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(data)

	data.Add("ak", c.ak)
	data.Add("service_id", c.serviceID)
	for k, v := range param {
		data.Add(k, v)
	}
	sn := c.sign(path, data)
	if sn != "" {
		data.Add("sn", sn)
	}

	_, body, err = c.httpClient.GetTimeout(nil, apiRootPath+path+"?"+fmt.Sprint(data.QueryString()), 10*time.Second)

	if err != nil {
		return nil, err
	}

	return body, nil
}
