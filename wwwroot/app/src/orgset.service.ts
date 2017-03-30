import {Injectable, Inject} from "@angular/core";
import {Http, Response, Headers} from "@angular/http";
import {alert} from "notie";

interface OrgSettings {
    organization: string
    participant: string
    period: string
    self_marks: boolean
    team: string
}

interface Category {
    id: number
    name: string
    negative_marks: boolean
}

interface TeamLeader {
    name: string
    surname: string
    middlename: string
    login: string
}

interface Team{
    id: number
    name: string
    leader: TeamLeader
    count: number
}

@Injectable()
export class OrgSetService {
    //Links
    private GetOrgSettingsLink = "https://api.forcamp.ga/orgset.settings.get";
    private GetCategoriesLink = "https://api.forcamp.ga/orgset.categories.get";
    private SetOrgSettingValueLink = "https://api.forcamp.ga/orgset.setting.set";
    private AddCategoryLink = "https://api.forcamp.ga/orgset.category.add";
    private DeleteCategoryLink = "https://api.forcamp.ga/orgset.category.delete";
    private EditCategoryLink = "https://api.forcamp.ga/orgset.category.edit";
    private GetTeamsLink = "https://api.forcamp.ga/orgset.teams.get";
    private EditTeamLink = "https://api.forcamp.ga/orgset.team.edit";
    private AddTeamLink = "https://api.forcamp.ga/orgset.team.add";
    private DeleteTeamLink = "https://api.forcamp.ga/orgset.team.delete";
    //Var's
    private PostHeaders: Headers = new Headers();
    public Token: string;
    public OrgSettings: OrgSettings = {
        organization: "загрузка...",
        period: "загрузка...",
        participant: "загрузка...",
        self_marks: false,
        team: "загрузка..."
    };
    public Teams: Team[] = [];
    public Categories: Category[] = [];
    public Preloader: boolean = false;
    public ParticipantValueEdit_Active: boolean = false;
    public PeriodValueEdit_Active: boolean = false;
    public TeamValueEdit_Active: boolean = false;
    public OrganizationValueEdit_Active: boolean = false;
    public AddCategory_Active: boolean = false;
    public AddTeam_Active: boolean = false;

    constructor(@Inject(Http) private http: Http,) {
        this.PostHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
    }

    // OrgSettings
    public GetOrgSettings() {
        this.PreloaderOn();
        this.http.get(this.GetOrgSettingsLink + "?token=" + this.Token).subscribe((data: Response) => this.getOrgSettingsFromResponse(data.json()))
    }

    public SetOrgSettingValue(name: string, value: string) {
        this.PreloaderOn();
        this.http.post(this.SetOrgSettingValueLink, "token=" + this.Token + "&name=" + name + "&value=" + value, {headers: this.PostHeaders}).subscribe((data: Response) => this.checkSetOrgSettingValueResponse(data.json(), name, value));
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

    private checkSetOrgSettingValueResponse(data: any, name: string, value: string) {
        if (data.code == 200) {
            this.OrgSettings[name] = value;
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    // Categories
    public GetCategories() {
        this.PreloaderOn();
        this.http.get(this.GetCategoriesLink + "?token=" + this.Token).subscribe((data: Response) => this.getCategoriesFromResponse(data.json()));
    }

    public AddCategory(name: string, negative_marks: boolean) {
        this.PreloaderOn();
        this.http.post(this.AddCategoryLink, "token=" + this.Token + "&name=" + name + "&negative_marks=" + negative_marks, {headers: this.PostHeaders}).subscribe((data: Response) => this.checkAddCategoryResponse(data.json(), name, negative_marks));
    }

    public DeleteCategory(id: number) {
        this.PreloaderOn();
        this.http.post(this.DeleteCategoryLink, "token=" + this.Token + "&id=" + id, {headers: this.PostHeaders}).subscribe((data: Response) => this.checkDeleteCategoryResponse(data.json(), id));
    }

    public EditCategory(category: Category) {
        this.PreloaderOn();
        this.http.post(this.EditCategoryLink, "token=" + this.Token + "&id=" + category.id + "&name=" + category.name + "&negative_marks=" + !category.negative_marks, {headers: this.PostHeaders}).subscribe((data: Response) => this.checkEditCategoryResponse(data.json()));
    }

    private checkAddCategoryResponse(data: any, name: string, negative_marks: boolean) {
        if (data.code == 200) {
            this.Categories.push({id: data.id, name: name, negative_marks: negative_marks});
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private getCategoriesFromResponse(data: any) {
        if (data.code == 200) {
            for (let i = 0; i < data.categories.length; i++) {
                this.Categories.push({
                    id: data.categories[i].id,
                    name: data.categories[i].name,
                    negative_marks: this.StringToBoolean(data.categories[i].negative_marks)
                });
            }
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkDeleteCategoryResponse(data: any, id: number) {
        if (data.code == 200) {
            for (let i = 0; i < this.Categories.length; i++) {
                if (this.Categories[i].id == id) {
                    this.Categories.splice(i, 1);
                    break
                }
            }
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkEditCategoryResponse(data: any) {
        if (data.code == 200) {
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    // Teams
    public GetTeams(){
        this.PreloaderOn();
        this.http.get(this.GetTeamsLink+"?token="+this.Token).subscribe((data: Response) => this.getTeamsFromResponse(data.json()));
    }

    public AddTeam(name: string){
        this.PreloaderOn();
        this.http.post(this.AddTeamLink, "token="+this.Token+"&name="+name, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkAddTeamResponse(data.json(), name));
    }

    public EditTeam(id: number,name: string){
        this.PreloaderOn();
        this.http.post(this.EditTeamLink, "token="+this.Token+"&id="+id+"&name="+name, { headers: this.PostHeaders}).subscribe((data: Response) => this.checkEditTeamResponse(data.json()));
    }

    public DeleteTeam(id: number){
        this.PreloaderOn();
        this.http.post(this.DeleteTeamLink, "token="+this.Token+"&id="+id, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkDeleteTeamResponse(data.json(), id));
    }

    private getTeamsFromResponse(data: any) {
        if(data.code == 200){
            this.Teams = data.teams;
            console.log(this.Teams);
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkAddTeamResponse(data: any, name: string){
        if(data.code == 200){
            this.Teams.push({id: data.id, name: name, count: 0, leader: {name: "", surname: "", middlename: "", login: ""}});
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkEditTeamResponse(data: any){
        if(data.code == 200){
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkDeleteTeamResponse(data: any, id: number){
        if(data.code == 200){
            for (let i = 0; i < this.Teams.length; i++) {
                if (this.Teams[i].id == id) {
                    this.Teams.splice(i, 1);
                    break
                }
            }
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }


    // Preloader
    public PreloaderOn() {
        this.Preloader = true;
    }

    public PreloaderOff() {
        this.Preloader = false;
    }

    // Tools
    private StringToBoolean(data: string): boolean {
        if (data == "false") {
            return false;
        } else {
            return true;
        }
    }

}