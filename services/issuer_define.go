package services

const (
	issuer                 = "http://localhost:8080"
	authorization_endpoint = "http://localhost:8080/authorize"
	token_endpoint         = "http://localhost:8080/token"
	userinfo_endpoint      = "http://localhost:8080/userinfo"
	jwks_uri               = "http://localhost:8080/.well-known/jwks.json"
	end_session_endpoint   = "http://localhost:8080/logout"
	//response_types_supported = ["code"]
	//grant_types_supported =  ["authorization_code", "refresh_token"]
	//"subject_types_supported" =  ["public"]
	//"id_token_signing_alg_values_supported": ["RS256"],
	//"scopes_supported": ["openid", "profile", "email"],
	//"token_endpoint_auth_methods_supported": ["client_secret_basic", "client_secret_post", "none"],
	//"code_challenge_methods_supported": ["S256"]
)

type IssuerStruct struct {
	Issuer                            string   `json:"issuer"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
	JwksUri                           string   `json:"jwks_uri"`
	EndSessionEndpoint                string   `json:"end_session_endpoint"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	GrantTypesSupported               []string `json:"grant_types_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	IdTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	ScopsSupported                    []string `json:"scops_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	CodeChallengeMethodsSupported     []string `json:"code_challenge_methods_supported"`
}

func NewIssuer() IssuerStruct {
	issuer := IssuerStruct{
		Issuer:                            issuer,
		AuthorizationEndpoint:             authorization_endpoint,
		TokenEndpoint:                     token_endpoint,
		UserinfoEndpoint:                  userinfo_endpoint,
		JwksUri:                           jwks_uri,
		EndSessionEndpoint:                end_session_endpoint,
		ResponseTypesSupported:            []string{"code"},
		GrantTypesSupported:               []string{"authorization_code", "refresh_token"},
		SubjectTypesSupported:             []string{"public"},
		IdTokenSigningAlgValuesSupported:  []string{"RS256"},
		ScopsSupported:                    []string{"openid", "profile", "email"},
		TokenEndpointAuthMethodsSupported: []string{"client_secret_basic", "client_secret_post", "none"},
		CodeChallengeMethodsSupported:     []string{"S256"},
	}
	return issuer
}
