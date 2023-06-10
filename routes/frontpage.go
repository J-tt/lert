package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"go.elara.ws/go-lemmy/types"
)

type Frontpage struct{}

func (R *Frontpage) Attach(group *gin.RouterGroup) error {
	group.GET("/", R.GET)
	return nil
}

func (S *Frontpage) GET(c *gin.Context) {
	l, err := ClientFromContext(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	posts, err := l.Posts(c, types.GetPosts{})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("couldn't retrieve lemmy posts: %w", err))
		return
	}
	var rPosts []reddit.Post
	for _, post := range posts.Posts {
		rPost := reddit.Post{
			Title: post.Post.Name,
		}
		rPosts = append(rPosts, rPost)
	}
	c.JSON(http.StatusOK, rPosts)
}
