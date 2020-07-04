package repository

import "chitchat/domains/model"

type PostRepository interface {
	Create(post model.Post) bool
}
