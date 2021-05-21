package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idx/core"
)

func TokenController(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	core.DumpRequest(os.Stdout, "Token", r)
	err := Srv.HandleTokenRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
