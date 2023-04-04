package services

import (
	"gorm.io/gorm"

	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/repository"
	"github.com/bivek/fmt_backend/utils"
)

// ChatRoomService -> struct
type ChatRoomService struct {
	repository repository.ChatRoomRepository
}

// NewChatRoomService  -> creates a new ChatRoomService
func NewChatRoomService(repository repository.ChatRoomRepository) ChatRoomService {
	return ChatRoomService{
		repository: repository,
	}
}

func (c ChatRoomService) WithTrx(trxHandle *gorm.DB) ChatRoomService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

// CreateChatRoom -> call to create the ChatRoom
func (c ChatRoomService) CreateChatRoom(ChatRoom models.ChatRoom) (models.ChatRoom, error) {
	println("I am at chat room services")
	return c.repository.Create(ChatRoom)
}

// GetAllChatRoom -> call to create the ChatRoom
func (c ChatRoomService) GetAllChatRoom(pagination utils.CursorPagination) ([]models.ChatRoom, int64, error) {
	return c.repository.GetAllChatRoom(pagination)
}

// GetOneChatRoom -> Get One ChatRoom By Id
func (c ChatRoomService) GetOneChatRoom(ID int64) (models.ChatRoom, error) {
	return c.repository.GetOneChatRoom(ID)
}

// UpdateOneChatRoom -> Update One ChatRoom By Id
func (c ChatRoomService) UpdateOneChatRoom(ChatRoom models.ChatRoom) error {
	return c.repository.UpdateOneChatRoom(ChatRoom)
}

// DeleteOneChatRoom -> Delete One ChatRoom By Id
func (c ChatRoomService) DeleteOneChatRoom(ID int64) error {
	return c.repository.DeleteOneChatRoom(ID)

}

func (c ChatRoomService) GetAllRoomByUserID(userID int64)([]models.RoomMember, int64, error){
 return c.repository.GetAllRoomByUserID(userID)
}
