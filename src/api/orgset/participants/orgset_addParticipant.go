package participants

import (
	"net/http"
	"forcamp/src/api/orgset"
	"forcamp/conf"
	"forcamp/src"
	"strconv"
	"github.com/tealeg/xlsx"
)

type addParticipant_Success struct{
	Login string `json:"login"`
}

func AddParticipant(token string, participant Participant, responseWriter http.ResponseWriter) bool {
	if orgset.CheckUserAccess(token, responseWriter){
		Organization, _, APIerr := orgset.GetUserOrganizationAndLoginByToken(token)
		if APIerr != nil {
			return APIerr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(Organization)
		if checkAddParticipantData(participant, responseWriter) {
			rawResp, APIerr := addParticipantRequest(participant, Organization)
			if APIerr != nil {
				return APIerr.Print(responseWriter)
			}
			resp := conf.ApiResponse{200, "success", rawResp}
			resp.Print(responseWriter)
		}
	}
	return true
}

func addParticipantRequest(participant Participant, organization string) (addParticipant_Success, *conf.ApiResponse){
	Password, Hash := orgset.GeneratePassword()
	login, APIerr := addParticipant_Main(organization, Hash)
	if APIerr != nil {
		return addParticipant_Success{}, APIerr
	}
	participant.Login = login
	APIerr = addParticipant_Organization(participant)
	if APIerr != nil {
		return addParticipant_Success{}, APIerr
	}
	APIerr = addParticipant_Excel(participant, organization, Password)
	if APIerr != nil {
		return addParticipant_Success{}, APIerr
	}
	return addParticipant_Success{login}, nil
}

func addParticipant_Main(organization string, hash string) (string, *conf.ApiResponse){
	Query, err := src.Connection.Prepare("INSERT INTO users(password,organization) VALUES(?,?)")
	if err != nil {
		return "", conf.ErrDatabaseQueryFailed
	}
	Resp, err := Query.Exec(hash, organization)
	if err != nil {
		return "", conf.ErrDatabaseQueryFailed
	}
	ID, err := Resp.LastInsertId()
	if err != nil {
		return "", conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	Query, err = src.Connection.Prepare("UPDATE users SET login=? WHERE id=?")
	if err != nil {
		return "", conf.ErrDatabaseQueryFailed
	}
	login := "participant_"+strconv.FormatInt(ID, 10)
	_, err = Query.Exec(login, ID)
	if err != nil {
		return "", conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	return login, nil
}

func addParticipant_Organization(participant Participant) *conf.ApiResponse{
	Query, err := src.CustomConnection.Prepare("INSERT INTO users(login,name,surname,middlename,team,access,sex,avatar) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(participant.Login, participant.Name, participant.Surname, participant.Middlename, participant.Team, 0, participant.Sex, "default.jpg")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	Query, err = src.CustomConnection.Prepare("INSERT INTO participants(login) VALUES(?)")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = Query.Exec(participant.Login)
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	Query.Close()
	return nil
}

func addParticipant_Excel(participant Participant, organization string, password string) *conf.ApiResponse{
	teamName, APIerr := getTeamNameById(participant.Team)
	if APIerr != nil {
		return APIerr
	}
	excelFilePath := conf.FOLDER_PARTICIPANTS+"/"+organization+".xlsx"
	xlFile, err := xlsx.OpenFile(excelFilePath)
	if err != nil {
		return conf.ErrOpenExcelFile
	}
	sheet := xlFile.Sheets[0]
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = participant.Surname
	cell = row.AddCell()
	cell.Value = participant.Name
	cell = row.AddCell()
	cell.Value = participant.Middlename
	cell = row.AddCell()
	cell.Value = teamName
	cell = row.AddCell()
	cell.Value = participant.Login
	cell = row.AddCell()
	cell.Value = password
	err = xlFile.Save(excelFilePath)
	if err != nil {
		return conf.ErrSaveExcelFile
	}
	return nil
}

func getTeamNameById(id int64) (string, *conf.ApiResponse){
	if id == 0{
		return "отуствует", nil
	} else {
		Query, err := src.CustomConnection.Query("SELECT name FROM teams WHERE id=?", id)
		if err != nil {
			return "", conf.ErrDatabaseQueryFailed
		}
		defer Query.Close()
		var name string
		for Query.Next(){
			err = Query.Scan(&name)
			if err != nil {
				return "", conf.ErrDatabaseQueryFailed
			}
		}
		return name, nil
	}

}

func checkAddParticipantData(participant Participant, w http.ResponseWriter) bool {
	if len(participant.Name) > 0 {
		if len(participant.Surname) > 0 {
			if len(participant.Middlename) > 0 {
				if participant.Sex == 0 || participant.Sex == 1 {
					if orgset.CheckTeamID(participant.Team, w) {
						return true
					} else {
						return false
					}
				} else {
					return conf.ErrSexIncorrect.Print(w)
				}
			} else {
				return conf.ErrMiddlenameEmpty.Print(w)
			}
		} else {
			return conf.ErrSurnameEmpty.Print(w)
		}
	} else {
		return conf.ErrNameEmpty.Print(w)
	}
}

