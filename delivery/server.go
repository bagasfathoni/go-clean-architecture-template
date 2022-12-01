package delivery

import (
	"github.com/bagasfathoni/go-clean-architecture-template/config"
	"github.com/bagasfathoni/go-clean-architecture-template/delivery/api/controller"
	"github.com/bagasfathoni/go-clean-architecture-template/manager"
	"github.com/gin-gonic/gin"
)

type appServer struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
	host           string
}

func (a *appServer) initController() {
	controller.NewVendorController(a.engine, a.usecaseManager)
}

func (a *appServer) Run() {
	a.initController()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.InitConfig()
	infraManager := manager.InitInfra(&appConfig)
	repoManager := manager.NewRepositoryManager(infraManager)
	usecaseManager := manager.NewUsecaseManager(repoManager)

	host := appConfig.Url
	return &appServer{
		usecaseManager: usecaseManager,
		engine:         r,
		host:           host,
	}
}
