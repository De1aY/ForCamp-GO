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
const core_1 = require("@angular/core");
const http_1 = require("@angular/http");
const notie_1 = require("notie");
let OrgSetService = class OrgSetService {
    constructor(http) {
        this.http = http;
        this.GetOrgSettingsLink = "https://api.forcamp.ga/orgset.settings.get";
        this.GetCategoriesLink = "https://api.forcamp.ga/orgset.categories.get";
        this.SetOrgSettingValueLink = "https://api.forcamp.ga/orgset.setting.set";
        this.AddCategoryLink = "https://api.forcamp.ga/orgset.category.add";
        this.DeleteCategoryLink = "https://api.forcamp.ga/orgset.category.delete";
        this.EditCategoryLink = "https://api.forcamp.ga/orgset.category.edit";
        this.GetTeamsLink = "https://api.forcamp.ga/orgset.teams.get";
        this.EditTeamLink = "https://api.forcamp.ga/orgset.team.edit";
        this.AddTeamLink = "https://api.forcamp.ga/orgset.team.add";
        this.DeleteTeamLink = "https://api.forcamp.ga/orgset.team.delete";
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
        this.Preloader = false;
        this.ParticipantValueEdit_Active = false;
        this.PeriodValueEdit_Active = false;
        this.TeamValueEdit_Active = false;
        this.OrganizationValueEdit_Active = false;
        this.AddCategory_Active = false;
        this.AddTeam_Active = false;
        this.PostHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
    }
    GetOrgSettings() {
        this.PreloaderOn();
        this.http.get(this.GetOrgSettingsLink + "?token=" + this.Token).subscribe((data) => this.getOrgSettingsFromResponse(data.json()));
    }
    SetOrgSettingValue(name, value) {
        this.PreloaderOn();
        this.http.post(this.SetOrgSettingValueLink, "token=" + this.Token + "&name=" + name + "&value=" + value, { headers: this.PostHeaders }).subscribe((data) => this.checkSetOrgSettingValueResponse(data.json(), name, value));
    }
    getOrgSettingsFromResponse(data) {
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
    }
    checkSetOrgSettingValueResponse(data, name, value) {
        if (data.code == 200) {
            this.OrgSettings[name] = value;
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    }
    GetCategories() {
        this.PreloaderOn();
        this.http.get(this.GetCategoriesLink + "?token=" + this.Token).subscribe((data) => this.getCategoriesFromResponse(data.json()));
    }
    AddCategory(name, negative_marks) {
        this.PreloaderOn();
        this.http.post(this.AddCategoryLink, "token=" + this.Token + "&name=" + name + "&negative_marks=" + negative_marks, { headers: this.PostHeaders }).subscribe((data) => this.checkAddCategoryResponse(data.json(), name, negative_marks));
    }
    DeleteCategory(id) {
        this.PreloaderOn();
        this.http.post(this.DeleteCategoryLink, "token=" + this.Token + "&id=" + id, { headers: this.PostHeaders }).subscribe((data) => this.checkDeleteCategoryResponse(data.json(), id));
    }
    EditCategory(category) {
        this.PreloaderOn();
        this.http.post(this.EditCategoryLink, "token=" + this.Token + "&id=" + category.id + "&name=" + category.name + "&negative_marks=" + !category.negative_marks, { headers: this.PostHeaders }).subscribe((data) => this.checkEditCategoryResponse(data.json()));
    }
    checkAddCategoryResponse(data, name, negative_marks) {
        if (data.code == 200) {
            this.Categories.push({ id: data.id, name: name, negative_marks: negative_marks });
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    }
    getCategoriesFromResponse(data) {
        if (data.code == 200) {
            for (let i = 0; i < data.categories.length; i++) {
                this.Categories.push({
                    id: data.categories[i].id,
                    name: data.categories[i].name,
                    negative_marks: this.StringToBoolean(data.categories[i].negative_marks)
                });
            }
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    }
    checkDeleteCategoryResponse(data, id) {
        if (data.code == 200) {
            for (let i = 0; i < this.Categories.length; i++) {
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
    }
    checkEditCategoryResponse(data) {
        if (data.code == 200) {
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    }
    GetTeams() {
        this.PreloaderOn();
        this.http.get(this.GetTeamsLink + "?token=" + this.Token).subscribe((data) => this.getTeamsFromResponse(data.json()));
    }
    AddTeam(name) {
        this.PreloaderOn();
        this.http.post(this.AddTeamLink, "token=" + this.Token + "&name=" + name, { headers: this.PostHeaders }).subscribe((data) => this.checkAddTeamResponse(data.json(), name));
    }
    EditTeam(id, name) {
        this.PreloaderOn();
        this.http.post(this.EditTeamLink, "token=" + this.Token + "&id=" + id + "&name=" + name, { headers: this.PostHeaders }).subscribe((data) => this.checkEditTeamResponse(data.json()));
    }
    DeleteTeam(id) {
        this.PreloaderOn();
        this.http.post(this.DeleteTeamLink, "token=" + this.Token + "&id=" + id, { headers: this.PostHeaders }).subscribe((data) => this.checkDeleteTeamResponse(data.json(), id));
    }
    getTeamsFromResponse(data) {
        if (data.code == 200) {
            this.Teams = data.teams;
            console.log(this.Teams);
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    }
    checkAddTeamResponse(data, name) {
        if (data.code == 200) {
            this.Teams.push({ id: data.id, name: name, count: 0, leader: { name: "", surname: "", middlename: "", login: "" } });
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    }
    checkEditTeamResponse(data) {
        if (data.code == 200) {
            notie_1.alert({ type: 1, text: "Операция успешно завершена!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    }
    checkDeleteTeamResponse(data, id) {
        if (data.code == 200) {
            for (let i = 0; i < this.Teams.length; i++) {
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
    }
    PreloaderOn() {
        this.Preloader = true;
    }
    PreloaderOff() {
        this.Preloader = false;
    }
    StringToBoolean(data) {
        if (data == "false") {
            return false;
        }
        else {
            return true;
        }
    }
};
OrgSetService = __decorate([
    core_1.Injectable(),
    __param(0, core_1.Inject(http_1.Http)),
    __metadata("design:paramtypes", [http_1.Http])
], OrgSetService);
exports.OrgSetService = OrgSetService;
//# sourceMappingURL=orgset.service.js.map