package entities

import "github.com/pborman/uuid"

type tweet struct {
	ID          string `json:id`
	Description string `json:description`
}

func NewTweet() *tweet {
	tweet := tweet{
		ID: uuid.New(),
	}

	return &tweet
}
