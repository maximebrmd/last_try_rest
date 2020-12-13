package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
)

type Stance string

// Stance
const (
	Regular Stance = "regular"
	Goofy   Stance = "goofy"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username *string            `json:"username" binding:"required" bson:"username"`
	Email    *string            `json:"email" binding:"required" bson:"email"`
	Password *string            `json:"-" binding:"required" bson:"password"`
	Avatar   *string            `json:"avatar" bson:"avatar"`
	Stance   *Stance            `json:"stance" bson:"stance"`
}

type UserForm struct {
	Username *string               `json:"username"`
	Email    *string               `json:"email"`
	Password *string               `json:"password"`
	Avatar   *multipart.FileHeader `json:"avatar"`
	Stance   *Stance               `json:"stance"`
}

func (s Stance) IsValid() error {
	switch s {
	case Regular, Goofy:
		return nil
	}

	return errors.New("Invalid TrickTipsCategory")
}
