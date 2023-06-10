package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type Me struct{}

func (M *Me) Attach(group *gin.RouterGroup) error {
	group.GET("/", M.GETMe)
	return nil
}

func (M *Me) GETMe(c *gin.Context) {
	c.JSON(http.StatusOK, reddit.User{})
	return
}
