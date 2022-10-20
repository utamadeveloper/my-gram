package model

import (
	"my-gram/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           string         `gorm:"size:36;primaryKey"`
	Username     string         `json:"username,omitempty" gorm:"size:20;index:idx_username,unique" form:"username" valid:"required~Username is required, minstringlength(3)~Username must be at least 3 characters"`
	Email        string         `json:"email,omitempty" gorm:"size:255;index:idx_email,unique" form:"email" valid:"required~Email is required, email~Email must valid email address"`
	Password     string         `json:"password,omitempty" gorm:"size:60" form:"password" valid:"required~Password is required, minstringlength(8)~Password must be at least 8 characters"`
	DOB          datatypes.Date `json:"dob,omitempty" form:"dob" valid:"required~DOB is required"`
	Age          int            `json:"age,omitempty" form:"age" valid:"required~Age is required"`
	SocialMedias []SocialMedia  `json:"social_medias" gorm:"foreignkey:UserID"`
	Photos       []Photo        `json:"photos" gorm:"foreignkey:UserID"`
	Comments     []Comment      `json:"comments" gorm:"foreignkey:UserID"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errValid := govalidator.ValidateStruct(user)

	if errValid != nil {
		err = errValid
		return
	}

	user.ID = uuid.NewString()
	user.Password = helpers.HashPassword(user.Password)

	err = nil
	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errValid := govalidator.ValidateStruct(user)

	if errValid != nil {
		err = errValid
		return
	}

	if user.Password != "" {
		user.Password = helpers.HashPassword(user.Password)
	}

	err = nil
	return
}
