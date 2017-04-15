import {Injectable, Inject} from "@angular/core";
import {Http, Response} from "@angular/http";
import {alert} from "notie";

interface Mark{
    id: number
    value: number
}

interface MarkPermission{
    id: number
    value: boolean
}

interface UserData {
    Name: string
    Surname: string
    Middlename: string
    Team: number
    Sex: number
    Access: number
    Avatar: string
    Organization: string
    Marks: Mark[]
    Permissions: MarkPermission[]
    Post: string
}

@Injectable()
export class UserService {
    private GetUserLoginLink:string = "https://api.forcamp.ga/user.login.get";
    private GetUserDataLink: string = "https://api.forcamp.ga/user.data.get";
    public SelfLogin: string;
    public SelfData: UserData = {
        Name: "загрузка...",
        Surname: "загрузка...",
        Middlename: "загрузка...",
        Team: 0,
        Avatar: "загрузка...",
        Sex: 0,
        Access: 0,
        Organization: "загрузка...",
        Marks: [],
        Permissions: [],
        Post: "загрузка..."
    };
    public UserData: UserData = {
        Name: "загрузка...",
        Surname: "загрузка...",
        Middlename: "загрузка...",
        Team: 0,
        Avatar: "загрузка...",
        Sex: 0,
        Access: 0,
        Organization: "загрузка...",
        Marks: [],
        Permissions: [],
        Post: "загрузка..."
    };
    public Token: string;

    constructor(
        @Inject(Http) private http: Http,
    ){}

    public GetSelfUserData(){
        this.http.get(this.GetUserLoginLink+"?token="+this.Token).subscribe((data: Response) => this.getSelfUserLoginFromResponse(data.json()));
    }

    public GetUserData(login: string){
        this.http.get(this.GetUserDataLink+"?token="+this.Token+"&login="+login).subscribe((data: Response) => this.getUserDataFromResponse(data.json()));
    }

    private getUserDataFromResponse(data: any){
        if (data.code == 200) {
            if (data.data.access > 0){
                this.UserData = {
                    Name: data.data.name,
                    Surname: data.data.surname,
                    Middlename: data.data.middlename,
                    Sex: data.data.sex,
                    Access: data.data.access,
                    Avatar: data.data.avatar,
                    Team: data.data.team,
                    Organization: data.data.organization,
                    Marks: [],
                    Permissions: data.data.permissions,
                    Post: data.data.post
                };
            } else {
                this.UserData = {
                    Name: data.data.name,
                    Surname: data.data.surname,
                    Middlename: data.data.middlename,
                    Sex: data.data.sex,
                    Access: data.data.access,
                    Avatar: data.data.avatar,
                    Team: data.data.team,
                    Organization: data.data.organization,
                    Marks: data.data.marks,
                    Permissions: [],
                    Post: ""
                };
            }
            for (let i = 0; i < this.UserData.Permissions.length; i++){
                this.UserData.Permissions[i].value = this.StringToBoolean(this.UserData.Permissions[i].value);
            }
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
    }

    private getSelfUserLoginFromResponse(data: any){
            if(data.code == 200){
                this.SelfLogin = data.login;
                this.getSelfUserData(this.Token, this.SelfLogin);
            } else {
                alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
            }
    }

    private getSelfUserData(token: string, login: string){
        this.http.get(this.GetUserDataLink+"?token="+token+"&login="+login).subscribe((data: Response) => this.getSelfUserDataFromResponse(data.json()));
    }

    private getSelfUserDataFromResponse(data: any) {
        if (data.code == 200) {
            if (data.data.access > 0){
                this.SelfData = {
                    Name: data.data.name,
                    Surname: data.data.surname,
                    Middlename: data.data.middlename,
                    Sex: data.data.sex,
                    Access: data.data.access,
                    Avatar: data.data.avatar,
                    Team: data.data.team,
                    Organization: data.data.organization,
                    Marks: [],
                    Permissions: data.data.permissions,
                    Post: data.data.post
                };
            } else {
                this.SelfData = {
                    Name: data.data.name,
                    Surname: data.data.surname,
                    Middlename: data.data.middlename,
                    Sex: data.data.sex,
                    Access: data.data.access,
                    Avatar: data.data.avatar,
                    Team: data.data.team,
                    Organization: data.data.organization,
                    Marks: data.data.marks,
                    Permissions: [],
                    Post: ""
                };
            }
            for (let i = 0; i < this.SelfData.Permissions.length; i++){
                this.SelfData.Permissions[i].value = this.StringToBoolean(this.SelfData.Permissions[i].value);
            }
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
    }

    private StringToBoolean(data: any): boolean {
        if (data == "false") {
            return false;
        } else {
            return true;
        }
    }

}