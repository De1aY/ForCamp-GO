package authorization

import (
	"crypto/sha512"
	"encoding/base64"
	"wplay/conf"
	"net/http"
	"time"
	"math/rand"
	"strconv"
)

const HASH_SALT = "ef203nsd313"

type AuthInf struct {
	Login    string
	Password string
}

type authorizationSuccess struct {
	Token string `json:"token"`
}

// Generate PasswordHash from string using SHA-512(Keccak)
func GeneratePasswordHash(password string) string {
	Hash := sha512.New()
	Bytes := []byte(password + HASH_SALT)
	Result := Hash.Sum(Bytes)
	return base64.URLEncoding.EncodeToString(Result)
}

// Generate Token from string using SHA-512(Keccak)
func generateTokenHash(login string) string {
	Hash := sha512.New()
	Time := strconv.FormatInt(time.Now().Unix(), 10)
	rand.Seed(time.Now().UnixNano())
	Bytes := []byte(login + Time + strconv.FormatInt(rand.Int63(), 10))
	Result := Hash.Sum(Bytes)
	return base64.URLEncoding.EncodeToString(Result)
}

func printToken(token string, responseWriter http.ResponseWriter) bool {
	resp := conf.ApiResponse{200, "success", authorizationSuccess{token}}
	resp.Print(responseWriter)
	return true
}

func IsTokenNotEmpty(token string, responseWriter http.ResponseWriter) bool{
	if len(token) > 0{
		return true
	} else {
		return conf.ErrUserTokenEmpty.Print(responseWriter)
	}
}
