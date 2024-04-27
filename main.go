package main

import (
	"yuuzin217/go-api-sample/controllers"
	"yuuzin217/go-api-sample/models"
	"yuuzin217/go-api-sample/repositories"
	"yuuzin217/go-api-sample/services"

	"github.com/gin-gonic/gin"
)

func getItemController() controllers.I_ItemController {
	items := []*models.Item{
		{ID: 1, Name: "商品１", Price: 1000, Description: "説明１", SoldOut: false},
		{ID: 2, Name: "商品２", Price: 2000, Description: "説明２", SoldOut: true},
		{ID: 3, Name: "商品３", Price: 3000, Description: "説明３", SoldOut: false},
	}
	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)
	return itemController
}

func main() {
	itemController := getItemController()
	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindByID)
	r.POST("/items", itemController.Create)
	r.PUT("/items/:id", itemController.Update)
	r.DELETE("/items/:id", itemController.Delete)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
