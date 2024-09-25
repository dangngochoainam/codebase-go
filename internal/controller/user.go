package controller

import (
	"example/internal/common/helper/responsehelper"
	"example/internal/dto"
	"example/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		GetUserList(ctx *gin.Context)
	}

	userController struct {
		userUseCase usecase.UserUseCase
	}
)

func NewUserController(userUseCase usecase.UserUseCase) UserController {
	return &userController{
		userUseCase: userUseCase,
	}
}

func (u *userController) GetUserList(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	// handle request from client
	input := &dto.FindUsersRequestDTO{}
	err := appC.C.ShouldBindQuery(input)
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	// call usecase handle business
	data, err := u.userUseCase.FindUsers(&dto.FindUsersRequestDTO{
		Username: input.Username,
	})

	// handle response into client
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}
	appC.Response(http.StatusOK, responsehelper.SUCCESS, data)
	return
}
