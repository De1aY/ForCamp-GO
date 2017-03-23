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
let UserService = class UserService {
    constructor(http) {
        this.http = http;
        this.GetUserLoginLink = "https://api.forcamp.ga/user.login";
        this.GetUserDataLink = "https://api.forcamp.ga/user.data";
    }
    GetUserLogin(token) {
        return this.http.get(this.GetUserLoginLink + "?token=" + token);
    }
    GetUserData(token, login) {
        return this.http.get(this.GetUserDataLink + "?token=" + token + "&login=" + login);
    }
};
UserService = __decorate([
    core_1.Injectable(),
    __param(0, core_1.Inject(http_1.Http)),
    __metadata("design:paramtypes", [http_1.Http])
], UserService);
exports.UserService = UserService;
//# sourceMappingURL=user.service.js.map