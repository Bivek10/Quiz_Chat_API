package controllers

import (
	"net/http"
	"strconv"

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

type QuestionsController struct {
	logger          infrastructure.Logger
	questionService services.QuestionServices
	env             infrastructure.Env
	firbaseServices services.FirebaseService
}

func NewQuestionController(logger infrastructure.Logger, questionService services.QuestionServices, friebaseService services.FirebaseService, evn infrastructure.Env) QuestionsController {
	return QuestionsController{
		logger:          logger,
		env:             evn,
		firbaseServices: friebaseService,
		questionService: questionService,
	}

}

// create ->quiz
func (qq QuestionsController) CreateQuestion(c *gin.Context) {
	quesitons := models.Questions{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)
	if err := c.ShouldBindJSON(&quesitons); err != nil {
		qq.logger.Zap.Error("Error [create quiz] (should Bind Json): ", err)
		responses.HandleError(c, err)
		return
	}

	if err := qq.questionService.WithTrx(trx).CreateQuestion(quesitons); err != nil {
		qq.logger.Zap.Error("Error [create user] [db create questions]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create questions")

		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Questions data created succesfully")
}

// get allquiz  data
func (qq QuestionsController) GetAllQuestion(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	quesitons, count, err := qq.questionService.GetAllQuestion(pagination)
	if err != nil {
		qq.logger.Zap.Error("Error geting quiz records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get quiz data")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, quesitons, count)
}

// get allquiz by quiz id data
func (qq QuestionsController) GetQuestionByID(c *gin.Context) {
	id := c.Param("quiz_id")
	pagination := utils.BuildPagination(c)
	quiz_id, errs := strconv.Atoi(id)
	if errs != nil {
		qq.logger.Zap.Error("Error converting the string into int", errs.Error())
		err := errors.InternalError.Wrap(errs, "Failed failed to convert error to int")
		responses.HandleError(c, err)
		return
	}
	quesitons, count, err := qq.questionService.GetByQuestionByQuizID(pagination, int64(quiz_id))
	if err != nil {
		qq.logger.Zap.Error("Error geting questions records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get questions data")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, quesitons, count)
}
