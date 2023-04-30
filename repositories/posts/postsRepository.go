package repositories

import (
	"go_practice/initializers"
	"go_practice/models"

	"gorm.io/gorm"
)

type CreatePostDataType struct {
	Title  string 
	Body string 
}
func CreatePost(data CreatePostDataType) (models.Post, *gorm.DB) {
	post := models.Post{Title: data.Title, Body: data.Body }
	result := initializers.DB.Create(&post)
	return post, result
}