package repository

import (
	"gorm.io/gorm"

	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/utils"
)

// ChatMemberRepository database structure
type ChatMemberRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewChatMemberRepository creates a new ChatMember repository
func NewChatMemberRepository(db infrastructure.Database, logger infrastructure.Logger) ChatMemberRepository {
	return ChatMemberRepository{
		db:     db,
		logger: logger,
	}
}

func (c ChatMemberRepository) WithTrx(trxHandle *gorm.DB) ChatMemberRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transction Database not found in gin context")
		return c
	}
	c.db.DB = trxHandle
	return c
}

// Create ChatMember
func (c ChatMemberRepository) Create(ChatMember models.ChatMember) (models.ChatMember, error) {
	return ChatMember, c.db.DB.Create(&ChatMember).Error
}

// GetAllChatMember -> Get All ChatMember
func (c ChatMemberRepository) GetAllChatMember(pagination utils.Pagination) ([]models.ChatMember, int64, error) {
	var ChatMember []models.ChatMember
	var totalRows int64 = 0
	queryBuider := c.db.DB.Model(&models.ChatMember{}).Offset(pagination.Offset).Order(pagination.Sort)

	if !pagination.All {
		queryBuider = queryBuider.Limit(pagination.PageSize)
	}

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuider.Where(c.db.DB.Where("`chatmember`.`title` LIKE ?", searchQuery))
	}

	err := queryBuider.
		Find(&ChatMember).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return ChatMember, totalRows, err
}

// GetOneChatMember -> Get One ChatMember By Id
func (c ChatMemberRepository) GetOneChatMember(userID int64, roomID int64) (models.ChatMember, error) {
	ChatMember := models.ChatMember{}
	return ChatMember, c.db.DB.
		Where("user_id = ?", userID).Where("room_id= ?", roomID).First(&ChatMember).Error
}

// UpdateOneChatMember -> Update One ChatMember By Id
func (c ChatMemberRepository) UpdateOneChatMember(ChatMember models.ChatMember) error {
	return c.db.DB.Model(&models.ChatMember{}).
		Where("id = ?", ChatMember.ID).
		Updates(map[string]interface{}{
			"user_id": ChatMember.UserID,
			"room_id": ChatMember.RoomID,
		}).Error
}

// DeleteOneChatMember -> Delete One ChatMember By Id
func (c ChatMemberRepository) DeleteOneChatMember(ID int64) error {
	return c.db.DB.
		Where("id = ?", ID).
		Delete(&models.ChatMember{}).
		Error
}
