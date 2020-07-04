package repository

import "chitchat/domains/model"

type PostRepository interface {
	func Create(post model.Post) bool
}
