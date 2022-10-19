package model

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	ID     string `gorm:"size:36;primaryKey"`
	Name   string `json:"name" form:"name" valid:"required~Name is required"`
	Url    string `json:"url"  form:"url" valid:"required~Url is required"`
	UserID string `json:"user_id" gorm:"size:36" form:"user_id" valid:"required~User ID is required"`
}

func (socialMedia *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValid := govalidator.ValidateStruct(socialMedia)

	if errValid != nil {
		err = errValid
		return
	}

	socialMedia.ID = uuid.NewString()

	err = nil
	return
}

func (socialMedia *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errValid := govalidator.ValidateStruct(socialMedia)

	if errValid != nil {
		err = errValid
		return
	}
	err = nil
	return
}
