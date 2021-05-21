package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
)

func LogoutHandler(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	store, err := session.Start(ctx.Request.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	store.Delete("LoggedInUserID")
	store.Save()

	http.Redirect(w, r, "/login", http.StatusFound)

}
