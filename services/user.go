package services

import (
	"chitchat/domains/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) Authenticate(email string, password string) bool {
	_, found := s.userRepository.FindByEmail(email)

	if found {
		return true
	}

	return false
}
