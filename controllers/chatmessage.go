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

// ChatMessageController -> struct
type ChatMessageController struct {
	logger             infrastructure.Logger
	ChatMessageService services.ChatMessageService
}

// NewChatMessageController -> constructor
func NewChatMessageController(
	logger infrastructure.Logger,
	ChatMessageService services.ChatMessageService,
) ChatMessageController {
	return ChatMessageController{
		logger:             logger,
		ChatMessageService: ChatMessageService,
	}
}

// CreateChatMessage -> Create ChatMessage
func (cc ChatMessageController) CreateChatMessage(c *gin.Context) {
	ChatMessage := models.ChatMessage{}

	if err := c.ShouldBindJSON(&ChatMessage); err != nil {
		cc.logger.Zap.Error("Error [CreateChatMessage] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind ChatMessage")
		responses.HandleError(c, err)
		return
	}

	if _, err := cc.ChatMessageService.CreateChatMessage(ChatMessage); err != nil {
		cc.logger.Zap.Error("Error [CreateChatMessage] [db CreateChatMessage]: ", err.Error())
		err := errors.BadRequest.Wrap(err, "Failed To Create ChatMessage")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "ChatMessage Created Sucessfully")
}

// GetAllChatMessage -> Get All ChatMessage
func (cc ChatMessageController) GetAllChatMessage(c *gin.Context) {
	roomID, err := strconv.ParseInt(c.Param("roomid"), 10, 64)

	if err != nil {
		cc.logger.Zap.Error("invalid room id", err.Error())
		err := errors.InternalError.Wrap(err, "invalid room id")
		responses.HandleError(c, err)
		return
	}
	pagination, err := utils.BuildCursorPagination(c)
	if err != nil {
		cc.logger.Zap.Error("invalid cursor type", err.Error())
		err := errors.InternalError.Wrap(err, "invalid cursor type")
		responses.HandleError(c, err)
		return
	}
	// pagination.Sort = "created_at desc"
	ChatMessage, count, err := cc.ChatMessageService.GetAllChatMessage(pagination, int(roomID))

	if err != nil {
		cc.logger.Zap.Error("Error finding ChatMessage records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed To Find ChatMessage")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, ChatMessage, count)

}

/*

// GetOneChatMessage -> Get One ChatMessage
func (cc ChatMessageController) GetOneChatMessage(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ChatMessage, err := cc.ChatMessageService.GetOneChatMessage(ID)

	if err != nil {
		cc.logger.Zap.Error("Error [GetOneChatMessage] [db GetOneChatMessage]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed To Find ChatMessage")
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, ChatMessage)

}

*/

/*

// UpdateOneChatMessage -> Update One ChatMessage By Id
func (cc ChatMessageController) UpdateOneChatMessage(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ChatMessage := models.ChatMessage{}

	if err := c.ShouldBindJSON(&ChatMessage); err != nil {
		cc.logger.Zap.Error("Error [UpdateChatMessage] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "failed to update ChatMessage")
		responses.HandleError(c, err)
		return
	}
	ChatMessage.ID = ID

	if err := cc.ChatMessageService.UpdateOneChatMessage(ChatMessage); err != nil {
		cc.logger.Zap.Error("Error [UpdateChatMessage] [db UpdateChatMessage]: ", err.Error())
		err := errors.InternalError.Wrap(err, "failed to update ChatMessage")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "ChatMessage Updated Sucessfully")
}

*/

// DeleteOneChatMessage -> Delete One ChatMessage By Id
func (cc ChatMessageController) DeleteOneChatMessage(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := cc.ChatMessageService.DeleteOneChatMessage(ID)

	if err != nil {
		cc.logger.Zap.Error("Error [DeleteOneChatMessage] [db DeleteOneChatMessage]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Delete ChatMessage")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "ChatMessage Deleted Sucessfully")
}
