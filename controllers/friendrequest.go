package controllers

import (
	"fmt"
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
	chatRoom             services.ChatRoomService
	chatMember           services.ChatMemberService
	env                  infrastructure.Env
	firbaseSerives       services.FirebaseService
}

func NewFriendRequestController(logger infrastructure.Logger,
	firedrequestservice services.FriendRequestService,
	env infrastructure.Env, firebaseService services.FirebaseService,

	chatRoom services.ChatRoomService,
	chatMember services.ChatMemberService,
) FriendRequestController {
	return FriendRequestController{
		logger:               logger,
		friendrequestService: firedrequestservice,
		env:                  env,
		firbaseSerives:       firebaseService,

		chatRoom:   chatRoom,
		chatMember: chatMember,
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

func (fc FriendRequestController) AcceptRequest(c *gin.Context) {
	friendsModel := models.FriendRequest{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&friendsModel); err != nil {
		fc.logger.Zap.Error("Error on Binding", err)
		responses.HandleError(c, err)
		return
	}

	if err := fc.friendrequestService.AcceptRequest(friendsModel); err != nil {
		fc.logger.Zap.Error("Error on sending request", err)
		err := errors.InternalError.Wrap(err, "Failed to Accept Request")

		responses.HandleError(c, err)
		return
	}

	chatRoomModel := models.ChatRoom{Name: "chatroom"}

	dbRoom, err := fc.chatRoom.WithTrx(trx).CreateChatRoom(chatRoomModel)
	fmt.Println("Chatroom id", dbRoom.ID)
	if err != nil {
		fc.logger.Zap.Error("Error [CreatRoom] (CreateRoom) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to Create Room")
		responses.HandleError(c, err)
		return
	}
	chatMemberModel := models.ChatMember{RoomID: dbRoom.ID, UserID: friendsModel.Sender}

	dbMember, err := fc.chatMember.WithTrx(trx).CreateChatMember(chatMemberModel)

	if err != nil {
		fc.logger.Zap.Error("Error [CreatMember] (CreateMember) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to Create Member")
		responses.HandleError(c, err)
		return
	}

	chatMemberModel1 := models.ChatMember{RoomID: dbRoom.ID, UserID: friendsModel.Receiver}

	dbMember1, err := fc.chatMember.WithTrx(trx).CreateChatMember(chatMemberModel1)

	if err != nil {
		fc.logger.Zap.Error("Error [CreatMember1] (CreateMember1) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to Create Member1")
		responses.HandleError(c, err)
		return
	}
	fmt.Println("Member added", dbMember.UserID)
	fmt.Println("Member added", dbMember1.UserID)

	//c.Set("senderid", friendsModel.Sender,)
	//c.Set("receiverid", friendsModel.Receiver)
	//create websocket now.
	//socket.ServeWs1(fc.wsServer, c)

	responses.SuccessJSON(c, http.StatusOK, "Request Accepted and Room created successfully")
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
