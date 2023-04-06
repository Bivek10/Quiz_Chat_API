package services

import (
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/repository"
	"github.com/bivek/fmt_backend/utils"
	"gorm.io/gorm"
)

type QuestionServices struct {
	repository repository.QuestionRepository
}

func NewQuestionServices(repo repository.QuestionRepository) QuestionServices {
	return QuestionServices{
		repository: repo,
	}
}

func (c QuestionServices) WithTrx(trxHandle *gorm.DB) QuestionServices {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}
func (c QuestionServices) CreateQuestion(questions models.Questions) error {
	err := c.repository.CreateQuestion(questions)
	return err
}

func (c QuestionServices) GetAllQuestion(pagination utils.Pagination) ([]models.Questions, int64, error) {
	return c.repository.GetAllQuestion(pagination)
}

func (c QuestionServices) GetByQuestionByQuizID(pagination utils.Pagination, quiz_id int64) ([]models.Questions, int64, error) {
	return c.repository.GetQuestionsByQuiz(pagination, quiz_id)
}
