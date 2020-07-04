package repository

import "chitchat/domains/model"

type UserRepository interface {
	func Create(user model.User) bool
	func Delete(user model.User) bool
}
