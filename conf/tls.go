/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package conf

// Paths for TLS certificates and keys
const (
	// Domain: "main"
	MAIN_SITE_TLS     = "./conf/tls/wplay_tls.pem"
	MAIN_SITE_TLS_KEY = "./conf/tls/wplay_tls_key.pem"
	/* Domain: "www.forcamp.nullteam.info"
	WWW_MAIN_SITE_TLS = "./conf/tls/wwwforcamp.pem"
	WWW_MAIN_SITE_TLS_KEY = "./conf/tls/wwwforcamp_key.pem" */
	// Domain: "api"
	API_SITE_TLS     = "./conf/tls/api_wplay_tls.pem"
	API_SITE_TLS_KEY = "./conf/tls/api_wplay_tls_key.pem"
)
