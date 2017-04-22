package conf

// Paths for TLS certificates and keys
const (
	// Domain: "forcamp.ga"
	MAIN_SITE_TLS = "./conf/tls/forcamp.pem"
	MAIN_SITE_TLS_KEY = "./conf/tls/forcamp_key.pem"
	// Domain: "www.forcamp.ga"
	WWW_MAIN_SITE_TLS = "./conf/tls/wwwforcamp.pem"
	WWW_MAIN_SITE_TLS_KEY = "./conf/tls/wwwforcamp_key.pem"
	// Domain: "api.forcamp.ga"
	API_SITE_TLS = "./conf/tls/apiforcamp.pem"
	API_SITE_TLS_KEY = "./conf/tls/apiforcamp_key.pem"
)
