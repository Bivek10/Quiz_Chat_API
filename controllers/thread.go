package controllers

import (
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/socket1"
	"github.com/gin-gonic/gin"
)

type ThreadController struct {
	logger   infrastructure.Logger
	db       infrastructure.Database
	wsServer *socket1.WsServer
}

func NewThreadController(
	logger infrastructure.Logger,
	db infrastructure.Database,
	wsServer *socket1.WsServer,

) ThreadController {
	return ThreadController{
		logger:   logger,
		db:       db,
		wsServer: wsServer,
	}
}

func (tc *ThreadController) ServeWs(c *gin.Context) {
	socket1.ServeWs(tc.wsServer, c)
}
