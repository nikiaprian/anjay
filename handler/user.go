package handler

import (
	"fmt"
	"kel15/models"
	"kel15/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) UserList(c *gin.Context) {
	utils.SetPaginationDefault(c)

	data, err := handler.Project.Usecase.UserList(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	pagination := utils.GetPagination(c)

	c.JSON(200, sendResponseSuccess{Success: true, Code: 200, Message: "", Data: data, Pagination: &pagination})
	return
}

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

func (handler *Handler) UserProfileUpdate(c *gin.Context) {
	file, fileHeader, err := utils.GetFileUpload(c)

	fmt.Println(handler.Project.Storage.BucketName)
	var User models.UserRegisterRequest
	User.Email = c.Request.FormValue("email")
	User.Password = c.Request.FormValue("password")
	User.Username = c.Request.FormValue("username")
	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	fileName, err := utils.UploadToS3(1, handler.Project.Storage, file, fileHeader)

	if err != nil {
		c.JSON(http.StatusBadRequest, sendResponseError{Success: false, Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, sendResponseSuccess{Success: true, Code: http.StatusOK, Message: "", Data: fileName})
}
