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
var cookies_service_1 = require("angular2-cookie/services/cookies.service");
var apanel_service_1 = require("../../src/apanel.service");
var ApanelComponent = (function () {
    function ApanelComponent(cookieService, ApanelService) {
        this.cookieService = cookieService;
        this.ApanelService = ApanelService;
        this.OrgName = "";
    }
    ApanelComponent.prototype.ngOnInit = function () {
        this.TokenInit();
        this.ServiceInit();
    };
    ApanelComponent.prototype.ServiceInit = function () {
        this.ApanelService.Token = this.Token;
        this.ApanelService.ServiceInit();
    };
    ApanelComponent.prototype.TokenInit = function () {
        this.Token = this.cookieService.get("token");
        if (this.Token == undefined) {
            window.location.href = "https://forcamp.ga/exit";
        }
    };
    ApanelComponent.prototype.CreateOrganization = function () {
        this.ApanelService.CreateOrganization(this.OrgName);
    };
    return ApanelComponent;
}());
ApanelComponent = __decorate([
    core_1.Component({
        selector: "apanel",
        templateUrl: "app/components/apanel/apanel.component.html",
        styleUrls: ["app/components/apanel/apanel.component.css"]
    }),
    __metadata("design:paramtypes", [cookies_service_1.CookieService,
        apanel_service_1.ApanelService])
], ApanelComponent);
exports.ApanelComponent = ApanelComponent;
//# sourceMappingURL=apanel.component.js.map