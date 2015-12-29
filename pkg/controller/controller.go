package controller

import (

	"github.com/gin-gonic/gin"
)


type Controller struct {
	e *gin.Engine
}

//New new a controller
func New( e *gin.Engine) *Controller {

	return &Controller{
		e: e,
	}

}

//SetupRouters setup all controllers
func (c *Controller) SetupRouters() {
	c.e.GET("/ping2", c.ping)

}


func (p *Controller) ping(ctx *gin.Context) {
	ctx.String(200, "pong pong")
}
