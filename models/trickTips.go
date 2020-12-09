package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"last_try_rest/utils"
	"mime/multipart"
)

type TrickTipsCategory string
type TrickTipsLevel string

// TrickTipsCategory
const (
	Flat       TrickTipsCategory = "flat"
	Grind      TrickTipsCategory = "grind"
	Transition TrickTipsCategory = "transition"
)

// TrickTipsLevel
const (
	Beginner TrickTipsLevel = "beginner"
	Advance  TrickTipsLevel = "advance"
	Expert   TrickTipsLevel = "expert"
)

type TrickTips struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Title           *string            `json:"title" binding:"required"`
	Thumbnail       *string            `json:"thumbnail" bson:"thumbnail"`
	Url             *string            `json:"url" binding:"required"`
	Category        *TrickTipsCategory `json:"category" binding:"required"`
	Level           *TrickTipsLevel    `json:"level" binding:"required"`
	Sequence        []*string          `json:"sequence" bson:"sequence"`
	DescriptionStep []*string          `json:"descriptionStep" binding:"required"`
	TotalStep       *uint              `json:"totalStep" binding:"required"`
}

type TrickTipsImages struct {
	TrickTipsID *string                 `json:"trickTipsID" form:"trickTipsID" binding:"required"`
	Thumbnail   *multipart.FileHeader   `json:"thumbnail" form:"thumbnail" binding:"required"`
	Sequence    []*multipart.FileHeader `json:"sequence" form:"sequence" binding:"required"`
}

func EncodeFile(file multipart.File) (*string, error) {
	buff, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	base64 := utils.Encode(buff)

	return &base64, nil
}

func (ttl TrickTipsLevel) IsValid() error {
	switch ttl {
	case Beginner, Advance, Expert:
		return nil
	}

	return errors.New("Invalid TrickTipsCategory")
}

func (ttc TrickTipsCategory) IsValid() error {
	switch ttc {
	case Flat, Grind, Transition:
		return nil
	}

	return errors.New("Invalid TrickTipsCategory")
}
