package repository

import "chitchat/domains/model"

type UserRepository interface {
	Create(user model.User) bool
	Delete(user model.User) bool
}
