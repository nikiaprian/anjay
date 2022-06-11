package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) UserRegister(c *gin.Context) {
	data, err := handler.Project.Usecase.UserRegister(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sendResponseSuccess{Success: true, Code: 200, Message: data.Message})
	return

}
