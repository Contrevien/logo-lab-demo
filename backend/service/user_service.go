package service

import (
	"context"
	"time"

	"logo-lab-demo/domain"
)

type userService struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserService(userRepository domain.UserRepository, timeout time.Duration) domain.UserService {
	return &userService{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *userService) GetProfileByEmail(c context.Context, userID string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.User{Name: user.Name, Email: user.Email}, nil
}
