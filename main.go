package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mooncyclone/golang_api/controllers"
	"github.com/mooncyclone/golang_api/models"
)

func main() {
	route := gin.Default()

	// Подключение к базе данных
	models.ConnectDB()

	// Маршруты
	route.GET("/tracks", controllers.GetAllTracks)
	route.POST("/tracks", controllers.CreateTrack) // Создание трека
	route.GET("/tracks/:id", controllers.GetTrack) // Получение конкретного трека

	// Запуск сервера
	route.Run()
}
