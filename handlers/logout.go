package handlers

import (
	"net/http"

	"github.com/go-session/session"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	store.Delete("LoggedInUserID")
	store.Save()

	http.Redirect(w, r, "/login", http.StatusFound)

}
