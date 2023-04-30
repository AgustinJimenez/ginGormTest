package controllers

import (
	"go_practice/initializers"
	"go_practice/models"
	repositories "go_practice/repositories/posts"
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl -v -i -X GET "http://127.0.0.1:3000/posts"
func PostsIndex(c *gin.Context){
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	  })
}

// curl -v POST -H "Content-Type: application/json" \-d '{"title": "Some title 4", "body": "lorem ipsum 332"}' "http://127.0.0.1:3000/posts"
func PostsCreate(c *gin.Context) {
	var body models.CreatePostPayload

	c.Bind(&body)
	post, result := repositories.CreatePost(repositories.CreatePostDataType{
		Title: body.Title, 
		Body: body.Body,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

    c.JSON(http.StatusOK, gin.H{
      "post": post,
    })
}

func PostsShow(c *gin.Context){
	id := c.Param("id")

	var post models.Post
	initializers.DB.Find(&post, id)

	if int(post.ID) == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	  })
}

type updatePostRequest struct {
    Body  string `json:"body"`
    Title string `json:"title"`
}



func PostsUpdate(c *gin.Context) {
	id := c.Param("id")
	var body updatePostRequest

	c.Bind(&body)

	var post models.Post
	initializers.DB.Find(&post, id)

	post.Title = body.Title
	post.Body = body.Body

	var result = initializers.DB.Save(&post)
	
	if result.Error != nil {
		c.Status(400)
		return
	}

    c.JSON(http.StatusOK, gin.H{
      "post": post,
    })
}


func PostsDelete(c *gin.Context) {
	id := c.Param("id")
	var result = initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

    c.Status(200)
}