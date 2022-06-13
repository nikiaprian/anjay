package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (blogHandler *Handler) GetAllBlog(c *gin.Context) {
	data, err := blogHandler.Project.Usecase.GetAllBlog(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data})
	return
}
func (blogHandler *Handler) CreateBlog(c *gin.Context) {
	data, err := blogHandler.Project.Usecase.CreateBlog(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sendResponseSuccess{Success: true, Code: 201, Data: data})
	return
}
func (blogHandler *Handler) UpdateBlog(c *gin.Context) {
	data, err := blogHandler.Project.Usecase.UpdateBlog(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data})
	return
}
