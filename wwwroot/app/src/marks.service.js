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
var orgset_service_1 = require("./orgset.service");
var user_service_1 = require("./user.service");
var MarksService = (function () {
    function MarksService(http, orgSetService, userService) {
        this.http = http;
        this.orgSetService = orgSetService;
        this.userService = userService;
        this.EditUserMarkLink = "https://api.forcamp.ga/mark.edit";
        this.PostHeaders = new http_1.Headers();
        this.Preloader = false;
        this.PostHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
    }
    MarksService.prototype.EditParticipantMark = function (participant, id, reason) {
        var _this = this;
        this.PreloaderOn();
        this.http.post(this.EditUserMarkLink, "token=" + this.Token + "&login=" + participant.login + "&category_id=" + id + "&reason_id=" + reason, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkEditParticipantMark(data.json(), participant, id, reason); });
    };
    MarksService.prototype.checkEditParticipantMark = function (data, participant, id, reason_id) {
        if (data.code == 200) {
            var reason = this.orgSetService.Reasons.filter(function (row) { return row.id == reason_id; })[0];
            participant.marks.filter(function (row) { return row.id == id; })[0].value += reason.change;
            this.Time = new Date();
            this.Time.setDate(this.Time.getDate());
            this.orgSetService.MarksChanges.push({ id: id, employee_login: this.userService.SelfLogin, participant_login: participant.login, text: reason.text, time: this.Time.toUTCString(), change: reason.change });
            this.orgSetService.GetData();
            notie_1.alert({ type: 1, text: "Балл успешно изменён!", time: 2 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
        this.PreloaderOff();
    };
    MarksService.prototype.PreloaderOn = function () {
        this.Preloader = true;
    };
    MarksService.prototype.PreloaderOff = function () {
        this.Preloader = false;
    };
    return MarksService;
}());
MarksService = __decorate([
    core_1.Injectable(),
    __param(0, core_1.Inject(http_1.Http)),
    __metadata("design:paramtypes", [http_1.Http,
        orgset_service_1.OrgSetService,
        user_service_1.UserService])
], MarksService);
exports.MarksService = MarksService;
//# sourceMappingURL=marks.service.js.map