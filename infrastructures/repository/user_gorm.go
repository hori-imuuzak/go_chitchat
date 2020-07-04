package repository

import (
	"github.com/jinzhu/gorm"
	"chitchat/domains/model"
	"chitchat/infrastructures/helper"
)

type UserRepositoryGorm struct{}

type UserEntity struct {
	gorm.Model
	Uuid string `gorm:"type:varchar(255);unique_index"`,
	Name string `gorm:"type:varchar(255)"`,
	Email string `gorm:"type:varchar(255)"`,
}

func (r *UserRepositoryGorm) Create(user model.User) bool {
	db, err := helper.DBOpen()
	defer helper.DBClose(&db)
	if err != nil {
		return false
	}

	u := UserEntity{
		Uuid: user.Uuid,
		Name: user.Name,
		Email: user.Email,
	}

	db.Create(&u)

	return true
}

func (r *UserRepositoryGorm) Delete(user model.User) bool {
	return false
}
