package routes

import (
	"github.com/bivek/fmt_backend/controllers"
	"github.com/bivek/fmt_backend/infrastructure"
)

// ThreadRoutes -> struct
type ThreadRoutes struct {
	logger           infrastructure.Logger
	router           infrastructure.Router
	threadController controllers.ThreadController
}

// NewThreadRoutes -> creates new user controller
func NewThreadRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	threadController controllers.ThreadController,

) ThreadRoutes {
	return ThreadRoutes{
		router:           router,
		logger:           logger,
		threadController: threadController,
	}
}

func (m ThreadRoutes) Setup() {
	m.logger.Zap.Info(" Setting up thread routes")
	threads := m.router.Gin.Group("websocket").Use()
	{
		threads.GET("", m.threadController.ServeWs)
	}
}
	