package controllers

import (
	"net/http"

	"github.com/bivek/fmt_backend/constants"
	"github.com/bivek/fmt_backend/errors"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/responses"
	"github.com/bivek/fmt_backend/services"
	"github.com/bivek/fmt_backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuizController struct {
	logger          infrastructure.Logger
	quizService     services.QuizService
	env             infrastructure.Env
	firbaseServices services.FirebaseService
}

func NewQuizController(logger infrastructure.Logger, quizService services.QuizService, friebaseService services.FirebaseService, evn infrastructure.Env) QuizController {
	return QuizController{
		logger:          logger,
		quizService:     quizService,
		env:             evn,
		firbaseServices: friebaseService,
	}

}

//create ->quiz

func (qq QuizController) CreateQuiz(c *gin.Context) {
	quiz := models.Quizs{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)
	if err := c.ShouldBindJSON(&quiz); err != nil {
		qq.logger.Zap.Error("Error [create quiz] (should Bind Json): ", err)
		responses.HandleError(c, err)
		return
	}

	if err := qq.quizService.WithTrx(trx).CreateQuiz(quiz); err != nil {
		qq.logger.Zap.Error("Error [create user] [db create quiz]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create quiz")

		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Quiz data created succesfully")

}

// get allquiz  data
func (qq QuizController) GetAllQuiz(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	quizs, count, err := qq.quizService.GetAllQuiz(pagination)
	if err != nil {
		qq.logger.Zap.Error("Error geting quiz records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get quiz data")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, quizs, count)
}
