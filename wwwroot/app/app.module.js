"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
var core_1 = require("@angular/core");
var platform_browser_1 = require("@angular/platform-browser");
var forms_1 = require("@angular/forms");
var http_1 = require("@angular/http");
var material_1 = require("@angular/material");
var router_1 = require("@angular/router");
var ngx_datatable_1 = require("@swimlane/ngx-datatable");
var ngx_charts_1 = require("@swimlane/ngx-charts");
var animations_1 = require("@angular/platform-browser/animations");
require("hammerjs");
var MDLInit_1 = require("./MDLInit");
var landing_component_1 = require("./components/landing/landing.component");
var app_component_1 = require("./app.component");
var orgset_componenet_1 = require("./components/orgset/orgset.componenet");
var participantValueEdit_component_1 = require("./components/orgset/participantValueEdit/participantValueEdit.component");
var periodValueEdit_component_1 = require("./components/orgset/periodValueEdit/periodValueEdit.component");
var teamValueEdit_component_1 = require("./components/orgset/teamValueEdit/teamValueEdit.component");
var organizationValueEdit_component_1 = require("./components/orgset/organizationValueEdit/organizationValueEdit.component");
var addCategory_component_1 = require("./components/orgset/addCategory/addCategory.component");
var addTeam_component_1 = require("./components/orgset/addTeam/addTeam.component");
var addParticipant_component_1 = require("./components/orgset/addParticipant/addParticipant.component");
var addEmployee_component_1 = require("./components/orgset/addEmployee/addEmployee.component");
var marks_component_1 = require("./components/marks/marks.component");
var addReason_component_1 = require("./components/orgset/addReason/addReason.component");
var profile_component_1 = require("./components/profile/profile.component");
var general_component_1 = require("./components/general/general.component");
var apanel_service_1 = require("./src/apanel.service");
var orgset_service_1 = require("./src/orgset.service");
var marks_service_1 = require("./src/marks.service");
var apanel_component_1 = require("./components/apanel/apanel.component");
var checkToken_service_1 = require("./src/checkToken.service");
var user_service_1 = require("./src/user.service");
var cookies_service_1 = require("angular2-cookie/services/cookies.service");
var appRoutes = [
    { path: 'orgset', component: orgset_componenet_1.OrgSetComponent },
    { path: 'marks', component: marks_component_1.MarksComponent },
    { path: '', component: landing_component_1.LandingComponent },
    { path: 'profile/:login', component: profile_component_1.ProfileComponent },
    { path: 'general', component: general_component_1.GeneralComponent },
    { path: 'apanel', component: apanel_component_1.ApanelComponent }
];
var AppModule = (function () {
    function AppModule() {
    }
    return AppModule;
}());
AppModule = __decorate([
    core_1.NgModule({
        imports: [platform_browser_1.BrowserModule, forms_1.FormsModule, http_1.HttpModule, material_1.MaterialModule, router_1.RouterModule.forRoot(appRoutes), ngx_datatable_1.NgxDatatableModule, ngx_charts_1.NgxChartsModule, animations_1.BrowserAnimationsModule],
        declarations: [
            landing_component_1.LandingComponent,
            app_component_1.AppComponent,
            orgset_componenet_1.OrgSetComponent,
            marks_component_1.MarksComponent,
            profile_component_1.ProfileComponent,
            general_component_1.GeneralComponent,
            apanel_component_1.ApanelComponent,
            MDLInit_1.MDL,
            participantValueEdit_component_1.ParticipantValueEditComponent,
            periodValueEdit_component_1.PeriodValueEditComponent,
            teamValueEdit_component_1.TeamValueEditComponent,
            organizationValueEdit_component_1.OrganizationValueEditComponent,
            addCategory_component_1.AddCategoryComponent,
            addTeam_component_1.AddTeamComponent,
            addParticipant_component_1.AddParticipantComponent,
            addEmployee_component_1.AddEmployeeComponent,
            addReason_component_1.AddReasonComponent,
        ],
        bootstrap: [app_component_1.AppComponent],
        providers: [cookies_service_1.CookieService, checkToken_service_1.CheckTokenService, user_service_1.UserService, orgset_service_1.OrgSetService, marks_service_1.MarksService, apanel_service_1.ApanelService],
    })
], AppModule);
exports.AppModule = AppModule;
//# sourceMappingURL=app.module.js.map