package resources

import (
	"github.com/gin-gonic/gin"
)

type Resources struct{}

func (R *Resources) Mount(router *gin.Engine) error {
	var Me *Me
	var Subreddits *Subreddits
	var Frontpage *Frontpage

	err := Subreddits.Attach(router.Group("/r"))
	if err != nil {
		return err
	}
	err = Frontpage.Attach(router.Group("/"))
	if err != nil {
		return err
	}

	apiG := router.Group("/api/v1")
	err = Me.Attach(apiG.Group("/me"))
	if err != nil {
		return err
	}
	return nil
}
