package models

type Post struct {
	BaseModel
	Title string
	Body string
}

type CreatePostPayload struct {
	Title  string `json:"title" binding:"required"`
	Body string `json:"body" binding:"required"`
}