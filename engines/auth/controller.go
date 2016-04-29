package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/itpkg/gails"
)

type Controller struct {
}

func (p *Controller) Mount(*gin.Engine) {

}

func (p *Controller) Assets() {

}

func init() {
	gails.Use(&Controller{})
}
