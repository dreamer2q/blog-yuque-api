package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	g *gin.Engine
)

func init() {
	g = gin.Default()

	//允许跨域调用
	g.Use(cors.Default())

	initRoute()
}

func initRoute() {
	api := g.Group("/api")
	{
		api.GET("articles", func(c *gin.Context) {
			articles, err := ProxyArticleList()
			if err != nil {
				MakeFail(c, http.StatusInternalServerError, err)
				return
			}
			MakeSuccess(c, articles)
		})
		api.GET("article/:slug", func(c *gin.Context) {
			slug := c.Param("slug")
			article, err := ProxyArticleDetail(slug)
			if err != nil {
				MakeFail(c, http.StatusInternalServerError, err)
				return
			}
			MakeSuccess(c, article)
		})
	}
}

func Serve() {
	if err := g.Run(config.Addr); err != nil {
		panic(err)
	}
	fmt.Printf("terminated")
}
