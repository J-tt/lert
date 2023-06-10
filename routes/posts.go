package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type Subreddits struct{}

func (S *Subreddits) Attach(group *gin.RouterGroup) error {
	group.GET("/hot", S.GETHot)
	return nil
}

func (S *Subreddits) GETHot(c *gin.Context) {
	c.JSON(http.StatusOK, []reddit.Post{})
}
