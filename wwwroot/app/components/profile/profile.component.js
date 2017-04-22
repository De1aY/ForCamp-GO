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
var marks_service_1 = require("../../src/marks.service");
var router_1 = require("@angular/router");
var ProfileComponent = (function () {
    function ProfileComponent(cookieService, checkTokenService, userService, orgSetService, marksService, activeRoute) {
        var _this = this;
        this.cookieService = cookieService;
        this.checkTokenService = checkTokenService;
        this.userService = userService;
        this.orgSetService = orgSetService;
        this.marksService = marksService;
        this.activeRoute = activeRoute;
        this.subscription = activeRoute.params.subscribe(function (params) { return _this.login = params['login']; });
        this.loginOld = this.login;
    }
    ProfileComponent.prototype.ngOnInit = function () {
        this.TokenInit();
        this.ServiceInit();
        this.InitRequestData();
    };
    ProfileComponent.prototype.ngDoCheck = function () {
        if (this.loginOld != this.login) {
            this.InitRequestData();
            this.loginOld = this.login;
        }
    };
    ProfileComponent.prototype.ngOnDestroy = function () {
        this.subscription.unsubscribe();
    };
    ProfileComponent.prototype.ServiceInit = function () {
        this.UserServiceInit();
        this.OrgSetServiceInit();
        this.MarksServiceInit();
    };
    ProfileComponent.prototype.OrgSetServiceInit = function () {
        if (this.orgSetService.Token == undefined) {
            this.orgSetService.Token = this.Token;
        }
        this.orgSetService.GetData();
    };
    ProfileComponent.prototype.MarksServiceInit = function () {
        this.marksService.Token = this.Token;
    };
    ProfileComponent.prototype.UserServiceInit = function () {
        if (this.userService.Token == undefined) {
            this.userService.Token = this.Token;
        }
        this.userService.GetData();
    };
    ProfileComponent.prototype.TokenInit = function () {
        this.Token = this.cookieService.get("token");
        if (this.Token == undefined) {
            window.location.href = "https://forcamp.ga/exit";
        }
        this.checkTokenService.CheckToken(this.Token);
    };
    ProfileComponent.prototype.InitRequestData = function () {
        this.userService.GetUserData(this.login);
    };
    return ProfileComponent;
}());
ProfileComponent = __decorate([
    core_1.Component({
        selector: "profile",
        templateUrl: "app/components/profile/profile.component.html",
        styleUrls: ["app/components/profile/profile.component.css"]
    }),
    __metadata("design:paramtypes", [core_2.CookieService,
        checkToken_service_1.CheckTokenService,
        user_service_1.UserService,
        orgset_service_1.OrgSetService,
        marks_service_1.MarksService,
        router_1.ActivatedRoute])
], ProfileComponent);
exports.ProfileComponent = ProfileComponent;
//# sourceMappingURL=profile.component.js.map