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
var UserService = (function () {
    function UserService(http) {
        this.http = http;
        this.GetUserLoginLink = "https://api.forcamp.ga/user.login.get";
        this.GetUserDataLink = "https://api.forcamp.ga/user.data.get";
        this.SelfData = {
            Name: "загрузка...",
            Surname: "загрузка...",
            Middlename: "загрузка...",
            Team: "загрузка...",
            Avatar: "загрузка...",
            Sex: 0,
            Access: 0,
            Organization: "загрузка..."
        };
        this.UserData = {
            Name: "загрузка...",
            Surname: "загрузка...",
            Middlename: "загрузка...",
            Team: "загрузка...",
            Avatar: "загрузка...",
            Sex: 0,
            Access: 0,
            Organization: "загрузка..."
        };
    }
    UserService.prototype.GetSelfUserData = function () {
        var _this = this;
        this.http.get(this.GetUserLoginLink + "?token=" + this.Token).subscribe(function (data) { return _this.getSelfUserLoginFromResponse(data.json()); });
    };
    UserService.prototype.GetUserData = function (login) {
        var _this = this;
        this.http.get(this.GetUserDataLink + "?token=" + this.Token + "&login=" + login).subscribe(function (data) { return _this.getUserDataFromResponse(data.json()); });
    };
    UserService.prototype.getUserDataFromResponse = function (data) {
        if (data.code == 200) {
            this.UserData = {
                Name: data.data.name,
                Surname: data.data.surname,
                Middlename: data.data.middlename,
                Sex: data.data.sex,
                Access: data.data.access,
                Avatar: data.data.avatar,
                Team: data.data.team,
                Organization: data.data.organization
            };
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
    };
    UserService.prototype.getSelfUserLoginFromResponse = function (data) {
        if (data.code == 200) {
            this.SelfLogin = data.login;
            this.getSelfUserData(this.Token, this.SelfLogin);
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
    };
    UserService.prototype.getSelfUserData = function (token, login) {
        var _this = this;
        this.http.get(this.GetUserDataLink + "?token=" + token + "&login=" + login).subscribe(function (data) { return _this.getSelfUserDataFromResponse(data.json()); });
    };
    UserService.prototype.getSelfUserDataFromResponse = function (data) {
        if (data.code == 200) {
            this.SelfData = {
                Name: data.data.name,
                Surname: data.data.surname,
                Middlename: data.data.middlename,
                Sex: data.data.sex,
                Access: data.data.access,
                Avatar: data.data.avatar,
                Team: data.data.team,
                Organization: data.data.organization
            };
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
    };
    return UserService;
}());
UserService = __decorate([
    core_1.Injectable(),
    __param(0, core_1.Inject(http_1.Http)),
    __metadata("design:paramtypes", [http_1.Http])
], UserService);
exports.UserService = UserService;
//# sourceMappingURL=user.service.js.map