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
Object.defineProperty(exports, "__esModule", { value: true });
var core_1 = require("@angular/core");
var core_2 = require("angular2-cookie/core");
var checkToken_service_1 = require("../../src/checkToken.service");
var user_service_1 = require("../../src/user.service");
var orgset_service_1 = require("../../src/orgset.service");
var notie_1 = require("notie");
var marks_service_1 = require("../../src/marks.service");
var MarksComponent = (function () {
    function MarksComponent(cookieService, checkTokenService, userService, orgSetService, marksService) {
        this.cookieService = cookieService;
        this.checkTokenService = checkTokenService;
        this.userService = userService;
        this.orgSetService = orgSetService;
        this.marksService = marksService;
        this.reasonID = 0;
        this.MarkEdit = {};
    }
    MarksComponent.prototype.ngOnInit = function () {
        this.TokenInit();
        this.ServiceInit();
    };
    MarksComponent.prototype.EditMark = function (login, categoryID, index) {
        if (this.CheckPermissions(categoryID)) {
            if (this.CheckSelfTeamMarks(login)) {
                this.MarkEdit[index + "-mark-" + categoryID] = true;
            }
            else {
                notie_1.alert({ type: 3, text: "Вы не можете изменять баллы своей команде", time: 2 });
            }
        }
        else {
            notie_1.alert({ type: 3, text: "Вы не можете редактировать данную категорию", time: 2 });
        }
    };
    MarksComponent.prototype.ServiceInit = function () {
        this.UserServiceInit();
        this.OrgSetServiceInit();
        this.MarksServiceInit();
    };
    MarksComponent.prototype.OrgSetServiceInit = function () {
        if (this.orgSetService.Token == undefined) {
            this.orgSetService.Token = this.Token;
        }
        this.orgSetService.GetOrgSettings();
        this.orgSetService.GetCategories();
        this.orgSetService.GetTeams();
        this.orgSetService.GetParticipants();
        this.orgSetService.GetEmployees();
        this.orgSetService.GetReasons();
    };
    MarksComponent.prototype.MarksServiceInit = function () {
        this.marksService.Token = this.Token;
    };
    MarksComponent.prototype.UserServiceInit = function () {
        if (this.userService.Token == undefined) {
            this.userService.Token = this.Token;
        }
        this.userService.GetSelfUserData();
    };
    MarksComponent.prototype.CheckSelfTeamMarks = function (login) {
        if (!this.orgSetService.OrgSettings.self_marks) {
            if (this.userService.SelfData.Team != 0) {
                for (var i = 0; i < this.orgSetService.Teams.length; i++) {
                    if (this.orgSetService.Teams[i].id == this.userService.SelfData.Team) {
                        for (var k = 0; k < this.orgSetService.Teams[i].participants.length; k++) {
                            if (this.orgSetService.Teams[i].participants[k] == login) {
                                return false;
                            }
                        }
                        return true;
                    }
                }
            }
            else {
                return true;
            }
        }
        else {
            return true;
        }
    };
    MarksComponent.prototype.CheckPermissions = function (categoryID) {
        for (var i = 0; i < this.userService.SelfData.Permissions.length; i++) {
            if (this.userService.SelfData.Permissions[i].id == categoryID) {
                return this.userService.SelfData.Permissions[i].value;
            }
        }
        return false;
    };
    MarksComponent.prototype.TokenInit = function () {
        this.Token = this.cookieService.get("token");
        if (this.Token == undefined) {
            window.location.href = "https://forcamp.ga/exit";
        }
        this.checkTokenService.CheckToken(this.Token);
    };
    return MarksComponent;
}());
MarksComponent = __decorate([
    core_1.Component({
        selector: "marks",
        templateUrl: "app/components/marks/marks.component.html",
        styleUrls: ["app/components/marks/marks.component.css"]
    }),
    __metadata("design:paramtypes", [core_2.CookieService,
        checkToken_service_1.CheckTokenService,
        user_service_1.UserService,
        orgset_service_1.OrgSetService,
        marks_service_1.MarksService])
], MarksComponent);
exports.MarksComponent = MarksComponent;
//# sourceMappingURL=marks.component.js.map