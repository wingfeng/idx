package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/wingfeng/idx/core"
)

var Jwks *core.JWKS

func JWKSHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(Jwks)
}
