package repository

import (
	"chitchat/domains/model"
	"chitchat/domains/repository"
	"chitchat/infrastructures/helper"

	"github.com/jinzhu/gorm"
)

type userRepositoryGorm struct{}

type UserEntity struct {
	gorm.Model
	Uuid  string `gorm:"type:varchar(255);unique_index"`
	Name  string `gorm:"type:varchar(255)"`
	Email string `gorm:"type:varchar(255)"`
}

func (u *UserEntity) TableName() string {
	return "users"
}

func NewUserRepository() repository.UserRepository {
	return &userRepositoryGorm{}
}

func (r *userRepositoryGorm) Create(user model.User) bool {
	db, err := helper.DBOpen()
	defer helper.DBClose(db)
	if err != nil {
		return false
	}

	u := UserEntity{
		Uuid:  user.Uuid,
		Name:  user.Name,
		Email: user.Email,
	}

	db.Create(&u)

	return true
}

func (r *userRepositoryGorm) Delete(user model.User) bool {
	return false
}
