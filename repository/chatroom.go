package repository

import (
	"gorm.io/gorm"

	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/utils"
)

// ChatRoomRepository database structure
type ChatRoomRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewChatRoomRepository creates a new ChatRoom repository
func NewChatRoomRepository(db infrastructure.Database, logger infrastructure.Logger) ChatRoomRepository {
	return ChatRoomRepository{
		db:     db,
		logger: logger,
	}
}

func (c ChatRoomRepository) WithTrx(trxHandle *gorm.DB) ChatRoomRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transction Database not found in gin context")
		return c
	}
	c.db.DB = trxHandle
	return c
}

// Create ChatRoom
func (c ChatRoomRepository) Create(ChatRoom models.ChatRoom) (models.ChatRoom, error) {
	// print("Create chat room in repo", ChatRoom)
	println("I am at chat room repo")
	err := c.db.DB.Create(&ChatRoom).Error
	println(err)
	return ChatRoom, err
	//return ChatRoom, c.db.DB.Create(&ChatRoom).Error

}

// GetAllChatRoom -> Get All ChatRoom
func (c ChatRoomRepository) GetAllChatRoom(pagination utils.CursorPagination) ([]models.ChatRoom, int64, error) {
	var ChatRoom []models.ChatRoom

	queryBuider := c.db.DB.Model(&models.ChatRoom{})

	err := queryBuider.
		Find(&ChatRoom).
		Where("id > ?", pagination.Cursor).
		Limit(pagination.Limit).
		Error

	var nextCursor int64
	if err != nil {
		nextCursor = ChatRoom[len(ChatRoom)-1].ID
	}
	return ChatRoom, nextCursor, err
}

// GetOneChatRoom -> Get One ChatRoom By Id
func (c ChatRoomRepository) GetOneChatRoom(ID int64) (models.ChatRoom, error) {
	ChatRoom := models.ChatRoom{}
	return ChatRoom, c.db.DB.
		Where("id = ?", ID).First(&ChatRoom).Error
}

// UpdateOneChatRoom -> Update One ChatRoom By Id
func (c ChatRoomRepository) UpdateOneChatRoom(ChatRoom models.ChatRoom) error {
	return c.db.DB.Model(&models.ChatRoom{}).
		Where("id = ?", ChatRoom.ID).
		Updates(map[string]interface{}{
			"name": ChatRoom.Name,
		}).Error
}

// DeleteOneChatRoom -> Delete One ChatRoom By Id
func (c ChatRoomRepository) DeleteOneChatRoom(ID int64) error {
	return c.db.DB.
		Where("id = ?", ID).
		Delete(&models.ChatRoom{}).
		Error
}
