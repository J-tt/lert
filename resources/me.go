package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func (R *Resources) AttachMe(group *gin.RouterGroup) error {
	group.GET("/", R.me)
	return nil
}

func (R *Resources) me(c *gin.Context) {
	c.JSON(http.StatusOK, reddit.User{})
	return
}
