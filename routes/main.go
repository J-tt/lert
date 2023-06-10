package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"
	"github.com/j-tt/lert/pkg/config"
	"go.elara.ws/go-lemmy"
)

const lemmyClientContext = "lemmy"

type Routes struct {
	l config.Lemmy
}

func New(l config.Lemmy) *Routes {
	routes := Routes{l: l}
	return &routes
}

func (R *Routes) Mount(router *gin.Engine) error {
	var Me *Me
	var Subreddits *Subreddits
	var Frontpage *Frontpage
	var OAuth *OAuth
	router.Use(R.AuthMiddleware())
	router.Use(ginsession.New())

	err := Subreddits.Attach(router.Group("/r"))
	if err != nil {
		return err
	}
	err = Frontpage.Attach(router.Group("/"))
	if err != nil {
		return err
	}
	OAuth.Attach(router)

	apiG := router.Group("/api/v1")
	err = Me.Attach(apiG.Group("/me"))
	if err != nil {
		return err
	}
	return nil
}

// Stub function until OAuth is implemented

func ClientFromContext(c *gin.Context) (*lemmy.Client, error) {
	l, ok := c.Get(lemmyClientContext)
	if !ok {
		return nil, fmt.Errorf("unable to find client")
	}
	return l.(*lemmy.Client), nil
}

func (R *Routes) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		l, err := lemmy.New(R.l.Url)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.Set(lemmyClientContext, l)

		c.Next()
	}
}
