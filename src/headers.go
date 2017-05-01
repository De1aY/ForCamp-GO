/*
	Copyright: "Null team", 2016 - 2017
	Author: "De1aY"
	Documentation: https://bitbucket.org/lyceumdevelopers/golang/wiki/Home
*/
package src

import (
	"net/http"
)

/*
SetHeaders_API writes security HTTP-headers (X-XSS-Protection, Referrer-Policy, e.t.c) for API site
 */
func SetHeaders_API(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET, POST")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Security-Policy", "default-src 'none'; img-src 'self'; script-src 'self'; style-src 'unsafe-inline'")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
	w.Header().Set("Referrer-Policy", "no-referrer")
}

/*
SetHeaders_Main writes security HTTP-headers (X-XSS-Protection, Referrer-Policy, e.t.c) for Main site
 */
func SetHeaders_Main(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Method", "GET")
	w.Header().Set("Content-Security-Policy", "default-src 'self' https://api.forcamp.ga; font-src 'self' https://fonts.gstatic.com https://cdnjs.cloudflare.com/; img-src 'self' data:; script-src 'unsafe-eval' 'self' https://www.google-analytics.com/analytics.js https://cdnjs.cloudflare.com 'unsafe-inline'; style-src 'unsafe-inline' 'self' https://fonts.googleapis.com https://cdnjs.cloudflare.com")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
	w.Header().Set("Referrer-Policy", "no-referrer")
}

/*
SetHeaders_API_Download security HTTP-headers (X-XSS-Protection, Referrer-Policy, e.t.c) for Download pages
 */
func SetHeaders_API_Download(w http.ResponseWriter, filename string, r *http.Request){
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Security-Policy", "default-src 'none'; img-src 'self'; script-src 'self'; style-src 'unsafe-inline'")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
	w.Header().Set("Referrer-Policy", "no-referrer")
}