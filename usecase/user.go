package usecase

import (
	"encoding/json"
	"errors"
	"kel15/models"

	"github.com/gin-gonic/gin"

	utils "kel15/utils"
)

func (usecase *Usecase) UserRegister(c *gin.Context) (*models.UserRegisterResponse, error) {
	var payload models.UserRegisterRequest
	validate := utils.NewValidator()
	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	err = validate.Struct(&payload)

	if err != nil {
		return nil, errors.New(utils.MessageErrorByValidation(err))
	}

	// err = utils.PasswordValidator(payload.Password)

	// if err != nil {
	// 	return nil, err
	// }

	user, err := usecase.repository.GetUserByEmailOrUsername(c, payload)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.New("User already exists")
	}

	err = usecase.repository.CreateUser(c, payload)

	if err != nil {
		return nil, err
	}

	return &models.UserRegisterResponse{
		Message: "User created successfully",
	}, nil
}
