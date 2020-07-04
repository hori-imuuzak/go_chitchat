package repository

import "chitchat/domains/model"

type UserRepository interface {
	FindByUUID(uuid string) (model.User, bool)
	FindByEmail(email string) (model.User, bool)
	Create(user model.User) bool
	Delete(user model.User) bool
}
