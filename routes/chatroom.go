package routes

import (
	"github.com/bivek/fmt_backend/controllers"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/middlewares"
)

// ChatRoomRoutes -> struct
type ChatRoomRoutes struct {
	logger             infrastructure.Logger
	router             infrastructure.Router
	chatRoomController controllers.ChatRoomController
	middleware         middlewares.FirebaseAuthMiddleware
}

// NewChatRoomRoutes -> creates new chatRoom controller
func NewChatRoomRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	chatRoomController controllers.ChatRoomController,
	middleware middlewares.FirebaseAuthMiddleware,
) ChatRoomRoutes {
	return ChatRoomRoutes{
		router:             router,
		logger:             logger,
		chatRoomController: chatRoomController,
		middleware:         middleware,
	}
}

// Setup chatRoom routes
func (c ChatRoomRoutes) Setup() {
	c.logger.Zap.Info(" Setting up chatRoom routes")
	chatRoom := c.router.Gin.Group("/chatroom")
	{
		chatRoom.POST("", c.chatRoomController.CreateChatRoom)
		chatRoom.GET("", c.chatRoomController.GetAllChatRoom)
		chatRoom.GET("/:id", c.chatRoomController.GetOneChatRoom)
		chatRoom.GET("member/:id", c.chatRoomController.GetAllChatRoomByUserID)
		chatRoom.PUT("/:id", c.chatRoomController.UpdateOneChatRoom)
		chatRoom.DELETE("/:id", c.chatRoomController.DeleteOneChatRoom)
	}
}
