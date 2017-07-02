package tests

import (
	"forcamp/src/api/handlers"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"forcamp/src/api/authorization"
	"forcamp/conf"
)

func TestLoginAndPasswordHandler(t *testing.T) {
	Request, err := http.NewRequest("GET", "/token.get?login=test_1&password=test", nil)
	if err != nil {
		t.Fatal(err)
	}
	Recorder := httptest.NewRecorder()
	Handler := http.HandlerFunc(handlers.LoginAndPasswordAuthHandler)
	Handler.ServeHTTP(Recorder, Request)
	if status := Recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", status, http.StatusOK)
	}
	var res authorization.Success
	err = json.Unmarshal(Recorder.Body.Bytes(), &res)
	if err != nil {
		t.Errorf(err.Error())
	}
	if res.Status != "success"{
		t.Errorf(string(res.Code))
	}
	t.Logf("Authorization successful!")
}

func TestTokenVerification(t *testing.T){
	Request, err := http.NewRequest("GET", "/token.verify?token=dGVzdF8xMTQ4OTQzMDk4NTU1NzcwMDY3OTE5NDc3Nzk0MTDPg-E1fu-4vfFUKFDWbYAH1iDkBQtXFdyD9Kkh02zpzkfQ0TxdhfKw_4MY0od-7C9juTG9R0F6gaU4Mnr5J9o-", nil)
	if err != nil{
		t.Fatal(err)
	}
	Recorder := httptest.NewRecorder()
	Handler := http.HandlerFunc(handlers.TokenVerificationHandler)
	Handler.ServeHTTP(Recorder, Request)
	if status := Recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", status, http.StatusOK)
	}
	var res conf.Success
	err = json.Unmarshal(Recorder.Body.Bytes(), &res)
	if err != nil {
		t.Errorf(err.Error())
	}
	if res.Status != "success"{
		t.Errorf(string(res.Code))
	}
	t.Logf("Token verification successful!")
}