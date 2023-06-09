// This package implements the /me routes. Currently it returns dummy data.
package resources

import (
	"github.com/gin-gonic/gin"
)

type Resources struct {
}

func (R *Resources) Mount(router *gin.Engine) error {
	me := router.Group("/me")
	R.AttachMe(me)
	return nil
}
