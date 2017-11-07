package yingyan

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"

	"github.com/valyala/fasthttp"
)

func (c *Client) sign(uri string, param *fasthttp.Args) (sn string) {
	if c.sk == "" {
		return
	}

	o := uri + "?" + fmt.Sprint(param.QueryString()) + c.sk
	hash := md5.New()
	hash.Write([]byte(url.QueryEscape(o)))
	sn = hex.EncodeToString(hash.Sum(nil))
	return
}
