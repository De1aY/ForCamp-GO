import { NgModule }      from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule }   from '@angular/forms';
import { HttpModule } from '@angular/http';
import { MaterialModule } from '@angular/material';

import { CookieService } from 'angular2-cookie/services/cookies.service';

import { AuthorizationComponent } from "./authorization.component";

@NgModule({
    imports:      [ BrowserModule, FormsModule, HttpModule, MaterialModule ],
    declarations: [ AuthorizationComponent ],
    bootstrap:    [ AuthorizationComponent ],
    providers:    [ CookieService ],
})
export class AuthorizationModule { }