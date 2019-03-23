package participants

import (
	"net/http"
	"wplay/src/api/orgset"
	"wplay/conf"
	"wplay/src"
	"strconv"
	"github.com/tealeg/xlsx"
)

type addParticipant_Success struct{
	ID int64 `json:"id"`
}

func AddParticipant(token string, participant Participant, responseWriter http.ResponseWriter) bool {
	if orgset.IsUserAdmin(token, responseWriter){
		organizationName, _, apiErr := orgset.GetUserOrganizationAndIdByToken(token); if apiErr != nil {
			return apiErr.Print(responseWriter)
		}
		src.CustomConnection = src.Connect_Custom(organizationName)
		if isAddParticipantDataValid(participant, responseWriter) {
			rawResp, apiErr := addParticipant(participant, organizationName); if apiErr != nil {
				return apiErr.Print(responseWriter)
			}
			resp := conf.ApiResponse{200, "success", rawResp}
			resp.Print(responseWriter)
		}
	}
	return true
}

func addParticipant(participant Participant, organization string) (addParticipant_Success, *conf.ApiResponse){
	password, hash := orgset.GeneratePassword()
	participant_id, participant_login, apiErr := addParticipant_Main(organization, hash); if apiErr != nil {
		return addParticipant_Success{}, apiErr
	}
	participant.ID = participant_id
	apiErr = addParticipant_Organization(participant); if apiErr != nil {
		return addParticipant_Success{}, apiErr
	}
	apiErr = addParticipant_Excel(participant, participant_login,organization, password); if apiErr != nil {
		return addParticipant_Success{}, apiErr
	}
	return addParticipant_Success{participant_id}, nil
}

func addParticipant_Main(organizationName string, hash string) (int64, string, *conf.ApiResponse){
	query, err := src.Connection.Prepare("INSERT INTO users(password,organization) VALUES(?,?)"); if err != nil {
		return 0, "", conf.ErrDatabaseQueryFailed
	}
	resp, err := query.Exec(hash, organizationName); if err != nil {
		return 0, "", conf.ErrDatabaseQueryFailed
	}
	participant_id, err := resp.LastInsertId(); if err != nil {
		return 0, "", conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.Connection.Prepare("UPDATE users SET login=? WHERE id=?"); if err != nil {
		return 0, "", conf.ErrDatabaseQueryFailed
	}
	participant_login := "participant_"+strconv.FormatInt(participant_id, 10)
	_, err = query.Exec(participant_login, participant_id); if err != nil {
		return 0, "", conf.ErrDatabaseQueryFailed
	}
	query.Close()
	return participant_id, participant_login, nil
}

func addParticipant_Organization(participant Participant) *conf.ApiResponse{
	query, err := src.CustomConnection.Prepare("INSERT INTO users(id,name,surname,middlename," +
		"team,access,sex,avatar) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(participant.ID, participant.Name, participant.Surname,
		participant.Middlename, participant.Team, 0, participant.Sex, "default.jpg")
	if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	query, err = src.CustomConnection.Prepare("INSERT INTO participants(id) VALUES(?)"); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	_, err = query.Exec(participant.ID); if err != nil {
		return conf.ErrDatabaseQueryFailed
	}
	query.Close()
	return nil
}

func addParticipant_Excel(participant Participant, participant_login string, organization string, password string) *conf.ApiResponse{
	teamName, APIerr := getTeamNameById(participant.Team); if APIerr != nil {
		return APIerr
	}
	excelFilePath := conf.FOLDER_PARTICIPANTS + "/" + organization + ".xlsx"
	xlFile, err := xlsx.OpenFile(excelFilePath); if err != nil {
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
	cell.Value = participant_login
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
		rows, err := src.CustomConnection.Query("SELECT name FROM teams WHERE id=?", id); if err != nil {
			return "", conf.ErrDatabaseQueryFailed
		}
		defer rows.Close()
		var name string
		for rows.Next(){
			err = rows.Scan(&name); if err != nil {
				return "", conf.ErrDatabaseQueryFailed
			}
		}
		return name, nil
	}

}

func isAddParticipantDataValid(participant Participant, w http.ResponseWriter) bool {
	if len(participant.Name) > 0 {
		if len(participant.Surname) > 0 {
			if len(participant.Middlename) > 0 {
				if participant.Sex == 0 || participant.Sex == 1 {
					if orgset.IsTeamExist(participant.Team, w) {
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

