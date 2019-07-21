package wechat

import (
  "fmt"
  "encoding/json"
  "booksapi/src/config"
  "github.com/jdomzhang/resty"
)

type LoginResponse struct {
  OpenID       string  `json:"openid"`
  SessionKey   string  `json:"session_key"`
  UnionID      string  `json:"unionid"`
  ErrCode      int     `json:"errcode"`
  ErrMsg       string  `json:"errmsg"`
}

func JSCode2Session(result *LoginResponse, code string) error {
  appid := config.All["wechat.app.appid"]
  secret := config.All["wechat.app.secret"]

  url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
      appid, secret, code)
  
  resp, err := resty.R().Get(url)

  if err != nil {
    return err
  }

  if err := json.Unmarshal(resp.Body(), result); err != nil {
    return err
  }

  if result.ErrCode != 0 {
    return fmt.Errorf("Error: %v", result)
  }

  return nil
}