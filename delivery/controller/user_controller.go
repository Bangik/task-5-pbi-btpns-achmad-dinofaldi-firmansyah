package controller

import (
	"net/http"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/delivery/middleware"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/helper/security"
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

	err = u.userUseCase.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

func (u *UserController) Login(c *gin.Context) {
	var userCred model.UserCredential
	err := c.ShouldBindJSON(&userCred)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	token, err := u.userUseCase.Login(userCred)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"token":  token,
	})
}

func (u *UserController) FindById(c *gin.Context) {
	id, err := security.GetIdFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	user, err := u.userUseCase.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func (u *UserController) Update(c *gin.Context) {
	var user model.User
	id := c.Param("id")

	_, err := u.userUseCase.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "user not found",
		})
		return
	}

	userId, err := security.GetIdFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if userId != id {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "You are not authorized to update this user",
		})
		return
	}

	user.Id = id
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	err = u.userUseCase.Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (u *UserController) Delete(c *gin.Context) {
	var user model.User
	id := c.Param("id")

	_, err := u.userUseCase.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "user not found",
		})
		return
	}

	userId, err := security.GetIdFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if userId != id {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "You are not authorized to delete this user",
		})
		return
	}

	user.Id = id
	err = u.userUseCase.Delete(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func NewUserController(router *gin.Engine, userUseCase usecase.UserUseCase) {
	controller := &UserController{
		router:      router,
		userUseCase: userUseCase,
	}

	routerGroup := router.Group("/users")
	routerGroup.POST("/register", controller.Register)
	routerGroup.POST("/login", controller.Login)
	routerGroup.GET("/", middleware.AuthMiddleware(), controller.FindById)
	routerGroup.PUT("/:id", middleware.AuthMiddleware(), controller.Update)
	routerGroup.DELETE("/:id", middleware.AuthMiddleware(), controller.Delete)
}
