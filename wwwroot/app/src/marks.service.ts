import {Injectable, Inject} from "@angular/core";
import {Http, Response, Headers} from "@angular/http";
import {alert} from "notie";
import {OrgSetService} from "./orgset.service";
import {UserService} from "./user.service";
import {Data} from "@angular/router";

@Injectable()
export class MarksService {
    private EditUserMarkLink:string = "https://api.forcamp.ga/mark.edit";
    private PostHeaders: Headers = new Headers();
    private Time: Date;
    public Preloader: boolean = false;
    public Token: string;

    constructor(
        @Inject(Http) private http: Http,
        public orgSetService: OrgSetService,
        public userService: UserService,
    ){
        this.PostHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
    }

    public EditParticipantMark(participant: any, id: number, reason: number){
        this.PreloaderOn();
        this.http.post(this.EditUserMarkLink, "token="+this.Token+"&login="+participant.login+"&category_id="+id+"&reason_id="+reason, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkEditParticipantMark(data.json(), participant, id, reason));
    }

    private checkEditParticipantMark(data: any, participant: any, id: number, reason_id: number){
        if (data.code == 200) {
            let reason = this.orgSetService.Reasons.filter(row => row.id == reason_id)[0];
            participant.marks.filter(row => row.id == id)[0].value += reason.change;
            this.Time = new Date();
            this.Time.setDate(this.Time.getDate());
            this.orgSetService.MarksChanges.push({id: id, employee_login: this.userService.SelfLogin, participant_login: participant.login, text: reason.text, time: this.Time.toUTCString(), change: reason.change});
            this.orgSetService.GetData();
            alert({type: 1, text: "Балл успешно изменён!", time: 2});
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
}