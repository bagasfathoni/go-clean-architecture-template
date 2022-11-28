package api

import (
	"log"
	"time"

	"github.com/bagasfathoni/go-clean-architecture-template/config"
	"github.com/bagasfathoni/go-clean-architecture-template/delivery/api/controller"
	"github.com/bagasfathoni/go-clean-architecture-template/manager"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type apiServer struct {
	usecManager manager.UsecasesManager
	engine      *gin.Engine
	host        string
}

func Server() *apiServer {
	r := initRouterConfiguration()
	cfg := config.InitConfig()
	infra := manager.InitInfra(&cfg)
	repoManager := manager.InitRepositoryManager(infra)
	usecaseManager := manager.InitUsecasesManager(repoManager)

	return &apiServer{
		usecManager: usecaseManager,
		engine:      r,
		host:        "http://localhost:8888",
	}
}

func (a *apiServer) Run() {
	a.initControllers()
	log.Println("Running server on", a.host)
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}

func (a *apiServer) initControllers() {
	controller.InitAController(a.engine, a.usecManager.AUsecases())
}

func initRouterConfiguration() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(corsConfiguration())
	return router
}

func corsConfiguration() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"Origin", "Date", "Content-Length", "Content-Type", "Content-Disposition", "Accept", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
}
