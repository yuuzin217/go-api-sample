package controllers

import (
	"net/http"
	"yuuzin217/go-api-sample/dto"
	"yuuzin217/go-api-sample/services"

	"github.com/gin-gonic/gin"
)

type I_ItemController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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
	itemID, err := parseItemID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id: " + err.Error()})
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

func (controller *ItemController) Create(ctx *gin.Context) {
	var input dto.CreateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newItem, err := controller.service.Create(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": newItem})
}

func (controller *ItemController) Update(ctx *gin.Context) {
	itemID, err := parseItemID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id: " + err.Error()})
		return
	}
	var input dto.UpdateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedItem, err := controller.service.Update(uint(itemID), &input)
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
	ctx.JSON(http.StatusOK, gin.H{"data": updatedItem})
}

func (controller *ItemController) Delete(ctx *gin.Context) {
	itemID, err := parseItemID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id: " + err.Error()})
		return
	}
	err = controller.service.Delete(uint(itemID))
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
	ctx.Status(http.StatusOK)
}
