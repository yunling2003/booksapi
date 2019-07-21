package controllers

import (
	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	c.String(200, "Books API")
}