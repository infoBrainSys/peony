package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"peony/router"
	"peony/utils"
	"time"
)

func init() {
	utils.NewUtilsBase()
}

func NewServer() {
	app := gin.Default()
	app.LoadHTMLGlob("utils/email/*.html")
	router.RegisterRouter(app)

	// 优雅启停
	srv := http.Server{
		Addr:           ":8080",
		Handler:        app,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen:%s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s", err)
	}
	log.Println("Server exiting...")
}
