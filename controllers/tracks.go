package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mooncyclone/golang_api/models"
	"net/http"
)

// GET /tracks
// Получаем список всех треков
func GetAllTracks(context *gin.Context) {
	var tracks []models.Track
	models.ConnectDB().Find(&tracks)

	context.JSON(http.StatusOK, gin.H{"tracks": tracks})
}

// POST /tracks
// Создание трека
func CreateTrack(context *gin.Context) {
	var input models.CreateTrackInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	track := models.Track{Artist: input.Artist, Title: input.Title}
	models.ConnectDB().Create(&track)

	context.JSON(http.StatusOK, gin.H{"tracks": track})
}

// GET /tracks/:id
// Получение одного трека по ID
func GetTrack(context *gin.Context) {
	// Проверяем имеется ли запись
	var track models.Track
	if err := models.ConnectDB().Where("id = ?", context.Param("id")).First(&track).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"tracks": track})
}

// PATCH /tracks/:id
// Изменения информации
func UpdateTrack(context *gin.Context) {
	// Проверяем имеется ли такая запись перед тем как её менять
	var track models.Track
	if err := models.ConnectDB().Where("id = ?", context.Param("id")).First(&track).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateTrackInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.ConnectDB().Model(&track).Update(input)

	context.JSON(http.StatusOK, gin.H{"tracks": track})
}
