package users

import (
	"forcamp/conf"
	"forcamp/src"
	"forcamp/src/api/authorization"
	"forcamp/src/api/orgset"
	"io"
	"net/http"
	"os"
	"strconv"
)

func ChangeAvatar(token string, request *http.Request, responseWriter http.ResponseWriter) bool {
	if authorization.IsTokenValid(token, responseWriter) {
		organization_name, user_id, apiErr := orgset.GetUserOrganizationAndIdByToken(token)
		if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(organization_name)
		response := changeAvatar(user_id, request)
		return response.Print(responseWriter)
	}
	return true
}

func changeAvatar(user_id int64, request *http.Request) *conf.ApiResponse {
	var err error
	err = request.ParseMultipartForm(2048)
	if err != nil {
		return conf.ErrFileUpload
	}
	file, _, err := request.FormFile("file")
	if err != nil {
		return conf.ErrFileUpload
	}
	defer file.Close()
	fileName := "user_" + strconv.FormatInt(user_id, 10) + ".png"
	f, err := os.OpenFile(conf.FOLDER_IMAGES+"/"+fileName,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return conf.ErrFileUpload
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		return conf.ErrFileUpload
	}
	req, err := src.CustomConnection.Prepare("UPDATE users SET avatar=? WHERE id=?")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	defer req.Close()
	_, err = req.Exec(&fileName)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return conf.RequestSuccess
}