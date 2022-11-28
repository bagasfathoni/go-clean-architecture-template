package controller

import (
	"strconv"

	"github.com/bagasfathoni/go-clean-architecture-template/model"
	"github.com/bagasfathoni/go-clean-architecture-template/usecases"
	"github.com/gin-gonic/gin"
)

type aController struct {
	router *gin.Engine
	aUsec  usecases.AUsecases
}

type AController interface {
	CreateNewA(c *gin.Context)
	UpdateNameById(c *gin.Context)
	UpdateStatusById(c *gin.Context)
	DeleteById(c *gin.Context)
	GetById(c *gin.Context)
	GetAllWithTrueStatus(c *gin.Context)
	GetByLookAlikeName(c *gin.Context)
}

func (a *aController) CreateNewA(c *gin.Context) {
	var newA model.A
	if err := c.ShouldBindJSON(&newA); err != nil {
		BadRequestErrorWithMessage(c, err.Error())
		return
	} else {
		if err := a.aUsec.CreateNewA(&newA); err != nil {
			InternalServerErrorWithMessage(c, err.Error())
			return
		}
		Success(c)
		return
	}
}

func (a *aController) DeleteById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		InternalServerErrorWithMessage(c, "Error when fetching data")
		return
	} else {
		if err := a.aUsec.DeleteById(idInt); err != nil {
			InternalServerErrorWithMessage(c, err.Error())
			return
		}
		Success(c)
		return
	}
}

func (a *aController) UpdateNameById(c *gin.Context) {
	var input struct {
		Id   int
		Name string
	}
	if err := a.aUsec.UpdateNameById(input.Id, input.Name); err != nil {
		InternalServerErrorWithMessage(c, err.Error())
		return
	} else {
		Success(c)
		return
	}
}

func (a *aController) UpdateStatusById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		InternalServerErrorWithMessage(c, "Error when fetching data")
		return
	} else {
		if err := a.aUsec.UpdateStatusById(idInt); err != nil {
			InternalServerErrorWithMessage(c, err.Error())
			return
		}
		Success(c)
		return
	}
}

func (a *aController) GetById(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		InternalServerErrorWithMessage(c, "Error when fetching data")
		return
	} else {
		result, err := a.aUsec.GetById(idInt)
		if err != nil {
			BadRequestErrorWithMessage(c, err.Error())
			return
		}
		SuccessWithMessage(c, result)
		return
	}
}

func (a *aController) GetAllWithTrueStatus(c *gin.Context) {
	result, err := a.aUsec.GetAllWithTrueStatus()
	if err != nil {
		BadRequestErrorWithMessage(c, err.Error())
		return
	} else {
		SuccessWithMessage(c, result)
		return
	}
}

func (a *aController) GetByLookAlikeName(c *gin.Context) {
	searchName := c.Query("name")
	result, err := a.aUsec.GetByLookAlikeName(searchName)
	if err != nil {
		BadRequestErrorWithMessage(c, err.Error())
		return
	} else {
		SuccessWithMessage(c, result)
		return
	}
}

func InitAController(r *gin.Engine, a usecases.AUsecases) AController {
	return &aController{router: r, aUsec: a}
}
