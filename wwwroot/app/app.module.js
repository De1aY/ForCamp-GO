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
const router_1 = require("@angular/router");
require("hammerjs");
const checkToken_service_1 = require("./src/checkToken.service");
const user_service_1 = require("./src/user.service");
const cookies_service_1 = require("angular2-cookie/services/cookies.service");
const MDLInit_1 = require("./MDLInit");
const orgmain_component_1 = require("./components/orgmain/orgmain.component");
const authorization_component_1 = require("./components/authorization/authorization.component");
const app_component_1 = require("./app.component");
const orgset_componenet_1 = require("./components/orgset/orgset.componenet");
const orgset_service_1 = require("./src/orgset.service");
const participantValueEdit_component_1 = require("./components/orgset/participantValueEdit/participantValueEdit.component");
const periodValueEdit_component_1 = require("./components/orgset/periodValueEdit/periodValueEdit.component");
const teamValueEdit_component_1 = require("./components/orgset/teamValueEdit/teamValueEdit.component");
const organizationValueEdit_component_1 = require("./components/orgset/organizationValueEdit/organizationValueEdit.component");
const appRoutes = [
    { path: 'orgset', component: orgset_componenet_1.OrgSetComponent },
    { path: 'main', component: orgmain_component_1.OrgMainComponent },
    { path: '', component: authorization_component_1.AuthorizationComponent },
];
let AppModule = class AppModule {
};
AppModule = __decorate([
    core_1.NgModule({
        imports: [platform_browser_1.BrowserModule, forms_1.FormsModule, http_1.HttpModule, material_1.MaterialModule, router_1.RouterModule.forRoot(appRoutes)],
        declarations: [orgmain_component_1.OrgMainComponent,
            authorization_component_1.AuthorizationComponent,
            app_component_1.AppComponent,
            orgset_componenet_1.OrgSetComponent,
            MDLInit_1.MDL,
            participantValueEdit_component_1.ParticipantValueEditComponent,
            periodValueEdit_component_1.PeriodValueEditComponent,
            teamValueEdit_component_1.TeamValueEditComponent,
            organizationValueEdit_component_1.OrganizationValueEditComponent],
        bootstrap: [app_component_1.AppComponent],
        providers: [cookies_service_1.CookieService, checkToken_service_1.CheckTokenService, user_service_1.UserService, orgset_service_1.OrgSetService],
    })
], AppModule);
exports.AppModule = AppModule;
//# sourceMappingURL=app.module.js.map