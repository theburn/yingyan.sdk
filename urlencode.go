package yingyan

import (
	"sort"
	"strings"
)

// 参考 net/http  url.Value::Encode
func appendQuotedPath(dst, src []byte) []byte {
	for _, c := range src {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' ||
			c == '/' || c == '.' || c == '=' || c == '&' || c == '~' || c == '-' || c == '_' {
			dst = append(dst, c)
		} else if c == ' ' {
			dst = append(dst, '+')
		} else {
			dst = append(dst, '%', hexCharUpper(c>>4), hexCharUpper(c&15))

		}

	}
	return dst
}

func hexCharUpper(c byte) byte {
	if c < 10 {
		return '0' + c

	}
	return c - 10 + 'A'

}

func encodeQueryString(querybytes []byte) string {

	var v = make([]byte, 0, len(querybytes))

	return string(appendQuotedPath(v, querybytes))
}

func sortParamKeys(param map[string]string) []string {
	keySlice := make([]string, 0, 10)

	for k, _ := range param {
		keySlice = append(keySlice, k)
	}

	sort.Strings(keySlice)

	return keySlice
}

func sortParam(param map[string]string) ([]string, string) {
	l := len(param)
	keySlice := make([]string, 0, l)
	sortedKVSlice := make([]string, 0, l)

	for k, _ := range param {
		keySlice = append(keySlice, k)
	}

	sort.Strings(keySlice)

	for i := 0; i < len(keySlice); i++ {
		sortedKVSlice = append(sortedKVSlice, keySlice[i]+"="+param[keySlice[i]])
	}

	return keySlice, encodeQueryString([]byte(strings.Join(sortedKVSlice, "&")))
}
