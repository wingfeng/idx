package handlers

import (
	"net/http"
)

func RevocateHandler(w http.ResponseWriter, r *http.Request) {
	_, err := verifyAuthorizationToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token := r.URL.Query().Get("token")
	Srv.Manager.RemoveAccessToken(r.Context(), token)

	w.Write([]byte("Token removed"))
}
