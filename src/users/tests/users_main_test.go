package tests

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"forcamp/src/handlers"
	"encoding/json"
	"forcamp/src/users"
)

func TestGetUserLogin(t *testing.T){
	Request, err := http.NewRequest("GET", "/user.login?token=dGVzdF8xMTQ4OTQzMDk4NTU1NzcwMDY3OTE5NDc3Nzk0MTDPg-E1fu-4vfFUKFDWbYAH1iDkBQtXFdyD9Kkh02zpzkfQ0TxdhfKw_4MY0od-7C9juTG9R0F6gaU4Mnr5J9o-", nil)
	if err != nil{
		t.Fatal(err)
	}
	Recorder := httptest.NewRecorder()
	Handler := http.HandlerFunc(handlers.GetUserLoginHandler)
	Handler.ServeHTTP(Recorder, Request)
	if status := Recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", status, http.StatusOK)
	}
	var res users.Success_GetUserLogin
	err = json.Unmarshal(Recorder.Body.Bytes(), &res)
	if err != nil {
		t.Errorf(err.Error())
	}
	if res.Status != "success"{
		t.Fatal(string(res.Status))
	}
	t.Logf("Login(%v) getting successful!", res.Login)
}

func TestGetUserData(t *testing.T){
	Request, err := http.NewRequest("GET", "/user.data?token=dGVzdF8xMTQ4OTQzMDk4NTU1NzcwMDY3OTE5NDc3Nzk0MTDPg-E1fu-4vfFUKFDWbYAH1iDkBQtXFdyD9Kkh02zpzkfQ0TxdhfKw_4MY0od-7C9juTG9R0F6gaU4Mnr5J9o-&login=fortest_1", nil)
	if err != nil{
		t.Fatal(err)
	}
	Recorder := httptest.NewRecorder()
	Handler := http.HandlerFunc(handlers.GetUserDataHandler)
	Handler.ServeHTTP(Recorder, Request)
	if status := Recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code : got %v want %v", status, http.StatusOK)
	}
	var res users.Success_GetUserData
	err = json.Unmarshal(Recorder.Body.Bytes(), &res)
	if err != nil {
		t.Errorf(err.Error())
	}
	if res.Status != "success"{
		t.Fatal(string(res.Status))
	}
	t.Logf("Data(%v %v) getting successful!", res.UserData.Surname, res.UserData.Name)
}