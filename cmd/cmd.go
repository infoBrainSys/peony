package cmd

import (
	"github.com/gin-gonic/gin"
	"peony/router"
)

func NewServer() {
	app := gin.Default()
	router.RegisterRouter(app)
	app.Run()
}
