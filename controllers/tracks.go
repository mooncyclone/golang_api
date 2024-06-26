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
	var input CreateTrackInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	track := models.Track{Artist: input.Artist, Title: input.Title}
	models.DB.Create(&track)

	context.JSON(http.StatusOK, gin.H{"tracks": track})
}
