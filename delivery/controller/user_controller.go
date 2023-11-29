package controller

import (
	"net/http"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/model"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router      *gin.Engine
	userUseCase usecase.UserUseCase
}

func (u *UserController) Register(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	user, err = u.userUseCase.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   user,
	})
}

func NewUserController(router *gin.Engine, userUseCase usecase.UserUseCase) {
	controller := &UserController{
		router:      router,
		userUseCase: userUseCase,
	}

	routerGroup := router.Group("/users")
	routerGroup.POST("/register", controller.Register)
}
