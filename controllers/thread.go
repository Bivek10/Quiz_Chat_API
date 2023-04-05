package controllers

import (
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/socket"
	"github.com/gin-gonic/gin"
)

type ThreadController struct {
	logger   infrastructure.Logger
	db       infrastructure.Database
	wsServer *socket.WsServer
}

func NewThreadController(
	logger infrastructure.Logger,
	db infrastructure.Database,
	wsServer *socket.WsServer,

) ThreadController {
	return ThreadController{
		logger:   logger,
		db:       db,
		wsServer: wsServer,
	}
}

func (tc *ThreadController) ServeWs(c *gin.Context) {
	socket.ServeWs(tc.wsServer, c)
}
