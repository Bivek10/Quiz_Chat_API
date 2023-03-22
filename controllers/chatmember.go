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

// ChatMemberController -> struct
type ChatMemberController struct {
	logger                 infrastructure.Logger
	ChatMemberService  services.ChatMemberService
}

// NewChatMemberController -> constructor
func NewChatMemberController(
	logger infrastructure.Logger,
	ChatMemberService services.ChatMemberService,
) ChatMemberController {
	return ChatMemberController{
		logger:                  logger,
		ChatMemberService:  ChatMemberService,
	}
}

// CreateChatMember -> Create ChatMember
func (cc ChatMemberController) CreateChatMember(c *gin.Context) {
	ChatMember := models.ChatMember{}

	if err := c.ShouldBindJSON(&ChatMember); err != nil {
		cc.logger.Zap.Error("Error [CreateChatMember] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind ChatMember")
		responses.HandleError(c, err)
		return
	}

	if _, err := cc.ChatMemberService.CreateChatMember(ChatMember); err != nil {
		cc.logger.Zap.Error("Error [CreateChatMember] [db CreateChatMember]: ", err.Error())
		err := errors.BadRequest.Wrap(err, "Failed To Create ChatMember")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "ChatMember Created Sucessfully")
}

// GetAllChatMember -> Get All ChatMember
func (cc ChatMemberController) GetAllChatMember(c *gin.Context) {

	pagination := utils.BuildPagination(c)
	pagination.Sort = "created_at desc"
	ChatMember, count, err := cc.ChatMemberService.GetAllChatMember(pagination)

	if err != nil {
		cc.logger.Zap.Error("Error finding ChatMember records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed To Find ChatMember")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, ChatMember, count)

}

// GetOneChatMember -> Get One ChatMember
func (cc ChatMemberController) GetOneChatMember(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ChatMember, err := cc.ChatMemberService.GetOneChatMember(ID)

	if err != nil {
		cc.logger.Zap.Error("Error [GetOneChatMember] [db GetOneChatMember]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed To Find ChatMember")
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, ChatMember)

}

// UpdateOneChatMember -> Update One ChatMember By Id
func (cc ChatMemberController) UpdateOneChatMember(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ChatMember := models.ChatMember{}

	if err := c.ShouldBindJSON(&ChatMember); err != nil {
		cc.logger.Zap.Error("Error [UpdateChatMember] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "failed to update ChatMember")
		responses.HandleError(c, err)
		return
	}
	ChatMember.ID = ID

	if err := cc.ChatMemberService.UpdateOneChatMember(ChatMember); err != nil {
		cc.logger.Zap.Error("Error [UpdateChatMember] [db UpdateChatMember]: ", err.Error())
		err := errors.InternalError.Wrap(err, "failed to update ChatMember")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "ChatMember Updated Sucessfully")
}

// DeleteOneChatMember -> Delete One ChatMember By Id
func (cc ChatMemberController) DeleteOneChatMember(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := cc.ChatMemberService.DeleteOneChatMember(ID)

	if err != nil {
		cc.logger.Zap.Error("Error [DeleteOneChatMember] [db DeleteOneChatMember]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Delete ChatMember")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "ChatMember Deleted Sucessfully")
}
