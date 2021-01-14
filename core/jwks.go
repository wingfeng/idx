package core

type JWKS struct {
	Keys []interface{} `json:"keys"`
}
type JWTKey struct {
	KeyType string `json:"kty"`
	Use     string `json:"use"`
	Kid     []byte `json:"kid"`
	//	X5t string `json:"x5t"`
	//	E   string `json:"e"`
	//	N   string `json:"n"`
	//	X5c string `json:"x5c"`
	Alg string `json:"alg"`
}

type RSAJWTKey struct {
	JWTKey
	E string `json:"e"` //The "e" (exponent) parameter contains the exponent value for the RSA	public key.
	N string `json:"n"` //The "n" (modulus) parameter contains the modulus value for the RSA public key.  It is represented as a Base64urlUInt-encoded value.
}

func NewRSAJWTKey() RSAJWTKey {
	key := RSAJWTKey{}
	key.JWTKey.KeyType = "RSA"
	key.Use = "sig"
	return key
}
