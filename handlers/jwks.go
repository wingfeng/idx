package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/wingfeng/idx/core"
)

var Jwks *core.JWKS

func JWKSHandler(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(Jwks)
}
