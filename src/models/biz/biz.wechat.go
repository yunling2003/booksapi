package biz

import (
  "booksapi/src/models/orm"
  "booksapi/src/models/wechat"
)

func LoginUserWithWechatJSCode(user *orm.User, code string) error {
  var result wechat.LoginResponse
  if err := wechat.JSCode2Session(&result, code); err != nil {
    return err
  }

  return GetOrCreateUser(user, result.SessionKey, result.OpenID, result.UnionID)
}