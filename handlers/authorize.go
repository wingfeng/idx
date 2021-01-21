package handlers

import (
	"net/http"
	"net/url"

	"github.com/go-session/session"
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var form map[string]interface{}

	if v, ok := store.Get("ReturnUri"); ok {
		if r.Form == nil {
			r.Form = make(url.Values)
		}
		form = v.(map[string]interface{})
		for m, val := range form {
			for _, vi := range val.([]interface{}) {
				r.Form.Set(m, vi.(string))
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
