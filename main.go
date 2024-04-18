package main

import (
	"yuuzin217/go-api-sample/controllers"
	"yuuzin217/go-api-sample/models"
	"yuuzin217/go-api-sample/repositories"
	"yuuzin217/go-api-sample/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []*models.Item{
		{ID: 1, Name: "商品１", Price: 1000, Description: "説明１", SoldOut: false},
		{ID: 2, Name: "商品２", Price: 2000, Description: "説明２", SoldOut: true},
		{ID: 3, Name: "商品３", Price: 3000, Description: "説明３", SoldOut: false},
	}
	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	ItemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", ItemController.FindAll)
	r.GET("/items/:id", ItemController.FindByID)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
