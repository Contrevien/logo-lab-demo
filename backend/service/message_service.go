package service

import (
	"context"
	"time"

	"logo-lab-demo/domain"
)

type messageService struct {
	messageRepository domain.MessageRepository
	contextTimeout    time.Duration
}

func NewMessageService(messageRepository domain.MessageRepository, timeout time.Duration) domain.MessageService {
	return &messageService{
		messageRepository: messageRepository,
		contextTimeout:    timeout,
	}
}

func (pu *messageService) GetMessages(c context.Context, page int) (*domain.MessageResponse, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	messages, hasMore, err := pu.messageRepository.Fetch(ctx, page)
	if err != nil {
		return nil, err
	}

	res := &domain.MessageResponse{
		Messages: messages,
		HasMore:  hasMore,
	}

	return res, nil
}

func (pu *messageService) InsertMessage(c context.Context, message domain.CreateMessage) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	err := pu.messageRepository.Create(ctx, message)
	if err != nil {
		return err
	}

	return nil
}
