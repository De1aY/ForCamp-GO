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
var CheckTokenService = (function () {
    function CheckTokenService(http) {
        this.http = http;
        this.Link = "https://api.forcamp.ga/token.verify";
    }
    CheckTokenService.prototype.CheckToken = function (Token) {
        var _this = this;
        this.http.get(this.Link + "?token=" + Token).subscribe(function (data) { return _this.checkToken(data.json()); });
    };
    CheckTokenService.prototype.checkToken = function (data) {
        if (data.code != 200) {
            window.location.href = "https://forcamp.ga/exit";
        }
    };
    return CheckTokenService;
}());
CheckTokenService = __decorate([
    core_1.Injectable(),
    __param(0, core_1.Inject(http_1.Http)),
    __metadata("design:paramtypes", [http_1.Http])
], CheckTokenService);
exports.CheckTokenService = CheckTokenService;
//# sourceMappingURL=checkToken.service.js.map