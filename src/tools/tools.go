package tools

import (
	"forcamp/src"
	"log"
	"strings"
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

func StringToBoolean(data string) bool {
	if data == "true" {
		return true
	} else {
		return false
	}
}

func IsLatinLetter (r rune) bool{
	return r < 'A' || r > 'z'
}

func ToTitleCase(data string) string {
	if strings.IndexFunc(data[:1], IsLatinLetter) != -1 {
		return strings.ToUpper(data[:2])+data[2:]
	} else {
		return strings.ToUpper(data[:1])+data[1:]
	}
}