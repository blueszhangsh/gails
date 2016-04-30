package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/itpkg/gails"
)

type Controller struct {
}

func (p *Controller) Mount(*gin.Engine) {

}

func (p *Controller) Assets() error {
	return nil
}

func init() {
	gails.Use(&Controller{})
}
