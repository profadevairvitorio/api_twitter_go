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
