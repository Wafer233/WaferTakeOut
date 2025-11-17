package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const Code2SessionURL = "https://api.weixin.qq.com/sns/jscode2session"

type WxResp struct {
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrMsg     string `json:"errmsg"`
	OpenId     string `json:"openid"`
	ErrCode    int32  `json:"errcode"`
}

func GetWxResp(code string) (*WxResp, error) {

	url := fmt.Sprintf(
		"%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		Code2SessionURL, APPID, SECRET, code,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("请求URL失败")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("读取响应失败")
	}

	var data WxResp
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, errors.New("响应反序列化失败")
	}

	if data.ErrCode != 0 {
		return nil, errors.New(data.ErrMsg)
	}

	return &data, nil
}
