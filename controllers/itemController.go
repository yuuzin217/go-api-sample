package controllers

import (
	"net/http"
	"strconv"
	"yuuzin217/go-api-sample/services"

	"github.com/gin-gonic/gin"
)

type I_ItemController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
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

func (controller *ItemController) FindByID(ctx *gin.Context) {
	itemID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	item, err := controller.service.FindByID(uint(itemID))
	if err != nil {
		switch {
		case err.Error() == "Item not found":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}
