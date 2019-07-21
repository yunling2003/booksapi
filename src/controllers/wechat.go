package controllers

import (
  "fmt"
  "booksapi/src/models/orm"
  "booksapi/src/models/biz"
  "booksapi/src/models/auth"
  "booksapi/src/models/vo"
	"github.com/gin-gonic/gin"
)

type Wechat struct{}

func (*Wechat) WeChatLogin(c *gin.Context) {
  code := c.Query("code")

  if ok := code != ""; !ok {
    renderErrorMessage(c, "code is required")
    return
  }

  var user orm.User
  if err := biz.LoginUserWithWechatJSCode(&user, code); err != nil {
    renderError(c, err)
    return
  }

  token := auth.GenUserJwt(&user)
  SetHeaderToken(c, token)

  renderJSON(c, &user)
}

func (*Wechat) WechatGetUserInfo(c *gin.Context) {
  var wxData vo.WXData
	if err := c.ShouldBind(&wxData); err != nil {
		renderError(c, err)
		return
	}

	lc := getLoginContext(c)
	userID := lc.UserID
	openID := wxData.WechatOpenID
	if openID == "" {
		openID = lc.OpenID
	}
	fmt.Printf("current user is [%d]\n", userID)

	if user, err := biz.SaveWXUserInfo(wxData, userID, openID); err != nil {
		renderError(c, err)
	} else {		
		token := auth.GenUserJwt(user)
		SetHeaderToken(c, token)
		renderJSON(c, user)
	}
}