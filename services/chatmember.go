package services

import (
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/repository"
	"github.com/bivek/fmt_backend/utils"
)

// ChatMemberService -> struct
type ChatMemberService struct {
	repository repository.ChatMemberRepository
}

// NewChatMemberService  -> creates a new ChatMemberService
func NewChatMemberService(repository repository.ChatMemberRepository) ChatMemberService {
	return ChatMemberService{
		repository: repository,
	}
}

// CreateChatMember -> call to create the ChatMember
func (c ChatMemberService) CreateChatMember(ChatMember models.ChatMember) (models.ChatMember, error) {
	return c.repository.Create(ChatMember)
}

// GetAllChatMember -> call to create the ChatMember
func (c ChatMemberService) GetAllChatMember(pagination utils.Pagination) ([]models.ChatMember, int64, error) {
	return c.repository.GetAllChatMember(pagination)
}

// GetOneChatMember -> Get One ChatMember By Id
func (c ChatMemberService) GetOneChatMember(userID int64, roomID int64) (models.ChatMember, error) {
	return c.repository.GetOneChatMember(userID, roomID)
}

// UpdateOneChatMember -> Update One ChatMember By Id
func (c ChatMemberService) UpdateOneChatMember(ChatMember models.ChatMember) error {
	return c.repository.UpdateOneChatMember(ChatMember)
}

// DeleteOneChatMember -> Delete One ChatMember By Id
func (c ChatMemberService) DeleteOneChatMember(ID int64) error {
	return c.repository.DeleteOneChatMember(ID)

}
