package controllers

import (
	"api/api/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (t *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.tweets)
}

func (t *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if error := ctx.BindJSON(&tweet); error != nil {
		return
	}

	t.tweets = append(t.tweets, *tweet)

	ctx.JSON(http.StatusOK, tweet)
}
