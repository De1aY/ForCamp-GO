/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package src

import (
	"wplay/conf"
	"net/http"
)

/*
SetHeaders_API writes security HTTP-headers (X-XSS-Protection, Referrer-Policy, e.t.c) for API site (GET requests)
*/
func SetHeaders_API_GET(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Security-Policy", "default-src 'none'; img-src 'self'; script-src 'self'; style-src 'unsafe-inline'")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
	w.Header().Set("Referrer-Policy", "no-referrer")
}

/*
SetHeaders_API writes security HTTP-headers (X-XSS-Protection, Referrer-Policy, e.t.c) for API site (POST-requests)
*/
func SetHeaders_API_POST(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Security-Policy", "default-src 'none'; img-src 'self'; script-src 'self'; style-src 'unsafe-inline'")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
	w.Header().Set("Referrer-Policy", "no-referrer")
}

/*
SetHeaders_Main writes security HTTP-headers (X-XSS-Protection, Content-Security-Policy, e.t.c) for Main site
*/
func SetHeaders_Main(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Method", "GET")
	w.Header().Set("Content-Security-Policy",
		"default-src 'self' "+conf.API_SITE_DOMAIN+" https://mc.yandex.ru/ https://googleads.g.doubleclick.net "+
			"https://content.googleapis.com/ https://accounts.google.com/; "+
			"font-src 'self' https://fonts.gstatic.com https://cdnjs.cloudflare.com/; "+
			"img-src 'self' https://www.google.com/ https://stats.g.doubleclick.net/ https://www.google-analytics.com/ https://mc.yandex.ru/ data:; "+
			"script-src 'unsafe-eval' 'unsafe-inline' 'self' https://yastatic.net/ https://www.google-analytics.com/analytics.js "+
			"https://mc.yandex.ru/ https://cdnjs.cloudflare.com https://pagead2.googlesyndication.com "+
			"https://apis.google.com ;"+
			"style-src 'self' 'unsafe-inline' https://fonts.googleapis.com https://cdnjs.cloudflare.com")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
}

/*
SetHeaders_API_Download security HTTP-headers (X-XSS-Protection, Referrer-Policy, e.t.c) for Download pages
*/
func SetHeaders_API_Download(w http.ResponseWriter, filename string, r *http.Request) {
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
