package yingyan

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"

	"github.com/valyala/fasthttp"
)

func (c *Client) sign(uri string, param *fasthttp.Args) (sn string) {
	if c.sk == "" {
		return
	}
	o := uri + "?" + param.Encode() + c.sk
	hash := md5.New()
	hash.Write([]byte(url.QueryEscape(o)))
	sn = hex.EncodeToString(hash.Sum(nil))
	return
}
