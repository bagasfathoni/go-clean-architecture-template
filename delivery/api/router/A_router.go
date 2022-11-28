package router

import (
	"github.com/bagasfathoni/go-clean-architecture-template/delivery/api/controller"
	"github.com/gin-gonic/gin"
)

type ARouter struct {
	router      *gin.Engine
	aController controller.AController
}

func InitARouter(r *gin.Engine, a controller.AController) *ARouter {
	config := ARouter{router: r, aController: a}
	aRouter := config.router.Group("/api/a")

	aRouter.POST("/insert", a.CreateNewA)
	aRouter.GET("/search", a.CreateNewA, a.GetByLookAlikeName)
	aRouter.GET("/search?status=true", a.GetAllWithTrueStatus)
	aRouter.PUT("/update/name", a.UpdateNameById)
	aRouter.PUT("/update/status", a.UpdateStatusById)
	aRouter.DELETE("/delete", a.DeleteById)

	return &config
}
