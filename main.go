package main

import (
	"github.com/gin-gonic/gin"
	"github.com/j-tt/lert/pkg/config"
	"github.com/j-tt/lert/routes"
	_ "go.elara.ws/go-lemmy"
)

func main() {
	// TODO: implement configuring Gin server defaults from .ENV or config file.
	conf, err := config.New()
	if err != nil {
		panic(err)
	}
	routes := routes.New(conf.Lemmy)

	r := gin.Default()
	routes.Mount(r)
	r.Run(conf.BindAddr)
}
