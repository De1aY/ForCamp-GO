package tools

import (
	"forcamp/src"
	"log"
)

func CheckToken(token string) bool {
	var count int
	err := src.Connection.QueryRow("SELECT COUNT(login) FROM sessions WHERE token=?", token).Scan(&count)
	if err != nil {
		log.Print(err)
		return false
	}
	if count > 0 {
		return true
	}
	return false
}