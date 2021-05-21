package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wingfeng/idx/core"
)

func WellknownHandler(ctx *gin.Context) {

	r := ctx.Request

	ctx.Header("Content-Type", "application/json")
	config := &core.OpenIDConfig{}
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	config.Issuer = fmt.Sprintf("%s://%s", scheme, r.Host)
	config.JwksURI = fmt.Sprintf("%s%s", config.Issuer, "/.well-known/openid-configuration/jwks")
	config.AuthorizationEndpoint = fmt.Sprintf("%s%s", config.Issuer, "/connect/authorize")
	config.TokenEndpoint = fmt.Sprintf("%s%s", config.Issuer, "/connect/token")
	config.UserInfoEndpoint = fmt.Sprintf("%s%s", config.Issuer, "/connect/userinfo")
	config.EndSessionEndpoint = fmt.Sprintf("%s%s", config.Issuer, "/connect/endsession")
	config.CheckSessionIframe = fmt.Sprintf("%s%s", config.Issuer, "/connect/checksession")
	config.RevocationEndpoint = fmt.Sprintf("%s%s", config.Issuer, "/connect/revocation")
	config.IntrospectionEndpoint = fmt.Sprintf("%s%s", config.Issuer, "/connect/intropect")
	config.DeviceAuthorizationEndpoint = fmt.Sprintf("%s%s", config.Issuer, "/connect/deviceauthorization")
	config.FrontchannelLogoutSupported = true
	config.FrontchannelLogoutSessionSupported = true
	config.BackchannelLogoutSupported = true
	config.BackchannelLogoutSessionSupported = true
	config.TokenEndpointAuthMethodsSupported = []string{"client_secret_basic", "client_secret_post"}
	config.GrantTypesSupported = []string{"authorization_code", "implicit", "refresh_token"}
	config.ResponseTypesSupported = []string{"code",
		"token",
		"id_token",
		"id_token token",
		"code id_token",
		"code token",
		"code id_token token"}
	config.CodeChallengeMethodsSupported = []string{"S256", "plain"}
	config.IDTokenSigningAlgValuesSupported = []string{"RS256"}
	config.SubjectTypesSupported = []string{"public"}
	config.RequestParameterSupported = true
	// e := json.NewEncoder(w)
	// e.SetIndent("", "  ")
	// e.Encode(config)
	ctx.JSON(200, config)
}
