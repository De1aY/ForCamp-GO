import {Injectable, Inject} from "@angular/core";
import {Http, Response, Headers} from "@angular/http";
import {alert} from "notie";

interface UserData {
    Name: string
    Surname: string
    Middlename: string
    Team: string
    Sex: number
    Access: number
    Avatar: string
    Organization: string
}

interface OrgSettings {
    organization: string
    participant: string
    period: string
    self_marks: boolean
    team: string
}

@Injectable()
export class OrgSetService {
    private GetOrgSettingsLink = "https://api.forcamp.ga/orgset.settings.get";
    private SetOrgSettingValueLink = "https://api.forcamp.ga/orgset.setting.set";
    private PostHeaders: Headers = new Headers();
    public Token: string;
    public OrgSettings: OrgSettings = {
        organization: "загрузка...",
        period: "загрузка...",
        participant: "загрузка...",
        self_marks: false,
        team: "загрузка..."
    };
    public Preloader: boolean = false;
    public ParticipantValueEdit_Active: boolean = false;
    public PeriodValueEdit_Active: boolean = false;
    public TeamValueEdit_Active: boolean = false;
    public OrganizationValueEdit_Active: boolean = false;

    constructor(@Inject(Http) private http: Http,) {
        this.PostHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
    }

    public GetOrgSettings() {
        this.PreloaderOn();
        this.http.get(this.GetOrgSettingsLink + "?token=" + this.Token).subscribe((data: Response) => this.getOrgSettingsFromResponse(data.json()))
    }

    public SetOrgSettingValue(name: string, value: string){
        this.PreloaderOn();
        this.http.post(this.SetOrgSettingValueLink, "token="+this.Token+"&name="+name+"&value="+value, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkSetOrgSettingValueResponse(data.json(), name, value));
    }

    public PreloaderOn(){
        this.Preloader = true;
    }

    public PreloaderOff(){
        this.Preloader = false;
    }

    private checkSetOrgSettingValueResponse(data: any, name: string, value: string){
        if(data.code == 200){
            this.OrgSettings[name] = value;
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private getOrgSettingsFromResponse(data: any) {
        if (data.code == 200) {
            this.OrgSettings = {
                    organization: data.settings.organization,
                    period: data.settings.period,
                    participant: data.settings.participant,
                    self_marks: this.StringToBoolean(data.settings.self_marks),
                    team: data.settings.team
            };
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private StringToBoolean(data: string): boolean{
        if(data == "false"){
            return false;
        } else {
            return true;
        }
    }

}