package handlers

import (
	"net/http"
)

func Token(w http.ResponseWriter, r *http.Request) {

	err := Srv.HandleTokenRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
