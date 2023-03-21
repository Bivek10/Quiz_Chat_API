package controllers

import (
	"net/http"
	"strconv"

	"github.com/bivek/fmt_backend/errors"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/responses"
	"github.com/bivek/fmt_backend/services"
	"github.com/bivek/fmt_backend/utils"
	"github.com/gin-gonic/gin"
)

// ChatRoomController -> struct
type ChatRoomController struct {
	logger                 infrastructure.Logger
	ChatRoomService  services.ChatRoomService
}

// NewChatRoomController -> constructor
func NewChatRoomController(
	logger infrastructure.Logger,
	ChatRoomService services.ChatRoomService,
) ChatRoomController {
	return ChatRoomController{
		logger:                  logger,
		ChatRoomService:  ChatRoomService,
	}
}

// CreateChatRoom -> Create ChatRoom
func (cc ChatRoomController) CreateChatRoom(c *gin.Context) {
	ChatRoom := models.ChatRoom{}

	if err := c.ShouldBindJSON(&ChatRoom); err != nil {
		cc.logger.Zap.Error("Error [CreateChatRoom] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind ChatRoom")
		responses.HandleError(c, err)
		return
	}

	if _, err := cc.ChatRoomService.CreateChatRoom(ChatRoom); err != nil {
		cc.logger.Zap.Error("Error [CreateChatRoom] [db CreateChatRoom]: ", err.Error())
		err := errors.BadRequest.Wrap(err, "Failed To Create ChatRoom")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "ChatRoom Created Sucessfully")
}

// GetAllChatRoom -> Get All ChatRoom
func (cc ChatRoomController) GetAllChatRoom(c *gin.Context) {

	pagination := utils.BuildPagination(c)
	pagination.Sort = "created_at desc"
	ChatRoom, count, err := cc.ChatRoomService.GetAllChatRoom(pagination)

	if err != nil {
		cc.logger.Zap.Error("Error finding ChatRoom records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed To Find ChatRoom")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, ChatRoom, count)

}

// GetOneChatRoom -> Get One ChatRoom
func (cc ChatRoomController) GetOneChatRoom(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ChatRoom, err := cc.ChatRoomService.GetOneChatRoom(ID)

	if err != nil {
		cc.logger.Zap.Error("Error [GetOneChatRoom] [db GetOneChatRoom]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed To Find ChatRoom")
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, ChatRoom)

}

// UpdateOneChatRoom -> Update One ChatRoom By Id
func (cc ChatRoomController) UpdateOneChatRoom(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ChatRoom := models.ChatRoom{}

	if err := c.ShouldBindJSON(&ChatRoom); err != nil {
		cc.logger.Zap.Error("Error [UpdateChatRoom] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "failed to update ChatRoom")
		responses.HandleError(c, err)
		return
	}
	ChatRoom.ID = ID

	if err := cc.ChatRoomService.UpdateOneChatRoom(ChatRoom); err != nil {
		cc.logger.Zap.Error("Error [UpdateChatRoom] [db UpdateChatRoom]: ", err.Error())
		err := errors.InternalError.Wrap(err, "failed to update ChatRoom")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "ChatRoom Updated Sucessfully")
}

// DeleteOneChatRoom -> Delete One ChatRoom By Id
func (cc ChatRoomController) DeleteOneChatRoom(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := cc.ChatRoomService.DeleteOneChatRoom(ID)

	if err != nil {
		cc.logger.Zap.Error("Error [DeleteOneChatRoom] [db DeleteOneChatRoom]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Delete ChatRoom")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "ChatRoom Deleted Sucessfully")
}
