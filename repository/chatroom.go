package repository

import (
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

// Create ChatRoom
func (c ChatRoomRepository) Create(ChatRoom models.ChatRoom) (models.ChatRoom, error) {
	return ChatRoom, c.db.DB.Create(&ChatRoom).Error
}

// GetAllChatRoom -> Get All ChatRoom
func (c ChatRoomRepository) GetAllChatRoom(pagination utils.Pagination) ([]models.ChatRoom, int64, error) {
	var ChatRoom []models.ChatRoom
	var totalRows int64 = 0
	queryBuider := c.db.DB.Model(&models.ChatRoom{}).Offset(pagination.Offset).Order(pagination.Sort)

	if !pagination.All {
		queryBuider = queryBuider.Limit(pagination.PageSize)
	}

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuider.Where(c.db.DB.Where("`ChatRoom`.`title` LIKE ?", searchQuery))
	}

	err := queryBuider.
		Find(&ChatRoom).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return ChatRoom, totalRows, err
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
