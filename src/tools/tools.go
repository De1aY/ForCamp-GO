package tools

import (
	"forcamp/src"
	"strings"
)

func CheckToken(token string) bool {
	var count int
	err := src.Connection.QueryRow("SELECT COUNT(login) FROM sessions WHERE token=?", token).Scan(&count)
	if err != nil {
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
	if len(data) > 2 {
		if strings.IndexFunc(data[:1], IsLatinLetter) != -1 {
			return strings.ToUpper(data[:2]) + data[2:]
		} else {
			return strings.ToUpper(data[:1]) + data[1:]
		}
	}
	return data
}

func TimestampToDate(t string) string {
	date := strings.Split(t, " ")
	ymd := strings.Split(date[0], "-")
	return ymd[2] + "." + ymd[1] + "." + ymd[0]
}

func IsNegative (val int64) bool {
	return val < 0
}