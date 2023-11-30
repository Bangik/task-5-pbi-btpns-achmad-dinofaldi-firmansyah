package controller

import (
	"net/http"
	"path/filepath"
	"strings"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/delivery/middleware"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/helper/common"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/helper/security"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/model"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/usecase"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	router       *gin.Engine
	photoUseCase usecase.PhotoUseCase
}

func (p *PhotoController) Create(c *gin.Context) {
	var photo model.Photo

	id, err := security.GetIdFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "failed",
			"error":  "invalid token " + err.Error(),
		})
		c.Abort()
		return
	}

	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "invalid request body " + err.Error(),
		})
		return
	}

	allowedExtensions := []string{".png", ".jpg", ".jpeg", ".webp"}
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	allowed := false
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			allowed = true
			break
		}
	}
	if !allowed {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "extention is not allowed",
		})
		return
	}

	sess := common.ConnectAws()
	pathFile, err := common.UploadFileToS3(sess, &file, fileHeader.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	photo.Url = pathFile
	photo.UserId = id
	err = c.ShouldBind(&photo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "invalid request body " + err.Error(),
		})

		err = common.DeleteFileFromS3(sess, pathFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}

		return
	}

	err = p.photoUseCase.Create(photo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})

		err = common.DeleteFileFromS3(sess, pathFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

func (p *PhotoController) Get(c *gin.Context) {
	var photo []model.PhotoResponse

	photo, err := p.photoUseCase.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   photo,
	})
}

func (p *PhotoController) Update(c *gin.Context) {
	var photo model.Photo
	photo.Id = c.Param("id")

	photoData, err := p.photoUseCase.FindById(photo.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "photo not found",
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

	if photoData.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "you are not authorized to update this photo",
		})
		return
	} else {
		photo.UserId = userId
	}

	photo.Url = photoData.Url
	err = c.ShouldBind(&photo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	file, fileHeader, err := c.Request.FormFile("image")
	if err == nil {
		allowedExtensions := []string{".png", ".jpg", ".jpeg", ".webp"}
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		allowed := false
		for _, allowedExt := range allowedExtensions {
			if ext == allowedExt {
				allowed = true
				break
			}
		}
		if !allowed {
			c.JSON(http.StatusBadRequest, map[string]any{
				"status":  "failed",
				"message": "extention is not allowed",
			})
			return
		}

		sess := common.ConnectAws()
		pathFile, err := common.UploadFileToS3(sess, &file, fileHeader.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{
				"status":  "failed",
				"message": err.Error(),
			})
			return
		}

		photo.Url = pathFile
		err = common.DeleteFileFromS3(sess, photoData.Url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": err.Error(),
			})
		}
	} else {
		photo.Url = photoData.Url
	}

	err = p.photoUseCase.Update(photo)
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

func (p *PhotoController) Delete(c *gin.Context) {
	var photo model.Photo
	photo.Id = c.Param("id")

	photoData, err := p.photoUseCase.FindById(photo.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "photo not found",
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

	if photoData.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "you are not authorized to update this photo",
		})
		return
	}

	sess := common.ConnectAws()
	err = common.DeleteFileFromS3(sess, photoData.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	err = p.photoUseCase.Delete(photo)
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

func NewPhotoController(router *gin.Engine, photoUseCase usecase.PhotoUseCase) {
	controller := &PhotoController{
		router:       router,
		photoUseCase: photoUseCase,
	}

	router.POST("/photos", middleware.AuthMiddleware(), controller.Create)
	router.GET("/photos", middleware.AuthMiddleware(), controller.Get)
	router.PUT("/photos/:id", middleware.AuthMiddleware(), controller.Update)
	router.DELETE("/photos/:id", middleware.AuthMiddleware(), controller.Delete)
}
