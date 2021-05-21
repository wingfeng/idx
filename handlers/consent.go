package handlers

import (
	"net/http"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
)

func Consent(ctx *gin.Context) {
	w := ctx.Writer
	r := ctx.Request

	store, err := session.Start(ctx.Request.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	var form map[string]interface{}
	if v, ok := store.Get("ReturnUri"); ok {
		// mapVal := v.(map[string]interface{})
		// for m, val := range mapVal {
		// 	form[m] = val.(string)
		// }
		form = v.(map[string]interface{})
	}
	if v, ok := store.Get("state"); ok {
		//r.Form.Set("state", v.(string))
		log.Infof("State:%s", v)
	}

	ctx.HTML(http.StatusOK, "consent.html", form)

}
