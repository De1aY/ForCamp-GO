package authorization

import (
	"crypto/sha512"
	"encoding/base64"
	"database/sql"
	"forcamp/conf"
	"net/http"
	"time"
	"math/rand"
	"strconv"
	"encoding/json"
	"fmt"
)

const HASH_SALT = "ef203nsd313"

type AuthInf struct {
	Login    string
	Password string
}

type Success struct {
	Code int `json:"code"`
	Status string `json:"status"`
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
	Bytes := []byte(login + Time + strconv.FormatInt(rand.Int63(), 10))
	Result := Hash.Sum(Bytes)
	return base64.URLEncoding.EncodeToString(Result)
}

func getCountVal(rows *sql.Rows, ResponseWriter http.ResponseWriter) (count int) {
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			conf.PrintError(conf.ErrDatabaseQueryFailed, ResponseWriter)
		}
	}
	return count
}

func printToken(token string, w http.ResponseWriter) bool {
	rawResp := &Success{200, "success", token}
	Response, _ := json.Marshal(rawResp)
	fmt.Fprintf(w, string(Response))
	return true
}