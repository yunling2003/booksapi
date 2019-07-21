package controllers

import (	
	"booksapi/src/models/biz"
	"booksapi/src/models/orm"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	var bookList []orm.Book
	if err := biz.GetAllBooks(&bookList); err != nil {
		renderError(c, err)
	} else {
		renderJSON(c, &bookList)
	}
}