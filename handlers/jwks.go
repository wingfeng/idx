package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idx/core"
)

var Jwks *core.JWKS

func JWKSHandler(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Content-Type", "application/json")
	// e := json.NewEncoder(w)
	// e.SetIndent("", "  ")
	// e.Encode(Jwks)
	ctx.JSON(200, Jwks)
}
