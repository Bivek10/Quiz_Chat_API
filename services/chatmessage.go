package services

import (
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/repository"
	"github.com/bivek/fmt_backend/utils"
)

// ChatMessageService -> struct
type ChatMessageService struct {
	repository repository.ChatMessageRepository
}

// NewChatMessageService  -> creates a new ChatMessageService
func NewChatMessageService(repository repository.ChatMessageRepository) ChatMessageService {
	return ChatMessageService{
		repository: repository,
	}
}

// CreateChatMessage -> call to create the ChatMessage
func (c ChatMessageService) CreateChatMessage(ChatMessage models.ChatMessage) (models.ChatMessage, error) {
	return c.repository.Create(ChatMessage)
}

// GetAllChatMessage -> call to create the ChatMessage
func (c ChatMessageService) GetAllChatMessage(pagination utils.CursorPagination, roomID int) ([]models.ChatMessage, int64, error) {
	return c.repository.GetAllChatMessage(pagination, roomID)
}

// GetOneChatMessage -> Get One ChatMessage By Id
// func (c ChatMessageService) GetOneChatMessage(ID int64) (models.ChatMessage, error) {
// 	return c.repository.GetOneChatMessage(ID)
// }

// UpdateOneChatMessage -> Update One ChatMessage By Id
// func (c ChatMessageService) UpdateOneChatMessage(ChatMessage models.ChatMessage) error {
// 	return c.repository.UpdateOneChatMessage(ChatMessage)
// }

// DeleteOneChatMessage -> Delete One ChatMessage By Id
func (c ChatMessageService) DeleteOneChatMessage(ID int64) error {
	return c.repository.DeleteOneChatMessage(ID)

}
