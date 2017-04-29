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
var ApanelService = (function () {
    function ApanelService(http) {
        this.http = http;
        this.CreateOrganizationLink = "https://api.forcamp.ga/apanel.organization.add";
        this.TokenVerifyLink = "https://api.forcamp.ga/token.verify";
        this.PostHeaders = new http_1.Headers();
        this.PostHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
    }
    ApanelService.prototype.ServiceInit = function () {
        this.checkToken();
    };
    ApanelService.prototype.CreateOrganization = function (orgname) {
        var _this = this;
        this.http.post(this.CreateOrganizationLink, "token=" + this.Token + "&orgname=" + orgname, { headers: this.PostHeaders }).subscribe(function (data) { return _this.checkCreateOrganizationResponse(data.json); });
    };
    ApanelService.prototype.checkToken = function () {
        var _this = this;
        this.http.get(this.TokenVerifyLink + "?token=" + this.Token).subscribe(function (data) { return _this.checkTokenResponse(data.json()); });
    };
    ApanelService.prototype.checkTokenResponse = function (data) {
        if (data.code != 200) {
            window.location.href = "https://forcamp.ga/exit";
        }
        else {
            if (data.admin_status != true) {
                window.location.href = "https://forcamp.ga";
            }
        }
    };
    ApanelService.prototype.checkCreateOrganizationResponse = function (data) {
        if (data.code == 200) {
            alert(data.login + " " + data.password);
        }
        else {
            alert(data.code);
        }
    };
    return ApanelService;
}());
ApanelService = __decorate([
    core_1.Injectable(),
    __param(0, core_1.Inject(http_1.Http)),
    __metadata("design:paramtypes", [http_1.Http])
], ApanelService);
exports.ApanelService = ApanelService;
//# sourceMappingURL=apanel.service.js.map