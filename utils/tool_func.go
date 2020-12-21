package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/seanlan/packages/logging"
	"sort"
	"strings"
)

func MapToUrlencoded(m map[string]string, secretKey string) string {
	var keys []string
	var _source []string
	for k := range m {
		keys = append(keys, k)
	}
	//字符串排序
	sort.Strings(keys)
	for _, k := range keys {
		_source = append(_source, fmt.Sprintf("%s=%s", k, m[k]))
	}
	sourceStr := strings.Join(_source, "&")
	sourceStr = fmt.Sprintf("%s%s", sourceStr, secretKey)
	//MD5加密
	logging.Info("sourceStr:", sourceStr)
	h := md5.New()
	h.Write([]byte(sourceStr))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
