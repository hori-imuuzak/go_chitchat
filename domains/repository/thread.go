package repository

import "chitchat/domains/model"

type ThreadRepository interface {
	Create(thread model.Thread) bool
	Delete(thread model.Thread) bool
}
