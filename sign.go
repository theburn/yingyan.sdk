package yingyan

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
)

var snCache map[string]string = make(map[string]string)

func (c *Client) sign(uri string, sortParamString string) (sn string) {
	if c.sk == "" {
		return
	}

	o := uri + "?" + sortParamString + c.sk

	if sn = snCache[o]; sn != "" {
		return
	} else {
		hash := md5.New()
		hash.Write([]byte(url.QueryEscape(o)))
		sn = hex.EncodeToString(hash.Sum(nil))
		snCache[o] = sn
		return

	}
}
