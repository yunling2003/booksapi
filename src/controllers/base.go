package controllers

import (
	"fmt"		
	"net/http"
	"strings"
	"booksapi/src/models/auth"
	"github.com/gin-gonic/gin"
)

const bearerPrefix = "Bearer "
const authorizationHeader = "Authorization"
const loginContextKey = "LoginContextKey"

type LoginContext struct {
	UserID      uint64
	OpenID      string
	Visitor     string
	IsValid     bool
	IsValidUser bool
	IsVisitor   bool
}

func GetHeaderToken(c *gin.Context) string {
	value := c.GetHeader(authorizationHeader)
	if index := strings.Index(value, bearerPrefix); index == 0 {
		value = value[index+len(bearerPrefix):]
	}
	return value
}

func SetHeaderToken(c *gin.Context, token string) {
	value := bearerPrefix + token
	c.Header(authorizationHeader, value)
}

func getLoginContext(c *gin.Context) LoginContext {
	if value, ok := c.Keys[loginContextKey]; ok {
		return value.(LoginContext)
	}
	
	token := GetHeaderToken(c)
	if token == "" {		
		token = RefreshTokenForUserOrVisitor("")
	}

	ok, payload := auth.Check(token)
	loginContext := LoginContext{
		IsValid:       ok,
		UserID:        payload.UserID,
		OpenID:        payload.WechatOpenID,
		IsValidUser:   payload.UserID > 0 && ok,
		Visitor:       payload.Visitor,
		IsVisitor:     payload.Visitor != "",
	}

	c.Set(loginContextKey, loginContext)

	return loginContext
}

func RefreshTokenForUserOrVisitor(token string) string {	
	if len(token) > 0 {
		if ok, payload := auth.Check(token); ok {
			return payload.Gen()
		}
	}

	return auth.GenVisitorJwt()
}

func renderErrorMessage(c *gin.Context, message string, a ...interface{}) {
	if a != nil {
		message = fmt.Sprintf(message, a...)
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": message})

	// stop running other handler in the chain
	c.Abort()
}

func renderError(c *gin.Context, err error) {
	renderErrorMessage(c, err.Error())

	// stop running other handler in the chain
	c.Abort()
}

func renderJSON(c *gin.Context, obj interface{}) {
	c.IndentedJSON(http.StatusOK, obj)
}

func renderString(c *gin.Context, format string, values ...interface{}) {
	c.String(http.StatusOK, format, values...)
}