package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenController(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	//	core.DumpRequest(os.Stdout, "Token", r)
	err := Srv.HandleTokenRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
