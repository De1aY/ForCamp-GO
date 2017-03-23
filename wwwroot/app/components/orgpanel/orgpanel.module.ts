import { NgModule }      from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule }   from '@angular/forms';
import { HttpModule } from '@angular/http';
import { MaterialModule } from '@angular/material';

import { CheckTokenService } from '../../src/checkToken.service';
import {UserService} from '../../src/user.service';
import { CookieService } from 'angular2-cookie/services/cookies.service';

import { OrgPanelComponent } from "./orgpanel.component";

@NgModule({
    imports:      [ BrowserModule, FormsModule, HttpModule, MaterialModule ],
    declarations: [ OrgPanelComponent ],
    bootstrap:    [ OrgPanelComponent ],
    providers:    [ CookieService, CheckTokenService, UserService ],
})
export class OrgPanelModule { }