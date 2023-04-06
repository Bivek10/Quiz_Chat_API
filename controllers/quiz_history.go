package controllers

import (
	"fmt"
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

type HistoryController struct {
	logger              infrastructure.Logger
	quizHistoryServices services.QuizHistoryServices
	env                 infrastructure.Env
	firbaseServices     services.FirebaseService
	gmailServices       services.GmailService
}

func NewHistoryController(gmailService services.GmailService, logger infrastructure.Logger, quizHistoryService services.QuizHistoryServices, friebaseService services.FirebaseService, evn infrastructure.Env) HistoryController {
	return HistoryController{
		logger:              logger,
		env:                 evn,
		firbaseServices:     friebaseService,
		quizHistoryServices: quizHistoryService,
		gmailServices:       gmailService,
	}

}

// create ->quiz
func (qq HistoryController) CreateHistory(c *gin.Context) {
	quizhistory := models.QuizHistory{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)
	if err := c.ShouldBindJSON(&quizhistory); err != nil {
		qq.logger.Zap.Error("Error [create History] (should Bind Json): ", err)
		responses.HandleError(c, err)
		return
	}
	qq.logger.Zap.Info("user id", quizhistory.Client_ID)
	if err := qq.quizHistoryServices.WithTrx(trx).CreateHistory(quizhistory); err != nil {
		qq.logger.Zap.Error("Error [create history] [db create history]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create history")
		responses.HandleError(c, err)
		return
	}

	user_id, err := strconv.Atoi(quizhistory.Client_ID)
	if err != nil {
		qq.logger.Zap.Error("error converting string to int", err.Error())
		err := errors.InternalError.Wrap(err, "error converting string to int!")
		responses.HandleError(c, err)
		return
	}

	clients, err := qq.quizHistoryServices.GetUserByID(user_id)
	if err != nil {
		qq.logger.Zap.Error("UserNot Found", err.Error())
		err := errors.InternalError.Wrap(err, "Iser not found!")
		responses.HandleError(c, err)
		return
	}

	qq.logger.Zap.Info("================Email Sending-----------", clients.Email)

	emailParms := models.EmailParams{
		To:           clients.Email,
		SubjectData:  "Quiz Score",
		BodyTemplate: "register_body.txt",
		BodyData: models.EmailBody{
			Score: fmt.Sprintf("Your score is %d", quizhistory.Score),
		},
	}

	isSend, err := qq.gmailServices.SendEmail(emailParms)

	qq.logger.Zap.Info("================Process Completed-----------", quizhistory.Client_ID)

	if err != nil {
		qq.logger.Zap.Error("Failed to send email", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to sent email")
		responses.HandleError(c, err)
		return
	}

	if isSend {
		responses.SuccessJSON(c, http.StatusOK, "History data created succesfully")
	}

}

// get allquiz  data
func (qq HistoryController) GetAllHistory(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	History, count, err := qq.quizHistoryServices.GetAllHistory(pagination)
	if err != nil {
		qq.logger.Zap.Error("Error geting history records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get history data")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, History, count)
}

// get allquiz by quiz id data
func (qq HistoryController) GetHistoryByUserID(c *gin.Context) {
	id := c.Param("user_id")
	pagination := utils.BuildPagination(c)
	user_id := id

	// if errs != nil {
	// 	qq.logger.Zap.Error("Error converting the string into int", errs.Error())
	// 	err := errors.InternalError.Wrap(errs, "Failed failed to convert error to int")
	// 	responses.HandleError(c, err)
	// 	return
	// }
	quizhistory, count, err := qq.quizHistoryServices.GetHistoryByUserID(pagination, user_id)
	if err != nil {
		qq.logger.Zap.Error("Error geting quiz history records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get quiz history data")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, quizhistory, count)
}

// send Email to user
