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
var OrgMainComponent = (function () {
    function OrgMainComponent(cookieService, checkTokenService, userService, orgSetService) {
        this.cookieService = cookieService;
        this.checkTokenService = checkTokenService;
        this.userService = userService;
        this.orgSetService = orgSetService;
    }
    OrgMainComponent.prototype.ngOnInit = function () {
        this.TokenInit();
        this.ServiceInit();
    };
    OrgMainComponent.prototype.ServiceInit = function () {
        this.UserServiceInit();
        this.OrgSetServiceInit();
    };
    OrgMainComponent.prototype.OrgSetServiceInit = function () {
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
    OrgMainComponent.prototype.UserServiceInit = function () {
        if (this.userService.Token == undefined) {
            this.userService.Token = this.Token;
        }
        this.userService.GetSelfUserData();
    };
    OrgMainComponent.prototype.TokenInit = function () {
        this.Token = this.cookieService.get("token");
        if (this.Token == undefined) {
            window.location.href = "https://forcamp.ga/exit";
        }
        this.checkTokenService.CheckToken(this.Token);
    };
    return OrgMainComponent;
}());
OrgMainComponent = __decorate([
    core_1.Component({
        selector: "org_main",
        templateUrl: "app/components/orgmain/orgmain.component.html",
        styleUrls: ["app/components/orgmain/orgmain.component.css"]
    }),
    __metadata("design:paramtypes", [core_2.CookieService,
        checkToken_service_1.CheckTokenService,
        user_service_1.UserService,
        orgset_service_1.OrgSetService])
], OrgMainComponent);
exports.OrgMainComponent = OrgMainComponent;
//# sourceMappingURL=orgmain.component.js.map