package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type fetchConfig struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Hour     int64  `form:"hour"`
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	setupRouter(r)

	start(&http.Server{
		Addr:    fmt.Sprintf(":%s", env("FC_SERVER_PORT", "9000")),
		Handler: r,
	})
}

func setupRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("pong"))
	})
	r.GET("/cctv.test", func(c *gin.Context) {
		var epgList = make([]string, 0)
		for _, c := range cctv {
			epgList = append(epgList, c.epg)
		}
		resp, err := fastTest(epgList)
		if err != nil {
			c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
			return
		}

		c.JSON(http.StatusOK, resp)
	})
	r.GET("/cctv.m3u", func(c *gin.Context) {
		var m3uList = make([]string, 0)
		for _, c := range cctv {
			m3uList = append(m3uList, c.m3u)
		}
		resp, err := fastGet(m3uList)
		if err != nil {
			c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
			return
		}
		fmt.Println("The winner is ", resp.URL)
		c.Data(http.StatusOK, resp.contentType, resp.body)
	})

	r.GET("/cctv.xml", func(c *gin.Context) {
		var epgList = make([]string, 0)
		for _, c := range cctv {
			epgList = append(epgList, c.epg)
		}
		resp, err := fastGet(epgList)
		if err != nil {
			c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(err.Error()))
			return
		}
		fmt.Println("The winner is ", resp.URL)
		c.Data(http.StatusOK, resp.contentType, resp.body)
	})
}

func start(srv *http.Server) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()

	log.Printf("Start Server @ %s", srv.Addr)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Print("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown:%s", err)
	}
	<-ctx.Done()
	log.Print("Server exiting")
}

func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func failed(msg string) gin.H {
	return gin.H{
		"msg":       msg,
		"timestamp": time.Now().Unix(),
	}
}

func data(data interface{}) gin.H {
	return gin.H{
		"msg":       "success",
		"data":      data,
		"timestamp": time.Now().Unix(),
	}
}
