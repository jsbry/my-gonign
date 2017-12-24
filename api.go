package main

import "github.com/gin-gonic/gin"

func Api(c *gin.Context) {
	c.HTML(200, "Api", gin.H{
		"title": "ダッシュボード",
		"vuejs": "api.js",
	})
	return
}
