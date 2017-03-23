import {Component, OnInit} from '@angular/core';
import {Http, Response} from '@angular/http';
import {alert} from "notie";
import {CookieService} from 'angular2-cookie/core';
import {CheckTokenService} from '../../src/checkToken.service';
import {UserService} from '../../src/user.service';

interface UserData{
    Name: string
    Surname: string
    Middlename: string
    Team: string
    Sex: number
    Access: number
    Avatar: string
}

@Component({
    selector: "org_panel",
    templateUrl: "app/components/orgpanel/orgpanel.component.html",
    styleUrls: ["app/components/orgpanel/orgpanel.component.css"]
})
export class OrgPanelComponent implements OnInit {
    Token: string;
    SelfLogin: string;
    SelfData: UserData;

    constructor(private cookieService: CookieService,
                private checkTokenService: CheckTokenService,
                private userService: UserService,) {
    }

    ngOnInit() {
        this.Token = this.cookieService.get("token");
        if (this.Token == undefined) {
            window.location.href = "https://forcamp.ga/exit";
        }
        this.checkTokenService.CheckToken(this.Token);
        this.userService.GetUserLogin(this.Token).subscribe((data: Response) => this.getUserLoginFromResponse(data.json()));

    }

    private getUserLoginFromResponse(data: any){
        if(data.code == 200){
            this.SelfLogin = data.login;
            this.getUserData();
        } else {
            alert({type: 3, text: "Произошла ошибка("+data.code+")!", time: 3});
        }
    }

    private getUserDataFromResponse(data: any){
        if(data.code == 200){
            this.SelfData = {Name: data.data.name, Surname: data.data.surname,
                Middlename: data.data.middlename, Sex: data.data.sex,
                Access: data.data.access, Avatar: data.data.avatar, Team: data.data.team};
        } else {
            alert({type: 3, text: "Произошла ошибка("+data.code+")!", time: 3});
        }
    }

    private getUserData(){
        this.userService.GetUserData(this.Token, this.SelfLogin).subscribe((data: Response) => this.getUserDataFromResponse(data.json()))
    }
}