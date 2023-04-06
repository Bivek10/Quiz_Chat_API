package repository

import (
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/utils"
)

// ChatMessageRepository database structure
type ChatMessageRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewChatMessageRepository creates a new ChatMessage repository
func NewChatMessageRepository(db infrastructure.Database, logger infrastructure.Logger) ChatMessageRepository {
	return ChatMessageRepository{
		db:     db,
		logger: logger,
	}
}

// Create ChatMessage
func (c ChatMessageRepository) Create(ChatMessage models.ChatMessage) (models.ChatMessage, error) {
	println("on message creation")
	return ChatMessage, c.db.DB.Create(&ChatMessage).Error
}

// GetAllChatMessage -> Get All ChatMessage
func (c ChatMessageRepository) GetAllChatMessage(pagination utils.CursorPagination, roomID int) ([]models.ChatMessage, int64, error) {
	var ChatMessage []models.ChatMessage

	queryBuider := c.db.DB.Model(&models.ChatMessage{}).
		Limit(2).
		Where("room_id = ?", roomID).
		Where("id > ?", pagination.Cursor)

	err := queryBuider.
		Find(&ChatMessage).
		Limit(pagination.Limit).
		Error
	var nextCursor int64

	if err == nil {
		if len(ChatMessage) > 0 {
			nextCursor = ChatMessage[len(ChatMessage)-1].ID
		}
	}
	return ChatMessage, nextCursor, err
}

// GetOneChatMessage -> Get One ChatMessage By Id
// func (c ChatMessageRepository) GetOneChatMessage(ID int64) (models.ChatMessage, error) {
// 	ChatMessage := models.ChatMessage{}
// 	return ChatMessage, c.db.DB.
// 		Where("id = ?", ID).First(&ChatMessage).Error
// }

// // UpdateOneChatMessage -> Update One ChatMessage By Id
// func (c ChatMessageRepository) UpdateOneChatMessage(ChatMessage models.ChatMessage) error {
// 	return c.db.DB.Model(&models.ChatMessage{}).
// 		Where("id = ?", ChatMessage.ID).
// 		Updates(map[string]interface{}{
// 			"user_id": ChatMessage.UserID,
// 			"room_id": ChatMessage.RoomID,
// 		}).Error
// }

// DeleteOneChatMessage -> Delete One ChatMessage By Id
func (c ChatMessageRepository) DeleteOneChatMessage(ID int64) error {
	return c.db.DB.
		Where("id = ?", ID).
		Delete(&models.ChatMessage{}).
		Error
}
