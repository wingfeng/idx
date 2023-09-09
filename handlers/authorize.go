package handlers

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
)

func Authorize(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request
	//	core.DumpRequest(os.Stdout, "Authorize", r)
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var form url.Values

	if v, ok := store.Get("ReturnUri"); ok {
		if r.Form == nil {
			r.Form = make(url.Values)
		}
		form = v.(url.Values)
		for m, val := range form {
			for _, vi := range val {
				r.Form.Set(m, vi)
			}

		}
	}

	store.Delete("ReturnUri")
	store.Save()

	err = Srv.HandleAuthorizeRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
