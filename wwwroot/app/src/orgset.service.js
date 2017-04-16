"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
var __param = (this && this.__param) || function (paramIndex, decorator) {
    return function (target, key) { decorator(target, key, paramIndex); }
};
Object.defineProperty(exports, "__esModule", { value: true });
var core_1 = require("@angular/core");
var http_1 = require("@angular/http");
var notie_1 = require("notie");
var OrgSetService = (function () {
    function OrgSetService(http) {
        this.http = http;
        this.GetOrgSettingsLink = "https://api.forcamp.ga/orgset.settings.get";
        this.GetCategoriesLink = "https://api.forcamp.ga/orgset.categories.get";
        this.SetOrgSettingValueLink = "https://api.forcamp.ga/orgset.setting.edit";
        this.AddCategoryLink = "https://api.forcamp.ga/orgset.category.add";
        this.DeleteCategoryLink = "https://api.forcamp.ga/orgset.category.delete";
        this.EditCategoryLink = "https://api.forcamp.ga/orgset.category.edit";
        this.GetTeamsLink = "https://api.forcamp.ga/orgset.teams.get";
        this.EditTeamLink = "https://api.forcamp.ga/orgset.team.edit";
        this.AddTeamLink = "https://api.forcamp.ga/orgset.team.add";
        this.DeleteTeamLink = "https://api.forcamp.ga/orgset.team.delete";
        this.GetParticipantsLink = "https://api.forcamp.ga/orgset.participants.get";
        this.EditParticipantLink = "https://api.forcamp.ga/orgset.participant.edit";
        this.DeleteParticipantLink = "https://api.forcamp.ga/orgset.participant.delete";
        this.ResetParticipantPasswordLink = "https://api.forcamp.ga/orgset.participant.password.reset";
        this.AddParticipantLink = "https://api.forcamp.ga/orgset.participant.add";
        this.GetEmployeesLink = "https://api.forcamp.ga/orgset.employees.get";
        this.EditEmployeeLink = "https://api.forcamp.ga/orgset.employee.edit";
        this.DeleteEmployeeLink = "https://api.forcamp.ga/orgset.employee.delete";
        this.ResetEmployeePasswordLink = "https://api.forcamp.ga/orgset.employee.password.reset";
        this.AddEmployeeLink = "https://api.forcamp.ga/orgset.employee.add";
        this.EditEmployeePermissionLink = "https://api.forcamp.ga/orgset.employee.permission.edit";
        this.GetReasonsLink = "https://api.forcamp.ga/orgset.reasons.get";
        this.AddReasonLink = "https://api.forcamp.ga/orgset.reason.add";
        this.EditReasonLink = "https://api.forcamp.ga/orgset.reason.edit";
        this.DeleteReasonLink = "https://api.forcamp.ga/orgset.reason.delete";
        this.GetMarksChangesLink = "https://api.forcamp.ga/marks.changes.get";
        this.DeleteMarkChangeLink = "https://api.forcamp.ga/mark.change.delete";
        this.PostHeaders = new http_1.Headers();
        this.OrgSettings = {
            organization: "загрузка...",
            period: "загрузка...",
            participant: "загрузка...",
            self_marks: false,
            team: "загрузка..."
        };
        this.Teams = [];
        this.Categories = [];
        this.Participants = [];
        this.Employees = [];
        this.Reasons = [];
        this.MarksChanges = [];
        this.Preloader = false;
        this.ParticipantValueEdit_Active = false;
        this.PeriodValueEdit_Active = false;
        this.TeamValueEdit_Active = false;
        this.OrganizationValueEdit_Active = false;
        this.AddCategory_Active = false;
        this.AddTeam_Active = false;
        this.AddParticipant_Active = false;
        this.AddEmployee_Active = false;
        this.AddReason_Active = false;
        this.PostHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
    }
    OrgSetService.prototype.GetData = function () {
        var _this = this;
        if (this.UpdateInterval == undefined) {
            this.UpdateInterval = setInterval(function () { _this.GetData(); }, 20000);
        }
        this.GetOrgSettings();
        this.GetCategories();
        this.GetTeams();
        this.GetParticipants();
        this.GetEmployees();
        this.GetReasons();
        this.GetMarksChanges();
    };
    OrgSetService.prototype.GetOrgSettings = function () {
        var _this = this;
        this.PreloaderOn();
        this.http.get(this.GetOrgSettingsLink + "?token=" + this.Token).subscribe(function (data) { return _this.getOrgSettingsFromResponse(data.json()); });
    };
    OrgSetService.prototype.SetOrgSettingValue = function (name, value) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.SetOrgSettingValueLink, "token=" + this.Token + "&name=" + name + "&value=" + value, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkSetOrgSettingValueResponse(data.json(), name, value); });
    };
    OrgSetService.prototype.getOrgSettingsFromResponse = function (data) {
        if (data.code == 200) {
            this.OrgSettings = {
                organization: data.settings.organization,
                period: data.settings.period,
                participant: data.settings.participant,
                self_marks: this.StringToBoolean(data.settings.self_marks),
                team: data.settings.team
            };
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkSetOrgSettingValueResponse = function (data, name, value) {
        if (data.code == 200) {
            this.OrgSettings[name] = value;
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.GetCategories = function () {
        var _this = this;
        this.PreloaderOn();
        this.http.get(this.GetCategoriesLink + "?token=" + this.Token).subscribe(function (data) { return _this.getCategoriesFromResponse(data.json()); });
    };
    OrgSetService.prototype.AddCategory = function (name, negative_marks) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.AddCategoryLink, "token=" + this.Token + "&name=" + name + "&negative_marks=" + negative_marks, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkAddCategoryResponse(data.json(), name, negative_marks); });
    };
    OrgSetService.prototype.DeleteCategory = function (id) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.DeleteCategoryLink, "token=" + this.Token + "&id=" + id, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkDeleteCategoryResponse(data.json(), id); });
    };
    OrgSetService.prototype.EditCategory = function (category) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.EditCategoryLink, "token=" + this.Token + "&id=" + category.id + "&name=" + category.name + "&negative_marks=" + !category.negative_marks, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkEditCategoryResponse(data.json()); });
    };
    OrgSetService.prototype.checkAddCategoryResponse = function (data, name, negative_marks) {
        if (data.code == 200) {
            this.Categories.push({ id: data.id, name: name, negative_marks: negative_marks });
            if (this.Participants != undefined) {
                for (var i = 0; i < this.Participants.length; i++) {
                    this.Participants[i].marks.push({ id: data.id, value: 0 });
                }
            }
            if (this.Employees != undefined) {
                for (var i = 0; i < this.Employees.length; i++) {
                    this.Employees[i].permissions.push({ id: data.id, value: true });
                }
            }
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.getCategoriesFromResponse = function (data) {
        if (data.code == 200) {
            this.Categories = data.categories;
            if (this.Categories != null) {
                for (var i = 0; i < data.categories.length; i++) {
                    this.Categories[i].negative_marks = this.StringToBoolean(data.categories[i].negative_marks);
                }
            }
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkDeleteCategoryResponse = function (data, id) {
        if (data.code == 200) {
            for (var i = 0; i < this.Categories.length; i++) {
                if (this.Categories[i].id == id) {
                    this.Categories.splice(i, 1);
                    break;
                }
            }
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkEditCategoryResponse = function (data) {
        if (data.code == 200) {
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.GetTeams = function () {
        var _this = this;
        this.PreloaderOn();
        this.http.get(this.GetTeamsLink + "?token=" + this.Token).subscribe(function (data) { return _this.getTeamsFromResponse(data.json()); });
    };
    OrgSetService.prototype.AddTeam = function (name) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.AddTeamLink, "token=" + this.Token + "&name=" + name, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkAddTeamResponse(data.json(), name); });
    };
    OrgSetService.prototype.EditTeam = function (id, name) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.EditTeamLink, "token=" + this.Token + "&id=" + id + "&name=" + name, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkEditTeamResponse(data.json()); });
    };
    OrgSetService.prototype.DeleteTeam = function (id) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.DeleteTeamLink, "token=" + this.Token + "&id=" + id, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkDeleteTeamResponse(data.json(), id); });
    };
    OrgSetService.prototype.getTeamsFromResponse = function (data) {
        if (data.code == 200) {
            this.Teams = data.teams;
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkAddTeamResponse = function (data, name) {
        if (data.code == 200) {
            this.Teams.push({ id: data.id, name: name, participants: [], leader: { name: "", surname: "", middlename: "", login: "" } });
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkEditTeamResponse = function (data) {
        if (data.code == 200) {
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkDeleteTeamResponse = function (data, id) {
        if (data.code == 200) {
            for (var i = 0; i < this.Teams.length; i++) {
                if (this.Teams[i].id == id) {
                    this.Teams.splice(i, 1);
                    break;
                }
            }
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.GetParticipantsExcel = function () {
        window.location.href = "https://api.forcamp.ga/orgset.participants.password.get?token=" + this.Token;
    };
    OrgSetService.prototype.GetParticipants = function () {
        var _this = this;
        this.http.get(this.GetParticipantsLink + "?token=" + this.Token).subscribe(function (data) { return _this.getParticipantsFromResponse(data.json()); });
    };
    OrgSetService.prototype.EditParticipant = function (participant) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.EditParticipantLink, "token=" + this.Token +
            "&login=" + participant.login +
            "&name=" + participant.name +
            "&surname=" + participant.surname +
            "&middlename=" + participant.middlename +
            "&sex=" + participant.sex +
            "&team=" + participant.team, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkEditParticipantResponse(data.json(), participant); });
    };
    OrgSetService.prototype.DeleteParticipant = function (login) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.DeleteParticipantLink, "token=" + this.Token + "&login=" + login, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkDeleteParticipantResponse(data.json(), login); });
    };
    OrgSetService.prototype.ResetParticipantPassword = function (login) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.ResetParticipantPasswordLink, "token=" + this.Token + "&login=" + login, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkResetParticipantPasswordResponse(data.json()); });
    };
    OrgSetService.prototype.AddParticipant = function (participant) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.AddParticipantLink, "token=" + this.Token +
            "&name=" + participant.name +
            "&surname=" + participant.surname +
            "&middlename=" + participant.middlename +
            "&sex=" + participant.sex +
            "&team=" + participant.team, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkAddParticipantResponse(data.json(), participant); });
    };
    OrgSetService.prototype.GetParticipantFullNameByLogin = function (participant_login) {
        try {
            var data = this.Participants.filter(function (part) { return part.login == participant_login; })[0];
            return data.surname + " " + data.name + " " + data.middlename;
        }
        catch (e) {
            return participant_login;
        }
    };
    OrgSetService.prototype.GetParticipantsByTeamID = function (team_id) {
        if (team_id == -1) {
            return this.Participants;
        }
        else {
            return this.Participants.filter(function (part) { return part.team == team_id; });
        }
    };
    OrgSetService.prototype.checkResetParticipantPasswordResponse = function (data) {
        if (data.code == 200) {
            notie_1.alert({ type: 1, text: "Новый пароль: " + data.password, stay: true });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkDeleteParticipantResponse = function (data, login) {
        if (data.code == 200) {
            for (var i = 0; i < this.Participants.length; i++) {
                if (this.Participants[i].login == login) {
                    if (this.Participants[i].team != 0) {
                        for (var i_1 = 0; i_1 < this.Teams.length; i_1++) {
                            if (this.Teams[i_1].id == this.Participants[i_1].team) {
                                this.Teams[i_1].participants.splice(this.Teams[i_1].participants.indexOf(login), 1);
                                break;
                            }
                        }
                    }
                    this.Participants.splice(i, 1);
                    break;
                }
            }
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkEditParticipantResponse = function (data, participant) {
        if (data.code == 200) {
            for (var i = 0; i < this.Teams.length; i++) {
                if (this.Teams[i].participants.indexOf(participant.login) != -1) {
                    this.Teams[i].participants.splice(this.Teams[i].participants.indexOf(participant.login), 1);
                }
                if (this.Teams[i].id == participant.team) {
                    this.Teams[i].participants.push(participant.login);
                }
            }
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkAddParticipantResponse = function (data, participant) {
        if (data.code == 200) {
            participant.login = data.login;
            participant.marks = [];
            for (var i = 0; i < this.Categories.length; i++) {
                participant.marks.push({ id: this.Categories[i].id, value: 0 });
            }
            this.Participants.push(participant);
            if (participant.team != 0) {
                for (var i = 0; i < this.Teams.length; i++) {
                    if (this.Teams[i].id == participant.team) {
                        this.Teams[i].participants.push(participant.login);
                        break;
                    }
                }
            }
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.getParticipantsFromResponse = function (data) {
        if (data.code == 200) {
            this.Participants = data.participants;
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
    };
    OrgSetService.prototype.GetEmployees = function () {
        var _this = this;
        this.http.get(this.GetEmployeesLink + "?token=" + this.Token).subscribe(function (data) { return _this.getEmployeesFromResponse(data.json()); });
    };
    OrgSetService.prototype.EditEmployee = function (employee) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.EditEmployeeLink, "token=" + this.Token +
            "&login=" + employee.login +
            "&name=" + employee.name +
            "&surname=" + employee.surname +
            "&middlename=" + employee.middlename +
            "&sex=" + employee.sex +
            "&team=" + employee.team +
            "&post=" + employee.post, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkEditEmployeeResponse(data.json(), employee); });
    };
    OrgSetService.prototype.DeleteEmployee = function (login) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.DeleteEmployeeLink, "token=" + this.Token + "&login=" + login, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkDeleteEmployeeResponse(data.json(), login); });
    };
    OrgSetService.prototype.ResetEmployeePassword = function (login) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.ResetEmployeePasswordLink, "token=" + this.Token + "&login=" + login, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkResetEmployeePasswordResponse(data.json()); });
    };
    OrgSetService.prototype.AddEmployee = function (employee) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.AddEmployeeLink, "token=" + this.Token +
            "&name=" + employee.name +
            "&surname=" + employee.surname +
            "&middlename=" + employee.middlename +
            "&sex=" + employee.sex +
            "&team=" + employee.team +
            "&post=" + employee.post, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkAddEmployeeResponse(data.json(), employee); });
    };
    OrgSetService.prototype.EditEmployeePermission = function (login, value, id) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.EditEmployeePermissionLink, "token=" + this.Token + "&id=" + id + "&value=" + value + "&login=" + login, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkEditEmployeePermissionResponse(data.json()); });
    };
    OrgSetService.prototype.GetEmployeeFullNameByLogin = function (employee_login) {
        try {
            var data = this.Employees.filter(function (empl) { return empl.login == employee_login; })[0];
            return data.surname + " " + data.name + " " + data.middlename;
        }
        catch (e) {
            return employee_login;
        }
    };
    OrgSetService.prototype.GetEmployeesByTeamID = function (team_id) {
        if (team_id == -1) {
            return this.Employees;
        }
        else {
            return this.Employees.filter(function (part) { return part.team == team_id; });
        }
    };
    OrgSetService.prototype.checkEditEmployeePermissionResponse = function (data) {
        if (data.code == 200) {
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkResetEmployeePasswordResponse = function (data) {
        if (data.code == 200) {
            notie_1.alert({ type: 1, text: "Новый пароль: " + data.password, stay: true });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkDeleteEmployeeResponse = function (data, login) {
        if (data.code == 200) {
            for (var i = 0; i < this.Employees.length; i++) {
                if (this.Employees[i].login == login) {
                    if (this.Employees[i].team != 0) {
                        for (var i_2 = 0; i_2 < this.Teams.length; i_2++) {
                            if (this.Teams[i_2].id == this.Employees[i_2].team) {
                                this.Teams[i_2].leader = { name: '', surname: '', middlename: '', login: '' };
                                break;
                            }
                        }
                    }
                    this.Employees.splice(i, 1);
                    break;
                }
            }
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkEditEmployeeResponse = function (data, employee) {
        if (data.code == 200) {
            for (var i = 0; i < this.Teams.length; i++) {
                if (this.Teams[i].leader.login == employee.login) {
                    this.Teams[i].leader = { name: '', surname: '', middlename: '', login: '' };
                }
                if (this.Teams[i].id == employee.team) {
                    this.Teams[i].leader = { name: employee.name, surname: employee.surname, middlename: employee.middlename, login: employee.login };
                }
            }
            for (var i = 0; i < this.Employees.length; i++) {
                if (this.Employees[i].team == employee.team && this.Employees[i].login != employee.login) {
                    this.Employees[i].team = 0;
                }
            }
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkAddEmployeeResponse = function (data, employee) {
        if (data.code == 200) {
            employee.login = data.login;
            employee.permissions = [];
            for (var i = 0; i < this.Categories.length; i++) {
                employee.permissions.push({ id: this.Categories[i].id, value: true });
            }
            if (employee.team != 0) {
                for (var i = 0; i < this.Teams.length; i++) {
                    if (this.Teams[i].id == employee.team) {
                        this.Teams[i].leader = { name: employee.name, surname: employee.surname, middlename: employee.middlename, login: employee.login };
                        break;
                    }
                }
            }
            this.Employees.push(employee);
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.getEmployeesFromResponse = function (data) {
        if (data.code == 200) {
            this.Employees = data.employees;
            if (this.Employees != null) {
                for (var i = 0; i < this.Employees.length; i++) {
                    if (this.Employees[i].permissions != null) {
                        for (var k = 0; k < this.Employees[i].permissions.length; k++) {
                            this.Employees[i].permissions[k].value = this.StringToBoolean(data.employees[i].permissions[k].value);
                        }
                    }
                }
            }
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
    };
    OrgSetService.prototype.GetReasons = function () {
        var _this = this;
        this.http.get(this.GetReasonsLink + "?token=" + this.Token).subscribe(function (data) { return _this.getReasonsFromResponse(data.json()); });
    };
    OrgSetService.prototype.AddReason = function (catID, text, change) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.AddReasonLink, "token=" + this.Token + "&cat_id=" + catID + "&text=" + text + "&change=" + change, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkAddReasonResponse(data.json(), catID, text, change); });
    };
    OrgSetService.prototype.EditReason = function (reason) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.EditReasonLink, "token=" + this.Token + "&id=" + reason.id + "&text=" + reason.text + "&change=" + reason.change + "&cat_id=" + reason.cat_id, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkEditReasonResponse(data.json()); });
    };
    OrgSetService.prototype.DeleteReason = function (id) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.DeleteReasonLink, "token=" + this.Token + "&id=" + id, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkDeleteReasonResponse(data.json(), id); });
    };
    OrgSetService.prototype.getReasonsFromResponse = function (data) {
        if (data.code == 200) {
            this.Reasons = data.reasons;
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
    };
    OrgSetService.prototype.checkAddReasonResponse = function (data, catID, text, change) {
        if (data.code == 200) {
            this.Reasons.push({ id: data.id, cat_id: catID, text: text, change: change });
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkEditReasonResponse = function (data) {
        if (data.code == 200) {
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.checkDeleteReasonResponse = function (data, id) {
        if (data.code == 200) {
            for (var i = 0; i < this.Reasons.length; i++) {
                if (this.Reasons[i].id == id) {
                    this.Reasons.splice(i, 1);
                    break;
                }
            }
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.GetMarksChanges = function () {
        var _this = this;
        this.http.get(this.GetMarksChangesLink + "?token=" + this.Token).subscribe(function (data) { return _this.getMarksChangesFromResponse(data.json()); });
    };
    OrgSetService.prototype.DeleteMarkChange = function (id) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.DeleteMarkChangeLink, "token=" + this.Token + "&id=" + id, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkDeleteMarkChangeResponse(data.json(), id); });
    };
    OrgSetService.prototype.GetMarksChangesByEmployeeLogin = function (employee_login) {
        return this.MarksChanges.filter(function (markChange) { return markChange.employee_login == employee_login; });
    };
    OrgSetService.prototype.GetMarksChangesByParticipantLogin = function (participant_login) {
        return this.MarksChanges.filter(function (markChange) { return markChange.participant_login == participant_login; });
    };
    OrgSetService.prototype.getMarksChangesFromResponse = function (data) {
        if (data.code == 200) {
            this.MarksChanges = data.marks_changes;
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
    };
    OrgSetService.prototype.checkDeleteMarkChangeResponse = function (data, id) {
        if (data.code == 200) {
            this.MarksChanges = this.MarksChanges.filter(function (markChange) { return markChange.id != id; });
            this.GetData();
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    OrgSetService.prototype.PreloaderOn = function () {
        this.Preloader = true;
    };
    OrgSetService.prototype.PreloaderOff = function () {
        this.Preloader = false;
    };
    OrgSetService.prototype.IntToSex = function (num) {
        if (num == 0) {
            return "мужской";
        }
        else {
            return "женский";
        }
    };
    OrgSetService.prototype.IdToTeamName = function (id) {
        for (var i = 0; i < this.Teams.length; i++) {
            if (this.Teams[i].id == id) {
                return this.Teams[i].name;
            }
        }
        return "отсутствует";
    };
    OrgSetService.prototype.CategoryIdToName = function (id) {
        for (var i = 0; i < this.Categories.length; i++) {
            if (this.Categories[i].id == id) {
                return this.Categories[i].name;
            }
        }
        return "Ошибка!";
    };
    OrgSetService.prototype.StringToBoolean = function (data) {
        if (data == "false") {
            return false;
        }
        else {
            return true;
        }
    };
    OrgSetService.prototype.GetReasonsByCatID = function (id) {
        return this.Reasons.filter(function (reason) { return reason.cat_id == id; });
    };
    OrgSetService.prototype.GetMarkByCategoryID = function (id, marks) {
        return marks.filter(function (mark) { return mark.id == id; });
    };
    OrgSetService.prototype.GetPermissionByCategoryID = function (id, permissions) {
        return permissions.filter(function (permission) { return permission.id == id; });
    };
    return OrgSetService;
}());
OrgSetService = __decorate([
    core_1.Injectable(),
    __param(0, core_1.Inject(http_1.Http)),
    __metadata("design:paramtypes", [http_1.Http])
], OrgSetService);
exports.OrgSetService = OrgSetService;
//# sourceMappingURL=orgset.service.js.map