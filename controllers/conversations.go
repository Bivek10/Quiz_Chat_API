package controllers

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/bivek/fmt_backend/constants"
	"github.com/bivek/fmt_backend/errors"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/responses"
	"github.com/bivek/fmt_backend/services"
	"github.com/bivek/fmt_backend/utils"
	"github.com/gin-gonic/gin"
)

type ConversationController struct {
	logger              infrastructure.Logger
	conversationService services.ConversationService
	env                 infrastructure.Env
	firbaseSerives      services.FirebaseService
}

func NewConversationController(logger infrastructure.Logger, conversationservice services.ConversationService, env infrastructure.Env, firebaseService services.FirebaseService) ConversationController {
	return ConversationController{
		logger:              logger,
		conversationService: conversationservice,
		env:                 env,
		firbaseSerives:      firebaseService,
	}

}

func (cc ConversationController) SaveMessage(c *gin.Context) {
	conversation := models.Conversation{}

	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&conversation); err != nil {
		cc.logger.Zap.Error("Error on Binding", err)
		responses.HandleError(c, err)
		return
	}

	if err := cc.conversationService.WithTrx(trx).SaveMessage(conversation); err != nil {
		cc.logger.Zap.Error("Error on saving message", err)
		err := errors.InternalError.Wrap(err, "Failed to save message")

		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Message Sent")
}

func (cc ConversationController) CancleMessage(c *gin.Context) {
	id := c.Param("id")

	messageID, errs := strconv.Atoi(id)
	if errs != nil {
		cc.logger.Zap.Error("Error converting the string into int", errs.Error())
		err := errors.InternalError.Wrap(errs, "Failed failed to convert error to int")
		responses.HandleError(c, err)
		return
	}
	if err := cc.conversationService.CancleMessage(messageID); err != nil {
		cc.logger.Zap.Error("Error on cancling request", err)
		err := errors.InternalError.Wrap(err, "Failed to cancle message")

		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Request Cancled")
}

func (cc ConversationController) GetOldMessage(c *gin.Context) {
	s_id := c.Param("senderid")
	r_id := c.Param("receiverid")
	pagination := utils.BuildPagination(c)

	senderID, errs := strconv.Atoi(s_id)
	if errs != nil {
		cc.logger.Zap.Error("Error converting the string into int", errs.Error())
		err := errors.InternalError.Wrap(errs, "Failed failed to convert error to int")
		responses.HandleError(c, err)
		return
	}
	receiverID, errs := strconv.Atoi(r_id)
	if errs != nil {
		cc.logger.Zap.Error("Error converting the string into int", errs.Error())
		err := errors.InternalError.Wrap(errs, "Failed failed to convert error to int")
		responses.HandleError(c, err)
		return
	}

	oldmessage, count, err := cc.conversationService.GetOldMessage(pagination, senderID, receiverID)

	if err != nil {
		cc.logger.Zap.Error("Error on geting message", err)
		err := errors.InternalError.Wrap(err, "Failed to get message")

		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, oldmessage, count)
}
