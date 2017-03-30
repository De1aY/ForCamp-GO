import {Component, OnInit} from '@angular/core';
import {Http, Response} from '@angular/http';
import {alert} from "notie";
import {CookieService} from 'angular2-cookie/core';
import {CheckTokenService} from '../../src/checkToken.service';
import {UserService} from '../../src/user.service';
import {OrgSetService} from "../../src/orgset.service";


@Component({
    selector: "org_main",
    templateUrl: "app/components/orgmain/orgmain.component.html",
    styleUrls: ["app/components/orgmain/orgmain.component.css"]
})
export class OrgMainComponent implements OnInit {
    private Token: string;

    constructor(private cookieService: CookieService,
                private checkTokenService: CheckTokenService,
                private userService: UserService,
                public orgSetService: OrgSetService,) {
    }

    ngOnInit() {
        this.TokenInit();
        this.ServiceInit();
    }

    private ServiceInit(){
        this.UserServiceInit();
        this.OrgSetServiceInit();
    }

    private OrgSetServiceInit(){
        if(this.orgSetService.Token == undefined){
            this.orgSetService.Token = this.Token;
            this.orgSetService.GetOrgSettings();
            this.orgSetService.GetCategories();
            this.orgSetService.GetTeams();
        }
    }

    private UserServiceInit(){
        if(this.userService.Token == undefined) {
            this.userService.Token = this.Token;
            this.userService.GetSelfUserData();
        }
    }

    private TokenInit(){
        this.Token = this.cookieService.get("token");
        if (this.Token == undefined) {
            window.location.href = "https://forcamp.ga/exit";
        }
        this.checkTokenService.CheckToken(this.Token);
    }

}