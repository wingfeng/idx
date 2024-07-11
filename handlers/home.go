package handlers

import "github.com/gin-gonic/gin"

func Home(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{
		"title":    "IDX Home",
		"userName": "Wing",
	})

}
