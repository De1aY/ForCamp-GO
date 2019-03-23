package users

import (
	"bytes"
	"encoding/base64"
	"wplay/conf"
	"wplay/src"
	"wplay/src/api/authorization"
	"wplay/src/api/orgset"
	"image/png"
	"io/ioutil"
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
	file, fileHeader, err := request.FormFile("file")
	if err != nil || fileHeader.Header.Get("Content-Type") != "image/png" {
		return conf.ErrFileUpload
	}
	defer file.Close()
	fileData, err := ioutil.ReadAll(file)
	rawAvatarImage := make([]byte, len(fileData))
	_, err = base64.StdEncoding.Decode(rawAvatarImage, fileData)
	rawAvatarImageReader := bytes.NewReader(rawAvatarImage)
	avatarImage, err := png.Decode(rawAvatarImageReader)
	fileName := "user_" + strconv.FormatInt(user_id, 10) + ".png"
	f, err := os.OpenFile(conf.FOLDER_IMAGES+"/"+fileName,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return conf.ErrFileUpload
	}
	defer f.Close()
	err = png.Encode(f, avatarImage)
	if err != nil {
		return conf.ErrFileUpload
	}
	req, err := src.CustomConnection.Prepare("UPDATE users SET avatar=? WHERE id=?")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	defer req.Close()
	_, err = req.Exec(&fileName, &user_id)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	return conf.RequestSuccess
}
