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
const core_1 = require("@angular/core");
const notie_1 = require("notie");
const core_2 = require("angular2-cookie/core");
const checkToken_service_1 = require("../../src/checkToken.service");
const user_service_1 = require("../../src/user.service");
let OrgPanelComponent = class OrgPanelComponent {
    constructor(cookieService, checkTokenService, userService) {
        this.cookieService = cookieService;
        this.checkTokenService = checkTokenService;
        this.userService = userService;
    }
    ngOnInit() {
        this.Token = this.cookieService.get("token");
        if (this.Token == undefined) {
            window.location.href = "https://forcamp.ga/exit";
        }
        this.checkTokenService.CheckToken(this.Token);
        this.userService.GetUserLogin(this.Token).subscribe((data) => this.getUserLoginFromResponse(data.json()));
    }
    getUserLoginFromResponse(data) {
        if (data.code == 200) {
            this.SelfLogin = data.login;
            this.getUserData();
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
    }
    getUserDataFromResponse(data) {
        if (data.code == 200) {
            this.SelfData = { Name: data.data.name, Surname: data.data.surname,
                Middlename: data.data.middlename, Sex: data.data.sex,
                Access: data.data.access, Avatar: data.data.avatar, Team: data.data.team };
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3 });
        }
    }
    getUserData() {
        this.userService.GetUserData(this.Token, this.SelfLogin).subscribe((data) => this.getUserDataFromResponse(data.json()));
    }
};
OrgPanelComponent = __decorate([
    core_1.Component({
        selector: "org_panel",
        templateUrl: "app/components/orgpanel/orgpanel.component.html",
        styleUrls: ["app/components/orgpanel/orgpanel.component.css"]
    }),
    __metadata("design:paramtypes", [core_2.CookieService,
        checkToken_service_1.CheckTokenService,
        user_service_1.UserService])
], OrgPanelComponent);
exports.OrgPanelComponent = OrgPanelComponent;
//# sourceMappingURL=orgpanel.component.js.map