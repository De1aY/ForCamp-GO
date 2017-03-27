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
const http_1 = require("@angular/http");
const notie_1 = require("notie");
const core_2 = require("angular2-cookie/core");
let AuthorizationComponent = class AuthorizationComponent {
    constructor(http, cookieService) {
        this.http = http;
        this.cookieService = cookieService;
        this.FormActive = false;
    }
    ngOnInit() {
        this.Token = this.cookieService.get("token");
    }
    SubmitSignInForm() {
        this.http.get("https://api.forcamp.ga/token.get?login=" + this.Login + "&password=" + this.Password).subscribe((data) => this.HandleResponse(data.json()));
        this.Login = '';
        this.Password = '';
        this.FormActive = false;
    }
    HandleResponse(data) {
        if (data.code === 200) {
            this.Time = new Date();
            this.Time.setDate(this.Time.getDate() + 365);
            this.cookieService.put("token", data.token, {
                path: "/",
                expires: this.Time.toUTCString(),
                secure: true,
            });
            this.Token = data.token;
            notie_1.alert({ type: 1, text: "Вход успешно выполнен", time: 3 });
        }
        else {
            notie_1.alert({ type: 3, text: "Произошла ошибка " + data.code, time: 3 });
        }
    }
};
AuthorizationComponent = __decorate([
    core_1.Component({
        selector: "sign_in",
        templateUrl: "app/components/authorization/authorization.component.html",
        styleUrls: ["app/components/authorization/authorization.component.css"]
    }),
    __metadata("design:paramtypes", [http_1.Http,
        core_2.CookieService])
], AuthorizationComponent);
exports.AuthorizationComponent = AuthorizationComponent;
//# sourceMappingURL=authorization.component.js.map