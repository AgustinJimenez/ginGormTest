package tests

import (
	"bytes"
	"encoding/json"
	"go_practice/models"
	repositories "go_practice/repositories/posts"
	"go_practice/tests"
	"go_practice/utils"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func generateAnyPost() models.CreatePostPayload {
	return models.CreatePostPayload{
		Title: tests.Fake.Lorem().Sentence(3),
		Body: tests.Fake.Lorem().Sentence(50),
	}
}

func TestCreatePostApi(t *testing.T) {
	tests.ResetApp()
	tests.BeginDbTransaction()


	newPostPayload := generateAnyPost()
	newPost, err := json.Marshal(newPostPayload)
	utils.CheckTestError(t, err)
	response := tests.TestHttpRequest("POST", "/posts", bytes.NewBuffer(newPost))
	assert.Equal(t, http.StatusOK, response.Code)	

	newWrongPostPayload1 := generateAnyPost()
	newWrongPostPayload1.Title = ""
	newPost, err = json.Marshal(newWrongPostPayload1)
	utils.CheckTestError(t, err)
	response = tests.TestHttpRequest("POST", "/posts", bytes.NewBuffer(newPost))

	assert.Equal(t, http.StatusBadRequest, response.Code)


	newWrongPostPayload2 := generateAnyPost()
	newWrongPostPayload2.Body = ""
	newPost, err = json.Marshal(newWrongPostPayload1)
	utils.CheckTestError(t, err)
	response = tests.TestHttpRequest("POST", "/posts", bytes.NewBuffer(newPost))
	assert.Equal(t, http.StatusBadRequest, response.Code)
	
	tests.RollbackDbTransaction()
}

func TestGetPostsApi(t *testing.T) {
	tests.ResetApp()
	tests.BeginDbTransaction()

	postsData := [4] models.CreatePostPayload {
		generateAnyPost(),
		generateAnyPost(),
		generateAnyPost(),
		generateAnyPost(),
	}
	var newPosts [len(postsData)] models.Post

	for i, post := range postsData {
		var result *gorm.DB;
		newPosts[i], result = repositories.CreatePost(repositories.CreatePostDataType{
			Title: post.Title,
			Body: post.Body,
		})
		utils.CheckDbTestError(t, result)
	}

	res := tests.TestHttpRequest("GET", "/posts", nil)

	assert.Equal(t, http.StatusOK, res.Code)	

	var response struct {
        Posts []models.Post `json:"posts"`
    }
	err := json.Unmarshal(res.Body.Bytes(), &response)
	utils.CheckTestError(t, err)
	var countDetectedNewPosts = 0
	for _, response_post := range response.Posts {	
		for _, newPost := range newPosts {	
			if newPost.ID == response_post.ID {
				countDetectedNewPosts++
			}
		}
	}

	assert.Equal(t, len(newPosts), countDetectedNewPosts)
	tests.RollbackDbTransaction()

}

func TestPostShowApi(t *testing.T){
	tests.ResetApp()
	tests.BeginDbTransaction()

	postData := generateAnyPost()
	newPost, result := repositories.CreatePost(repositories.CreatePostDataType{
		Title: postData.Title,
		Body: postData.Body,
	})
	utils.CheckDbTestError(t, result)
	
	res := tests.TestHttpRequest("GET", "/posts/" + strconv.Itoa(int(newPost.ID)), nil)

	assert.Equal(t, http.StatusOK, res.Code)
	var response struct {
        Post models.Post `json:"post"`
    }
	err := json.Unmarshal(res.Body.Bytes(), &response)
	utils.CheckTestError(t, err)
	assert.Equal(t, int(response.Post.ID), int(newPost.ID) )
	assert.Equal(t, response.Post.Title, newPost.Title )
	assert.Equal(t, response.Post.Body, newPost.Body )
	assert.Equal(t, 3, utils.CountFields(response.Post) )
}

func TestPostUpdateApi(t *testing.T){
	tests.ResetApp()
	tests.BeginDbTransaction()

	postData := generateAnyPost()
	newPost, result := repositories.CreatePost(repositories.CreatePostDataType{
		Title: postData.Title,
		Body: postData.Body,
	})
	utils.CheckDbTestError(t, result)
	newPost.Body = "body was updated"
	newPost.Title = "title was updated"
	newPostUpdated, err := json.Marshal(newPost)
	utils.CheckTestError(t, err)
	res := tests.TestHttpRequest("PUT", "/posts/" + strconv.Itoa(int(newPost.ID)), bytes.NewBuffer(newPostUpdated))
	assert.Equal(t, http.StatusOK, res.Code)
	var response struct {
        Post models.Post `json:"post"`
    }
	err = json.Unmarshal(res.Body.Bytes(), &response)
	utils.CheckTestError(t, err)

	assert.Equal(t, int(response.Post.ID), int(newPost.ID) )
	assert.Equal(t, newPost.Title, response.Post.Title  )
	assert.Equal(t, newPost.Body, response.Post.Body )
}

func TestPostDeleteApi(t *testing.T){
	tests.ResetApp()
	tests.BeginDbTransaction()

	postData := generateAnyPost()

	newPost, result := repositories.CreatePost(repositories.CreatePostDataType{
		Title: postData.Title,
		Body: postData.Body,
	})
	utils.CheckDbTestError(t, result)

	response := tests.TestHttpRequest("DELETE", "/posts/" + strconv.Itoa(int(newPost.ID)), nil)

	assert.Equal(t, http.StatusOK, response.Code)

	response = tests.TestHttpRequest("GET", "/posts/" + strconv.Itoa(int(newPost.ID)), nil)

	assert.Equal(t, http.StatusNotFound, response.Code)
}