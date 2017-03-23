"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
const core_1 = require("@angular/core");
const platform_browser_1 = require("@angular/platform-browser");
const forms_1 = require("@angular/forms");
const http_1 = require("@angular/http");
const material_1 = require("@angular/material");
const cookies_service_1 = require("angular2-cookie/services/cookies.service");
const authorization_component_1 = require("./authorization.component");
let AuthorizationModule = class AuthorizationModule {
};
AuthorizationModule = __decorate([
    core_1.NgModule({
        imports: [platform_browser_1.BrowserModule, forms_1.FormsModule, http_1.HttpModule, material_1.MaterialModule],
        declarations: [authorization_component_1.AuthorizationComponent],
        bootstrap: [authorization_component_1.AuthorizationComponent],
        providers: [cookies_service_1.CookieService],
    })
], AuthorizationModule);
exports.AuthorizationModule = AuthorizationModule;
//# sourceMappingURL=authorization.module.js.map