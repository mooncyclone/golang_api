package models

type Track struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

type CreateTrackInput struct {
	Artist string `json:"artist" binding:"required"`
	Title  string `json:"title" binding:"required"`
}
