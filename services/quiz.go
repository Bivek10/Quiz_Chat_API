package services

import (
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/repository"
	"github.com/bivek/fmt_backend/utils"
	"gorm.io/gorm"
)

type QuizService struct {
	repository repository.QuizRepository
}

func NewQuizService(repo repository.QuizRepository) QuizService {
	return QuizService{
		repository: repo,
	}
}

func (c QuizService) WithTrx(trxHandle *gorm.DB) QuizService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

//create ->quiz type

func (c QuizService) CreateQuiz(quiz models.Quizs) error {
	err := c.repository.CreateQuiz(quiz)
	return err
}

//Get -> all quize list

func (c QuizService) GetAllQuiz(pagination utils.Pagination) ([]models.Quizs, int64, error) {
	return c.repository.GetAllQuiz(pagination)
}
