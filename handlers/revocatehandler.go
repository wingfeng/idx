package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RevocateHandler(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer

	_, err := verifyAuthorizationToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token := r.URL.Query().Get("token")
	Srv.Manager.RemoveAccessToken(r.Context(), token)

	w.Write([]byte("Token removed"))
}
