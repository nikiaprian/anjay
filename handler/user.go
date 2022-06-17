package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) UserLogin(c *gin.Context) {
	data, err := handler.Project.Usecase.UserLogin(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(200, sendResponseSuccess{Success: true, Code: 200, Message: "", Data: data})
	return

}

func (handler *Handler) UserRegister(c *gin.Context) {
	data, err := handler.Project.Usecase.UserRegister(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sendResponseSuccess{Success: true, Code: http.StatusCreated, Message: data.Message})
	return
}

func (handler *Handler) UserLoginByProvider(c *gin.Context) {
	url, err := handler.Project.Usecase.UserLoginByProvider(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
	return
}

func (handler *Handler) UserLoginByProviderCallback(c *gin.Context) {
	data, err := handler.Project.Usecase.UserLoginByProviderCallback(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sendResponseSuccess{Success: true, Code: http.StatusCreated, Message: "Success created Account", Data: data})
	return

}
