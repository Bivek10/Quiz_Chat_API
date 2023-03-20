package services

import (
	"gorm.io/gorm"

	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/repository"
	"github.com/bivek/fmt_backend/utils"
)

type ConversationService struct {
	repository repository.ConversationRepository
}

func NewConversationService(repo repository.ConversationRepository) ConversationService {
	return ConversationService{
		repository: repo,
	}
}

//withTrx

func (c ConversationService) WithTrx(trxHandle *gorm.DB) ConversationService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

func (c ConversationService) SaveMessage(converstationRequest models.Conversation) error {
	err := c.repository.SaveMessage(converstationRequest)
	return err
}

func (c ConversationService) CancleMessage(messageID int) error {
	err := c.repository.CancleMessage(messageID)
	return err
}

func (c ConversationService) GetOldMessage(pagination utils.Pagination, senderID int, receiverID int) ([]models.Conversation, int64, error) {
	conversationlist, count, err := c.repository.GetOldMessage(pagination, senderID, receiverID)
	return conversationlist, count, err
}
