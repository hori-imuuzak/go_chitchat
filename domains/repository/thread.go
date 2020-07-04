package repository

import "chitchat/domains/model"

type ThreadRepository interface {
	func Create(thread model.Thread) bool
	func Delete(thread model.Thread) bool
}
