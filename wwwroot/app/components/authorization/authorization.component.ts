import {Component, OnInit} from '@angular/core';
import {Http, Response} from '@angular/http';
import {alert} from "notie";
import {CookieService} from 'angular2-cookie/core';

@Component({
    selector: "sign_in",
    templateUrl: "app/components/authorization/authorization.component.html",
    styleUrls: ["app/components/authorization/authorization.component.css"]
})
export class AuthorizationComponent implements OnInit{
    Login: string;
    Password: string;
    Token: string;
    Time: Date;
    FormActive: boolean = false;

    constructor(private http: Http,
                private cookieService: CookieService,) {
    }

    ngOnInit(){
        this.Token = this.cookieService.get("token");
    }

    SubmitSignInForm() {
        this.http.get("https://api.forcamp.ga/token.get?login=" + this.Login + "&password=" + this.Password).subscribe((data: Response) => this.HandleResponse(data.json()));
        this.FormActive = false;
    }

    HandleResponse(data: any) {
        if (data.code === 200) {
            this.Time = new Date();
            this.Time.setDate(this.Time.getDate() + 365);
            this.cookieService.put("token", data.token, {
                path: "/",
                expires: this.Time.toUTCString(),
                secure: true,
            });
            this.Token = data.token;
            alert({type: 1, text: "Вход успешно выполнен", time: 3});
        } else {
            alert({type: 3, text: "Произошла ошибка " + data.code, time: 3});
        }
    }

    GoToMainSite(){
        window.location.href = "https://forcamp.ga/main.html";
    }
}