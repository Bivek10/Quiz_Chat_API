package repository

import (
	"gorm.io/gorm"

	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/utils"
)

type ConversationRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewConversationRepository(db infrastructure.Database, logger infrastructure.Logger) ConversationRepository {
	return ConversationRepository{
		db:     db,
		logger: logger,
	}
}

func (c ConversationRepository) WithTrx(trxHandle *gorm.DB) ConversationRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found")
		return c
	}
	c.db.DB = trxHandle

	return c
}

//save message

func (c ConversationRepository) SaveMessage(conversation models.Conversation) error {
	err := c.db.DB.Create(conversation).Error
	return err
}

// cancle message
func (c ConversationRepository) CancleMessage(messageID int) error {

	err := c.db.DB.Where("id = ? ", messageID).Delete(&models.Conversation{}).Error

	return err

}

//get old message

func (c ConversationRepository) GetOldMessage(pagination utils.Pagination, senderID int, receiverID int) ([]models.Conversation, int64, error) {
	var conversationList []models.Conversation

	var count int64

	queryBuilder := c.db.DB.Model(&models.Conversation{}).Where("sender = ?", senderID).Where("receiver =? ", receiverID)

	queryBuilder = queryBuilder.Offset(pagination.Offset).Order("created_at desc")

	if !pagination.All {
		queryBuilder = queryBuilder.Limit(pagination.PageSize)
	}

	err := queryBuilder.Find(&conversationList).Error

	return conversationList, count, err
}
