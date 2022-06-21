package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (forumHandler *Handler) GetAllForum(c *gin.Context) {
	data, err := forumHandler.Project.Usecase.GetAllForum(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: 200, Data: data})
	return
}
