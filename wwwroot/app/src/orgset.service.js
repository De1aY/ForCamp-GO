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
        this.SetOrgSettingValueLink = "https://api.forcamp.ga/orgset.setting.set";
        this.PostHeaders = new http_1.Headers();
        this.OrgSettings = {
            organization: "загрузка...",
            period: "загрузка...",
            participant: "загрузка...",
            self_marks: false,
            team: "загрузка..."
        };
        this.Preloader = false;
        this.ParticipantValueEdit_Active = false;
        this.PeriodValueEdit_Active = false;
        this.TeamValueEdit_Active = false;
        this.OrganizationValueEdit_Active = false;
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
    PreloaderOn() {
        this.Preloader = true;
    }
    PreloaderOff() {
        this.Preloader = false;
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