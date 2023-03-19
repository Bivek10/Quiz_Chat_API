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

type FriendRequestController struct {
	logger               infrastructure.Logger
	friendrequestService services.FriendRequestService
	env                  infrastructure.Env
	firbaseSerives       services.FirebaseService
}

func NewFirebaseController(logger infrastructure.Logger, firedrequestservice services.FriendRequestService, env infrastructure.Env, firebaseService services.FirebaseService) FriendRequestController {
	return FriendRequestController{
		logger:               logger,
		friendrequestService: firedrequestservice,
		env:                  env,
		firbaseSerives:       firebaseService,
	}

}

func (fc FriendRequestController) SendRequest(c *gin.Context) {
	friendsModel := models.FriendRequest{}

	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&friendsModel); err != nil {
		fc.logger.Zap.Error("Error on Binding", err)
		responses.HandleError(c, err)
		return
	}

	if err := fc.friendrequestService.WithTrx(trx).SendRequest(friendsModel); err != nil {
		fc.logger.Zap.Error("Error on sending request", err)
		err := errors.InternalError.Wrap(err, "Failed to send request")

		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Request Sent")
}

func (fc FriendRequestController) CancleRequest(c *gin.Context) {
	id := c.Param("id")

	clientID, errs := strconv.Atoi(id)
	if errs != nil {
		fc.logger.Zap.Error("Error converting the string into int", errs.Error())
		err := errors.InternalError.Wrap(errs, "Failed failed to convert error to int")
		responses.HandleError(c, err)
		return
	}
	if err := fc.friendrequestService.CancleRequest(clientID); err != nil {
		fc.logger.Zap.Error("Error on cancling request", err)
		err := errors.InternalError.Wrap(err, "Failed to send request")

		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Request Cancled")
}

func (fc FriendRequestController) GetAcceptedFriend(c *gin.Context) {
	id := c.Param("id")
	pagination := utils.BuildPagination(c)

	clientID, errs := strconv.Atoi(id)
	if errs != nil {
		fc.logger.Zap.Error("Error converting the string into int", errs.Error())
		err := errors.InternalError.Wrap(errs, "Failed failed to convert error to int")
		responses.HandleError(c, err)
		return
	}

	friendlist, count, err := fc.friendrequestService.GetAcceptedFriend(pagination, clientID)

	if err != nil {
		fc.logger.Zap.Error("Error on geting friend", err)
		err := errors.InternalError.Wrap(err, "Failed to send request")

		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, friendlist, count)
}

func (fc FriendRequestController) GetPendingFriend(c *gin.Context) {
	id := c.Param("id")
	pagination := utils.BuildPagination(c)

	clientID, errs := strconv.Atoi(id)
	if errs != nil {
		fc.logger.Zap.Error("Error converting the string into int", errs.Error())
		err := errors.InternalError.Wrap(errs, "Failed failed to convert error to int")
		responses.HandleError(c, err)
		return
	}

	friendlist, count, err := fc.friendrequestService.GetPendingFriend(pagination, clientID)

	if err != nil {
		fc.logger.Zap.Error("Error on geting friend", err)
		err := errors.InternalError.Wrap(err, "Failed to send request")

		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, friendlist, count)
}

func (fc FriendRequestController) GetUnFriend(c *gin.Context) {
	id := c.Param("id")
	pagination := utils.BuildPagination(c)

	clientID, errs := strconv.Atoi(id)
	if errs != nil {
		fc.logger.Zap.Error("Error converting the string into int", errs.Error())
		err := errors.InternalError.Wrap(errs, "Failed failed to convert error to int")
		responses.HandleError(c, err)
		return
	}

	friendlist, count, err := fc.friendrequestService.GetUnFriend(pagination, clientID)

	if err != nil {
		fc.logger.Zap.Error("Error on geting friend", err)
		err := errors.InternalError.Wrap(err, "Failed to send request")

		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, friendlist, count)
}
