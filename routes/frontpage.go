package resources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type Frontpage struct{}

func (R *Frontpage) Attach(group *gin.RouterGroup) error {
	group.GET("/", R.GET)
	return nil
}

func (S *Frontpage) GET(c *gin.Context) {
	c.JSON(http.StatusOK, []reddit.Post{})
}
