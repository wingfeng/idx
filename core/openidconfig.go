package core

type OpenIDConfig struct {
	Issuer                             string   `json:"issuer"`
	JwksURI                            string   `json:"jwks_uri"`
	AuthorizationEndpoint              string   `json:"authorization_endpoint"`
	TokenEndpoint                      string   `json:"token_endpoint"`
	UserInfoEndpoint                   string   `json:"userinfo_endpoint"`
	EndSessionEndpoint                 string   `json:"end_session_endpoint"`
	CheckSessionIframe                 string   `json:"check_session_iframe"`
	RevocationEndpoint                 string   `json:"revocation_endpoint"`
	IntrospectionEndpoint              string   `json:"introspection_endpoint"`
	DeviceAuthorizationEndpoint        string   `json:"device_authorization_endpoint"`
	FrontchannelLogoutSupported        bool     `json:"frontchannel_logout_supported"`
	FrontchannelLogoutSessionSupported bool     `json:"frontchannel_logout_session_supported"`
	BackchannelLogoutSupported         bool     `json:"backchannel_logout_supported"`
	BackchannelLogoutSessionSupported  bool     `json:"backchannel_logout_session_supported"`
	ScopesSupported                    []string `json:"scopes_supported"`
	ClaimsSupported                    []string `json:"claims_supported"`
	GrantTypesSupported                []string `json:"grant_types_supported"`
	ResponseTypesSupported             []string `json:"response_types_supported"`
	ResponseModesSupported             []string `json:"response_modes_supported"`
	TokenEndpointAuthMethodsSupported  []string `json:"token_endpoint_auth_methods_supported"`
	IDTokenSigningAlgValuesSupported   []string `json:"id_token_signing_alg_values_supported"`
	SubjectTypesSupported              []string `json:"subject_types_supported"`
	CodeChallengeMethodsSupported      []string `json:"code_challenge_methods_supported"`
	RequestParameterSupported          bool     `json:"request_parameter_supported"`
}
