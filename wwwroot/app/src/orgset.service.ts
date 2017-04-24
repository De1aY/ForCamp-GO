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
    participants: string[]
}

interface Mark{
    id: number
    value: number
}

interface MarkPermission{
    id: number
    value: boolean
}

interface Participant{
    login: string
    name: string
    surname: string
    middlename: string
    sex: number
    team: number
    marks: Mark[]
    sum: number
}

interface Employee{
    login: string
    name: string
    surname: string
    middlename: string
    sex: number
    team: number
    post: string
    permissions: MarkPermission[]
}

interface Reason{
    id: number
    cat_id: number
    text: string
    change: number
}

interface MarksChange{
    id: number
    employee_login: string
    participant_login: string
    text: string
    change: number
    time: string
}

@Injectable()
export class OrgSetService{
    //Links: OrgSettings
    private GetOrgSettingsLink = "https://api.forcamp.ga/orgset.settings.get";
    private GetCategoriesLink = "https://api.forcamp.ga/orgset.categories.get";
    private SetOrgSettingValueLink = "https://api.forcamp.ga/orgset.setting.edit";
    //Links: Categories
    private AddCategoryLink = "https://api.forcamp.ga/orgset.category.add";
    private DeleteCategoryLink = "https://api.forcamp.ga/orgset.category.delete";
    private EditCategoryLink = "https://api.forcamp.ga/orgset.category.edit";
    //Links: Teams
    private GetTeamsLink = "https://api.forcamp.ga/orgset.teams.get";
    private EditTeamLink = "https://api.forcamp.ga/orgset.team.edit";
    private AddTeamLink = "https://api.forcamp.ga/orgset.team.add";
    private DeleteTeamLink = "https://api.forcamp.ga/orgset.team.delete";
    //Links: Participants
    private GetParticipantsLink = "https://api.forcamp.ga/orgset.participants.get";
    private EditParticipantLink = "https://api.forcamp.ga/orgset.participant.edit";
    private DeleteParticipantLink = "https://api.forcamp.ga/orgset.participant.delete";
    private ResetParticipantPasswordLink = "https://api.forcamp.ga/orgset.participant.password.reset";
    private AddParticipantLink = "https://api.forcamp.ga/orgset.participant.add";
    //Links: Employees
    private GetEmployeesLink = "https://api.forcamp.ga/orgset.employees.get";
    private EditEmployeeLink = "https://api.forcamp.ga/orgset.employee.edit";
    private DeleteEmployeeLink = "https://api.forcamp.ga/orgset.employee.delete";
    private ResetEmployeePasswordLink = "https://api.forcamp.ga/orgset.employee.password.reset";
    private AddEmployeeLink = "https://api.forcamp.ga/orgset.employee.add";
    private EditEmployeePermissionLink = "https://api.forcamp.ga/orgset.employee.permission.edit";
    //Links: Reasons
    private GetReasonsLink = "https://api.forcamp.ga/orgset.reasons.get";
    private AddReasonLink = "https://api.forcamp.ga/orgset.reason.add";
    private EditReasonLink = "https://api.forcamp.ga/orgset.reason.edit";
    private DeleteReasonLink = "https://api.forcamp.ga/orgset.reason.delete";
    //Links: Marks
    private GetMarksChangesLink = "https://api.forcamp.ga/marks.changes.get";
    private DeleteMarkChangeLink = "https://api.forcamp.ga/mark.change.delete";
    //Var's
    private PostHeaders: Headers = new Headers();
    public Token: string;
    // Data
    public OrgSettings: OrgSettings = {
        organization: "загрузка...",
        period: "загрузка...",
        participant: "загрузка...",
        self_marks: false,
        team: "загрузка..."
    };
    public Teams: Team[] = [];
    public Categories: Category[] = [];
    public Participants: Participant[] = [];
    public Employees: Employee[] = [];
    public Reasons: Reason[] = [];
    public MarksChanges: MarksChange[] = [];
    public Preloader: boolean = false;
    // Pop-up
    public ParticipantValueEdit_Active: boolean = false;
    public PeriodValueEdit_Active: boolean = false;
    public TeamValueEdit_Active: boolean = false;
    public OrganizationValueEdit_Active: boolean = false;
    public AddCategory_Active: boolean = false;
    public AddTeam_Active: boolean = false;
    public AddParticipant_Active: boolean = false;
    public AddEmployee_Active: boolean = false;
    public AddReason_Active: boolean = false;
    // Data for charts
    public ParticipantsVerticalBar: object = [];

    constructor(@Inject(Http) private http: Http,) {
        this.PostHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
    }

    public GetData(){
        this.GetOrgSettings();
        this.GetCategories();
        this.GetTeams();
        this.GetParticipants();
        this.GetEmployees();
        this.GetReasons();
        this.GetMarksChanges();
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
            if(this.Participants != undefined) {
                for (let i = 0; i < this.Participants.length; i++) {
                    this.Participants[i].marks.push({id: data.id, value: 0})
                }
            }
            if(this.Employees != undefined) {
                for (let i = 0; i < this.Employees.length; i++) {
                    this.Employees[i].permissions.push({id: data.id, value: true});
                }
            }
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private getCategoriesFromResponse(data: any) {
        if (data.code == 200) {
            this.Categories = data.categories;
            if(this.Categories != null) {
                for (let i = 0; i < data.categories.length; i++) {
                    this.Categories[i].negative_marks = this.StringToBoolean(data.categories[i].negative_marks)
                }
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
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkAddTeamResponse(data: any, name: string){
        if(data.code == 200){
            this.Teams.push({id: data.id, name: name, participants: [], leader: {name: "", surname: "", middlename: "", login: ""}});
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

    // Participants
    public GetParticipantsExcel(){
        window.location.href = "https://api.forcamp.ga/orgset.participants.password.get?token="+this.Token;
    }

    public GetParticipants(){
        this.http.get(this.GetParticipantsLink+"?token="+this.Token).subscribe((data: Response) => this.getParticipantsFromResponse(data.json()));
    }

    public EditParticipant(participant: Participant){
        this.PreloaderOn();
        this.http.post(this.EditParticipantLink,
            "token="+this.Token+
            "&login="+participant.login+
            "&name="+participant.name+
            "&surname="+participant.surname+
            "&middlename="+participant.middlename+
            "&sex="+participant.sex+
            "&team="+participant.team,
            { headers: this.PostHeaders }).subscribe((data: Response) => this.checkEditParticipantResponse(data.json(), participant));
    }

    public DeleteParticipant(login: string){
        this.PreloaderOn();
        this.http.post(this.DeleteParticipantLink, "token="+this.Token+"&login="+login, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkDeleteParticipantResponse(data.json(), login))
    }

    public ResetParticipantPassword(login: string){
        this.PreloaderOn();
        this.http.post(this.ResetParticipantPasswordLink, "token="+this.Token+"&login="+login, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkResetParticipantPasswordResponse(data.json()));
    }

    public AddParticipant(participant: Participant){
        this.PreloaderOn();
        this.http.post(this.AddParticipantLink,
            "token="+this.Token+
            "&name="+participant.name+
            "&surname="+participant.surname+
            "&middlename="+participant.middlename+
            "&sex="+participant.sex+
            "&team="+participant.team,
            { headers: this.PostHeaders }).subscribe((data: Response) => this.checkAddParticipantResponse(data.json(), participant));
    }

    public GetParticipantFullNameByLogin(participant_login: string): string {
        try {
            let data = this.Participants.filter(part => part.login == participant_login)[0];
            return data.surname + " " + data.name + " " + data.middlename
        } catch (e) {
            return participant_login;
        }
    }

    public GetParticipantsByTeamID(team_id: number): Participant[] {
        if (team_id == -1) {
            return this.Participants;
        } else {
            return this.Participants.filter(part => part.team == team_id);
        }
    }

    private checkResetParticipantPasswordResponse(data: any){
        if(data.code == 200){
            alert({type: 1, text: "Новый пароль: "+data.password, stay: true});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkDeleteParticipantResponse(data: any, login: string){
        if(data.code == 200){
            for(let i = 0; i < this.Participants.length; i++){
                if(this.Participants[i].login == login){
                    if(this.Participants[i].team != 0){
                        for(let i = 0; i < this.Teams.length; i++){
                            if(this.Teams[i].id == this.Participants[i].team){
                                this.Teams[i].participants.splice(this.Teams[i].participants.indexOf(login), 1);
                                break
                            }
                        }
                    }
                    this.Participants.splice(i, 1);
                    break
                }
            }
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkEditParticipantResponse(data: any, participant: Participant){
        if(data.code == 200){
            for(let i = 0; i < this.Teams.length; i++){
                if(this.Teams[i].participants.indexOf(participant.login) != -1){
                    this.Teams[i].participants.splice(this.Teams[i].participants.indexOf(participant.login), 1);
                }
                if(this.Teams[i].id == participant.team){
                    this.Teams[i].participants.push(participant.login)
                }
            }
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkAddParticipantResponse(data: any, participant: Participant){
        if(data.code == 200){
            participant.login = data.login;
            participant.marks = [];
            for(let i = 0; i < this.Categories.length; i++){
                participant.marks.push({id: this.Categories[i].id, value: 0});
            }
            this.Participants.push(participant);
            if(participant.team != 0){
                for(let i = 0; i < this.Teams.length; i++){
                    if(this.Teams[i].id == participant.team){
                        this.Teams[i].participants.push(participant.login);
                        break
                    }
                }
            }
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private getParticipantsFromResponse(data: any){
        if(data.code == 200){
            this.Participants = data.participants;
            for (let i = 0; i < this.Participants.length; i++) {
                this.Participants[i].sum = 0;
                for (let j = 0; j < this.Participants[i].marks.length; j++) {
                    this.Participants[i].sum += this.Participants[i].marks[j].value;
                }
            }
            this.getParticipantsForVerticalBarChart();
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
    }

    private getParticipantsForVerticalBarChart(){
        let participantsVerticalBar = [];
        for (let i = 0; i < this.Participants.length; i++) {
            if (i == 10) {
                break
            }
            let surname = this.Participants[i].surname;
            surname = surname.charAt(0).toUpperCase() + surname.substr(1);
            let name = this.Participants[i].name;
            name = name.charAt(0).toUpperCase() + name.substr(1);
            participantsVerticalBar.push({name: surname + ' ' + name, value: this.Participants[i].sum});
        }
        participantsVerticalBar = participantsVerticalBar.sort((n1, n2) => { if (n1.value > n2.value) return -1; if (n1.value < n2.value) return 1; return 0});
        this.ParticipantsVerticalBar = participantsVerticalBar;
    }

    // Employees
    public GetEmployees(){
        this.http.get(this.GetEmployeesLink+"?token="+this.Token).subscribe((data: Response) => this.getEmployeesFromResponse(data.json()));
    }

    public EditEmployee(employee: Employee){
        this.PreloaderOn();
        this.http.post(this.EditEmployeeLink,
            "token="+this.Token+
            "&login="+employee.login+
            "&name="+employee.name+
            "&surname="+employee.surname+
            "&middlename="+employee.middlename+
            "&sex="+employee.sex+
            "&team="+employee.team+
            "&post="+employee.post,
            { headers: this.PostHeaders }).subscribe((data: Response) => this.checkEditEmployeeResponse(data.json(), employee));
    }

    public DeleteEmployee(login: string){
        this.PreloaderOn();
        this.http.post(this.DeleteEmployeeLink, "token="+this.Token+"&login="+login, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkDeleteEmployeeResponse(data.json(), login))
    }

    public ResetEmployeePassword(login: string){
        this.PreloaderOn();
        this.http.post(this.ResetEmployeePasswordLink, "token="+this.Token+"&login="+login, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkResetEmployeePasswordResponse(data.json()));
    }

    public AddEmployee(employee: Employee){
        this.PreloaderOn();
        this.http.post(this.AddEmployeeLink,
            "token="+this.Token+
            "&name="+employee.name+
            "&surname="+employee.surname+
            "&middlename="+employee.middlename+
            "&sex="+employee.sex+
            "&team="+employee.team+
            "&post="+employee.post,
            { headers: this.PostHeaders }).subscribe((data: Response) => this.checkAddEmployeeResponse(data.json(), employee));
    }

    public EditEmployeePermission(login: string, value: string, id: number){
        this.PreloaderOn();
        this.http.post(this.EditEmployeePermissionLink, "token="+this.Token+"&id="+id+"&value="+value+"&login="+login, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkEditEmployeePermissionResponse(data.json()));
    }

    public GetEmployeeFullNameByLogin(employee_login: string): string {
        try {
            let data = this.Employees.filter(empl => empl.login == employee_login)[0];
            return data.surname + " " + data.name + " " + data.middlename
        } catch (e) {
            return employee_login;
        }
    }

    public GetEmployeesByTeamID(team_id: number): Employee[] {
        if (team_id == -1) {
            return this.Employees;
        } else {
            return this.Employees.filter(part => part.team == team_id);
        }
    }

    private checkEditEmployeePermissionResponse(data: any){
        if(data.code == 200){
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkResetEmployeePasswordResponse(data: any){
        if(data.code == 200){
            alert({type: 1, text: "Новый пароль: "+data.password, stay: true});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkDeleteEmployeeResponse(data: any, login: string){
        if(data.code == 200){
            for(let i = 0; i < this.Employees.length; i++){
                if(this.Employees[i].login == login){
                    if(this.Employees[i].team != 0){
                        for(let i = 0; i < this.Teams.length; i++){
                            if(this.Teams[i].id == this.Employees[i].team){
                                this.Teams[i].leader = {name: '', surname: '', middlename: '', login: ''};
                                break
                            }
                        }
                    }
                    this.Employees.splice(i, 1);
                    break
                }
            }
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkEditEmployeeResponse(data: any, employee: Employee){
        if(data.code == 200){
            for(let i = 0; i < this.Teams.length; i++){
                if(this.Teams[i].leader.login == employee.login){
                    this.Teams[i].leader = {name: '', surname: '', middlename: '', login: ''};
                }
                if(this.Teams[i].id == employee.team){
                    this.Teams[i].leader = {name: employee.name, surname: employee.surname, middlename: employee.middlename, login: employee.login};
                }
            }
            for(let i = 0; i < this.Employees.length; i++){
                if(this.Employees[i].team == employee.team && this.Employees[i].login != employee.login){
                    this.Employees[i].team = 0;
                }
            }
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkAddEmployeeResponse(data: any, employee: Employee){
        if(data.code == 200){
            employee.login = data.login;
            employee.permissions = [];
            for(let i = 0; i < this.Categories.length; i++){
                employee.permissions.push({id: this.Categories[i].id, value: true});
            }
            if(employee.team != 0){
                for(let i = 0; i < this.Teams.length; i++){
                    if(this.Teams[i].id == employee.team){
                        this.Teams[i].leader = {name: employee.name, surname: employee.surname, middlename: employee.middlename, login: employee.login};
                        break
                    }
                }
            }
            this.Employees.push(employee);
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private getEmployeesFromResponse(data: any){
        if(data.code == 200){
            this.Employees = data.employees;
            if(this.Employees != null) {
                for (let i = 0; i < this.Employees.length; i++){
                    if(this.Employees[i].permissions != null) {
                        for (let k = 0; k < this.Employees[i].permissions.length; k++) {
                            this.Employees[i].permissions[k].value = this.StringToBoolean(data.employees[i].permissions[k].value)
                        }
                    }
                }
            }
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
    }

    // Reasons
    public GetReasons(){
        this.http.get(this.GetReasonsLink+"?token="+this.Token).subscribe((data: Response) => this.getReasonsFromResponse(data.json()));
    }

    public AddReason(catID: number, text: string, change: number){
        this.PreloaderOn();
        this.http.post(this.AddReasonLink, "token="+this.Token+"&cat_id="+catID+"&text="+text+"&change="+change, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkAddReasonResponse(data.json(), catID, text, change));
    }

    public EditReason(reason: Reason){
        this.PreloaderOn();
        this.http.post(this.EditReasonLink, "token="+this.Token+"&id="+reason.id+"&text="+reason.text+"&change="+reason.change+"&cat_id="+reason.cat_id, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkEditReasonResponse(data.json()));
    }

    public DeleteReason(id: number){
        this.PreloaderOn();
        this.http.post(this.DeleteReasonLink, "token="+this.Token+"&id="+id, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkDeleteReasonResponse(data.json(), id));
    }

    private getReasonsFromResponse(data: any){
        if(data.code == 200){
            this.Reasons = data.reasons;
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
    }

    private checkAddReasonResponse(data: any, catID: number, text: string, change: number){
        if(data.code == 200){
            this.Reasons.push({id: data.id, cat_id: catID, text: text, change: change});
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkEditReasonResponse(data: any){
        if(data.code == 200){
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    private checkDeleteReasonResponse(data: any, id: number){
        if(data.code == 200){
            for (let i = 0; i < this.Reasons.length; i++){
                if(this.Reasons[i].id == id){
                    this.Reasons.splice(i, 1);
                    break
                }
            }
            alert({type: 1, text: "Операция успешно завершена!", time: 2});
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
        this.PreloaderOff();
    }

    // MarksChanges
    public GetMarksChanges(){
        this.http.get(this.GetMarksChangesLink+"?token="+this.Token).subscribe((data: Response) => this.getMarksChangesFromResponse(data.json()));
    }

    public DeleteMarkChange(id: number){
        this.PreloaderOn();
        this.http.post(this.DeleteMarkChangeLink, "token="+this.Token+"&id="+id, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkDeleteMarkChangeResponse(data.json(), id));
    }

    public GetMarksChangesByEmployeeLogin(employee_login: string): MarksChange[] {
        return this.MarksChanges.filter(markChange => markChange.employee_login == employee_login);
    }

    public GetMarksChangesByParticipantLogin(participant_login: string): MarksChange[] {
        return this.MarksChanges.filter(markChange => markChange.participant_login == participant_login);
    }

    private getMarksChangesFromResponse(data: any){
        if(data.code == 200){
            this.MarksChanges = data.marks_changes;
        } else {
            alert({type: 3, text: "Произошла ошибка(" + data.code + ")!", time: 3});
        }
    }

    private checkDeleteMarkChangeResponse(data: any, id: number){
        if(data.code == 200){
            this.MarksChanges = this.MarksChanges.filter(markChange => markChange.id != id);
            this.GetData();
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
    public IntToSex(num: number): string{
        if(num == 0){
            return "мужской";
        } else {
            return "женский";
        }
    }

    public IdToTeamName(id: number): string{
        for(let i = 0; i < this.Teams.length; i++){
            if(this.Teams[i].id == id){
                return this.Teams[i].name
            }
        }
        return "отсутствует"
    }

    public CategoryIdToName(id: number){
        for(let i = 0; i < this.Categories.length; i++){
            if(this.Categories[i].id == id){
                return this.Categories[i].name;
            }
        }
        return "Ошибка!"
    }

    private StringToBoolean(data: string): boolean {
        if (data == "false") {
            return false;
        } else {
            return true;
        }
    }

    public GetReasonsByCatID(id: number){
        return this.Reasons.filter(reason => reason.cat_id == id);
    }

    public GetMarkByCategoryID(id: number, marks: any){
        return marks.filter(mark => mark.id == id);
    }

    public GetPermissionByCategoryID(id: number, permissions: any){
        return permissions.filter(permission => permission.id == id);
    }

}