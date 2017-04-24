import {Component, OnInit, ViewChild} from '@angular/core';
import {CookieService} from 'angular2-cookie/core';
import {CheckTokenService} from '../../src/checkToken.service';
import {UserService} from '../../src/user.service';
import {OrgSetService} from "../../src/orgset.service";
import {alert} from "notie";

@Component({
    selector: "general",
    templateUrl: "app/components/general/general.component.html",
    styleUrls: ["app/components/general/general.component.css"],
})
export class GeneralComponent implements OnInit {
    private Token: string;
    BarChartColorScheme = {
        domain: ['#5AA454', '#A10A28', '#C7B42C', '#3F51B5', '#FF772B', '#8649BA', '#FF0000']
    };

    constructor(private cookieService: CookieService,
                private checkTokenService: CheckTokenService,
                public userService: UserService,
                public orgSetService: OrgSetService) {
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
        }
        this.orgSetService.GetData();
    }
    private UserServiceInit(){
        if(this.userService.Token == undefined) {
            this.userService.Token = this.Token;
        }
        this.userService.GetData();
    }

    private TokenInit(){
        this.Token = this.cookieService.get("token");
        if (this.Token == undefined) {
            window.location.href = "https://forcamp.ga/exit";
        }
        this.checkTokenService.CheckToken(this.Token);
    }

}