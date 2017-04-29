import {NgModule}      from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import {FormsModule}   from '@angular/forms';
import {HttpModule} from '@angular/http';
import {MaterialModule} from '@angular/material';
import {RouterModule, Routes} from "@angular/router";
import {NgxDatatableModule} from "@swimlane/ngx-datatable";
import {NgxChartsModule} from "@swimlane/ngx-charts";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import 'hammerjs'

import {MDL} from "./MDLInit";

// Components:
import {LandingComponent} from "./components/landing/landing.component";
import {AppComponent} from "./app.component";
import {OrgSetComponent} from "./components/orgset/orgset.componenet";
import {ParticipantValueEditComponent} from "./components/orgset/participantValueEdit/participantValueEdit.component";
import {PeriodValueEditComponent} from "./components/orgset/periodValueEdit/periodValueEdit.component";
import {TeamValueEditComponent} from "./components/orgset/teamValueEdit/teamValueEdit.component";
import {OrganizationValueEditComponent} from "./components/orgset/organizationValueEdit/organizationValueEdit.component";
import {AddCategoryComponent} from "./components/orgset/addCategory/addCategory.component";
import {AddTeamComponent} from "./components/orgset/addTeam/addTeam.component";
import {AddParticipantComponent} from "./components/orgset/addParticipant/addParticipant.component";
import {AddEmployeeComponent} from "./components/orgset/addEmployee/addEmployee.component";
import {MarksComponent} from "./components/marks/marks.component";
import {AddReasonComponent} from "./components/orgset/addReason/addReason.component";
import {ProfileComponent} from "./components/profile/profile.component";
import {GeneralComponent} from "./components/general/general.component";
// Services:
import {ApanelService} from "./src/apanel.service";
import {OrgSetService} from "./src/orgset.service";
import {MarksService} from "./src/marks.service";
import {ApanelComponent} from "./components/apanel/apanel.component";
import {CheckTokenService} from './src/checkToken.service';
import {UserService} from './src/user.service';
import {CookieService} from 'angular2-cookie/services/cookies.service';

const appRoutes: Routes = [
    {path: 'orgset', component: OrgSetComponent},
    {path: 'marks', component: MarksComponent},
    {path: '', component: LandingComponent},
    {path: 'profile/:login', component: ProfileComponent},
    {path: 'general', component: GeneralComponent},
    {path: 'apanel', component: ApanelComponent}
];
@NgModule({
    imports: [BrowserModule, FormsModule, HttpModule, MaterialModule, RouterModule.forRoot(appRoutes), NgxDatatableModule, NgxChartsModule, BrowserAnimationsModule ],
    declarations: [
        LandingComponent,
        AppComponent,
        OrgSetComponent,
        MarksComponent,
        ProfileComponent,
        GeneralComponent,
        ApanelComponent,
        MDL,
        ParticipantValueEditComponent,
        PeriodValueEditComponent,
        TeamValueEditComponent,
        OrganizationValueEditComponent,
        AddCategoryComponent,
        AddTeamComponent,
        AddParticipantComponent,
        AddEmployeeComponent,
        AddReasonComponent,],
    bootstrap: [AppComponent],
    providers: [CookieService, CheckTokenService, UserService, OrgSetService, MarksService, ApanelService],
})
export class AppModule {
}