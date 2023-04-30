package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body string
}

type CreatePostPayload struct {
	Title  string `json:"title" binding:"required"`
	Body string `json:"body" binding:"required"`
}