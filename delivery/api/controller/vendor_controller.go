package controller

import (
	"net/http"

	"github.com/bagasfathoni/go-clean-architecture-template/manager"
	"github.com/gin-gonic/gin"
)

type VendorController struct {
	router         *gin.Engine
	usecaseManager manager.UsecaseManager
}

func (v *VendorController) GetVendorByName(c *gin.Context) {
	name := c.Param("name")
	res, err := v.usecaseManager.VendorUsecase().GetVendorByName(c, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "FAILED",
			"messege": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"data":   res,
	})
	return
}

func (v *VendorController) GetAllVendor(c *gin.Context) {
	res, err := v.usecaseManager.VendorUsecase().GetAllVendor(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "FAILED",
			"messege": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"data":   res,
	})
	return
}

func NewVendorController(router *gin.Engine, usecaseManager manager.UsecaseManager) *VendorController {
	controller := VendorController{
		router:         router,
		usecaseManager: usecaseManager,
	}

	routerGroup := router.Group("/api")
	routerGroup.GET("/vendor/:name", controller.GetVendorByName)
	routerGroup.GET("/vendor", controller.GetAllVendor)

	return &controller
}
