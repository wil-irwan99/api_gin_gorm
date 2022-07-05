package delivery

import (
	"api_gin_ref/config"
	"api_gin_ref/delivery/controller"
	"api_gin_ref/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	managerUscs manager.UseCaseManager
	engine      *gin.Engine
	host        string
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(&appConfig)
	managerRepo := manager.NewRepositoryManager(infra)
	managerUseCase := manager.NewUseCaseManager(managerRepo)
	host := appConfig.Url
	return &appServer{
		managerUscs: managerUseCase,
		engine:      r,
		host:        host,
	}
}

func (a *appServer) initControllers() {
	controller.NewProductController(a.engine, a.managerUscs.CreateProductUseCase(), a.managerUscs.GetProductUseCase())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
