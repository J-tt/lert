package main

import (
	"github.com/gin-gonic/gin"
	"github.com/j-tt/lert/pkg/config"
	"github.com/j-tt/lert/resources"
	_ "go.elara.ws/go-lemmy"
)

func main() {
	// TODO: implement configuring Gin server defaults from .ENV or config file.
	conf, err := config.New()
	if err != nil {
		panic(err)
	}
	var resources resources.Resources

	r := gin.Default()
	resources.Mount(r)
	r.Run(conf.BindAddr)
}
