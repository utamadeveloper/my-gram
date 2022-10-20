package model

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	ID       string    `gorm:"size:36;primaryKey"`
	Title    string    `json:"title" form:"title" valid:"required~Title is required"`
	Caption  string    `json:"caption" form:"caption" valid:"required~Caption is required"`
	Url      string    `json:"url"  form:"url" valid:"required~Url is required"`
	UserID   string    `json:"user_id" gorm:"size:36" form:"user_id" valid:"required~User ID is required"`
	Comments []Comment `json:"comments,omitempty" gorm:"foreignkey:PhotoID"`
}

func (photo *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValid := govalidator.ValidateStruct(photo)

	if errValid != nil {
		err = errValid
		return
	}

	photo.ID = uuid.NewString()

	err = nil
	return
}

func (photo *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errValid := govalidator.ValidateStruct(photo)

	if errValid != nil {
		err = errValid
		return
	}
	err = nil
	return
}
