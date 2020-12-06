package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Url             *string            `json:"url" binding:"required"`
	Category        *TrickTipsCategory `json:"category" binding:"required"`
	Level           *TrickTipsLevel    `json:"level" binding:"required"`
	Sequence        []*string          `json:"sequence" binding:"required"`
	DescriptionStep []*string          `json:"descriptionStep" binding:"required"`
	TotalStep       *uint              `json:"totalStep" binding:"required"`
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
