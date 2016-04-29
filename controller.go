package gails

import "github.com/gin-gonic/gin"

//Controller controller
type Controller interface {
	Mount(*gin.Engine)
}
