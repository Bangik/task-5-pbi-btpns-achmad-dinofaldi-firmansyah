package delivery

import (
	"fmt"
	"log"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/config"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/delivery/controller"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/delivery/middleware"
	"task-5-pbi-btpns-achmad-dinofaldi-firmansyah/manager"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type appServer struct {
	usecaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
	log            *logrus.Logger
}

func (a *appServer) initController() {
	a.engine.Use(middleware.LogRequestMiddleware(a.log))
	controller.NewUserController(a.engine, a.usecaseManager.UserUseCase())
	controller.NewPhotoController(a.engine, a.usecaseManager.PhotoUseCase())
}

func (a *appServer) Run() {
	a.initController()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err.Error())
	}
}

func Server() *appServer {
	engine := gin.Default()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln("Error Config : ()", err.Error())
	}

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatalln("Error Conection : ", err.Error())
	}

	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	host := fmt.Sprintf("%s:%s", cfg.APIHost, cfg.APIPort)

	return &appServer{
		engine:         engine,
		host:           host,
		usecaseManager: useCaseManager,
		log:            logrus.New(),
	}
}
