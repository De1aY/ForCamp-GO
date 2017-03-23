import {Injectable, Inject} from "@angular/core";
import {Http, Response} from "@angular/http";

@Injectable()
export class UserService {
    private GetUserLoginLink:string = "https://api.forcamp.ga/user.login";
    private GetUserDataLink: string = "https://api.forcamp.ga/user.data";

    constructor(
        @Inject(Http) private http: Http,
    ){}

    public GetUserLogin(token: string){
        return this.http.get(this.GetUserLoginLink+"?token="+token);
    }

    public GetUserData(token: string, login: string){
        return this.http.get(this.GetUserDataLink+"?token="+token+"&login="+login);
    }

}