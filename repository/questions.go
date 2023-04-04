package repository

import (
	"gorm.io/gorm"

	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/utils"
)

type QuestionRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewQuestionRepository(db infrastructure.Database, logger infrastructure.Logger) QuestionRepository {
	return QuestionRepository{
		db:     db,
		logger: logger,
	}
}

func (q QuestionRepository) WithTrx(trxHandle *gorm.DB) QuestionRepository {
	if trxHandle == nil {
		q.logger.Zap.Error("Transction Database not found in gin context")
		return q
	}
	q.db.DB = trxHandle
	return q
}

// save -> quiz data.
func (q QuestionRepository) CreateQuestion(Question models.Questions) error {
	return q.db.DB.Create(&Question).Error
}

func (c QuestionRepository) GetAllQuestion(pagination utils.Pagination) ([]models.Questions, int64, error) {
	var question []models.Questions
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.Questions{})

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`quizs`.`name` LIKE ?", searchQuery))
	}
	err := queryBuilder.Find(&question).Offset(-1).Limit(-1).Count(&totalRows).Error
	return question, totalRows, err
}

func (c QuestionRepository) GetQuestionsByQuiz(pagination utils.Pagination, quiz_id int64) ([]models.Questions, int64, error) {
	var questions []models.Questions
	//var choices []models.Choices
	//var questionchoices []models.QuestionChoices

	var totalRows int64 = 0

	//queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	//queryBuilder.Joins("questions")
	// queryBuilder = queryBuilder.Model(&models.Questions{}).Joins()
	//queryBuilder.Where(&models.Questions{Quiz_ID: quiz_id})
	//queryBuilder := c.db.DB.Model(&models.Questions{}).Joins("left join `choices` on `questions`.`q_id` = `choices`.q_id")
	err := c.db.DB.Model(&models.Questions{}).Preload("Choices").Where(&models.Questions{Quiz_ID: quiz_id}).Find(&questions).Error
	//err = queryBuilder.Find(&questions).Where(&models.Questions{Quiz_ID: quiz_id}).Error
	println(len(questions) - 1)
	return questions, totalRows, err

}
