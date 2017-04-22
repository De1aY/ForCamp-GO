import {Component, OnInit, OnDestroy, DoCheck} from '@angular/core';
import {CookieService} from 'angular2-cookie/core';
import {CheckTokenService} from '../../src/checkToken.service';
import {UserService} from '../../src/user.service';
import {OrgSetService} from "../../src/orgset.service";
import {alert} from "notie";
import {MarksService} from "../../src/marks.service";
import {Subscription} from "rxjs";
import {ActivatedRoute} from "@angular/router";

@Component({
    selector: "profile",
    templateUrl: "app/components/profile/profile.component.html",
    styleUrls: ["app/components/profile/profile.component.css"]
})
export class ProfileComponent implements OnInit, OnDestroy, DoCheck {
    private Token: string;
    private login: string;
    private loginOld: string;
    private subscription: Subscription;

    constructor(private cookieService: CookieService,
                private checkTokenService: CheckTokenService,
                public userService: UserService,
                public orgSetService: OrgSetService,
                public marksService: MarksService,
                private activeRoute: ActivatedRoute) {
        this.subscription = activeRoute.params.subscribe(params => this.login = params['login']);
        this.loginOld = this.login;
    }

    ngOnInit() {
        this.TokenInit();
        this.ServiceInit();
        this.InitRequestData();
    }

    ngDoCheck() {
        if (this.loginOld != this.login) {
            this.InitRequestData();
            this.loginOld = this.login;
        }
    }

    ngOnDestroy() {
        this.subscription.unsubscribe();
    }

    private ServiceInit() {
        this.UserServiceInit();
        this.OrgSetServiceInit();
        this.MarksServiceInit();
    }

    private OrgSetServiceInit() {
        if (this.orgSetService.Token == undefined) {
            this.orgSetService.Token = this.Token;
        }
        this.orgSetService.GetData();
    }

    private MarksServiceInit() {
        this.marksService.Token = this.Token;
    }

    private UserServiceInit() {
        if (this.userService.Token == undefined) {
            this.userService.Token = this.Token;
        }
        this.userService.GetData();
    }

    private TokenInit() {
        this.Token = this.cookieService.get("token");
        if (this.Token == undefined) {
            window.location.href = "https://forcamp.ga/exit";
        }
        this.checkTokenService.CheckToken(this.Token);
    }

    private InitRequestData() {
        this.userService.GetUserData(this.login);
    }

}