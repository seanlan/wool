package wool_sdk

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"math/rand"
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
	//map URL加入密钥拼接
	_source = append(_source, fmt.Sprintf("%s=%s", "key", secretKey))
	sourceStr := strings.Join(_source, "")
	//MD5加密
	h := md5.New()
	h.Write([]byte(sourceStr))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

type WoolSDK struct {
	ApiGateway string
	AppKey     string
	AppSecret  string
}

type ApiResult struct {
	Error        int         `json:"error"`
	ErrorMessage string      `json:"error_msg"`
	Data         interface{} `json:"data,omitempty"`
}

//API参数签名
func (client *WoolSDK) GetSign(jsonObject map[string]string) string {
	return MapToUrlencoded(jsonObject, client.AppSecret)
}

//API参数签名
func (client *WoolSDK) GetNonce() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//基础请求接口
func (client *WoolSDK) Request(apiUrl string, jsonObject map[string]string) (string, error) {
	request := gorequest.New()
	_, body, errs := request.Post(apiUrl).
		Type("form").
		Send(jsonObject).End()
	var err error
	if len(errs) > 0 {
		err = errs[0]
	} else {
		err = nil
	}
	return body, err
}

//API封装请求
func (client *WoolSDK) ApiRequest(method string, jsonObject map[string]string) (ApiResult, error) {
	//API接口请求
	var buf bytes.Buffer
	buf.WriteString(client.ApiGateway)
	buf.WriteString(method)
	apiUrl := buf.String()
	sign := MapToUrlencoded(jsonObject, client.AppSecret)
	jsonObject["sign"] = sign
	body, err := client.Request(apiUrl, jsonObject)
	if err != nil {
		return ApiResult{}, err
	}
	var result ApiResult
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return ApiResult{}, err
	}
	return result, nil
}

//获取用户IMtoken
func (client *WoolSDK) GetIMToken(userID string) (ApiResult, error) {
	return client.ApiRequest("/api/v1/im/get_token",
		map[string]string{
			"appkey": client.AppKey,
			"uid":    userID,
			"nonce":  client.GetNonce(),
		})
}
