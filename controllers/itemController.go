package controllers

import (
	"net/http"
	"yuuzin217/go-api-sample/services"

	"github.com/gin-gonic/gin"
)

type I_ItemController interface {
	FindAll(ctx *gin.Context)
}

type ItemController struct {
	service services.I_ItemService
}

func NewItemController(service services.I_ItemService) I_ItemController {
	return &ItemController{service: service}
}

func (controller *ItemController) FindAll(ctx *gin.Context) {
	items, err := controller.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return // 早期リターン
	}
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
