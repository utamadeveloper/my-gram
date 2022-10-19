package model

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID      string `gorm:"size:36;primaryKey"`
	Message string `json:"message" form:"message" valid:"required~Message is required"`
	UserID  string `json:"user_id" gorm:"size:36" form:"user_id" valid:"required~User ID is required"`
	PhotoID string `json:"photo_id" gorm:"size:36" form:"photo_id" valid:"required~Photo ID is required"`
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValid := govalidator.ValidateStruct(comment)

	if errValid != nil {
		err = errValid
		return
	}

	comment.ID = uuid.NewString()

	err = nil
	return
}

func (comment *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errValid := govalidator.ValidateStruct(comment)

	if errValid != nil {
		err = errValid
		return
	}
	err = nil
	return
}
