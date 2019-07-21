package controllers

import (
	"booksapi/src/models/biz"	

	"github.com/gin-gonic/gin"
)

func GetMyUserInfo(c *gin.Context) {
	userID := getLoginContext(c).UserID

	if user, err := biz.GetUserByID(userID); err != nil {
		renderError(c, err)
	} else {
		renderJSON(c, user)
	}
}