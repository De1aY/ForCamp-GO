import {Injectable, Inject} from "@angular/core";
import {Http, Response, Headers} from "@angular/http";
import {alert} from "notie";

@Injectable()
export class MarksService {
    private EditUserMarkLink:string = "https://api.forcamp.ga/mark.edit";
    private PostHeaders: Headers = new Headers();
    public Preloader: boolean = false;
    public Token: string;

    constructor(
        @Inject(Http) private http: Http,
    ){
        this.PostHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
    }

    public EditParticipantMark(login: string, id: number, reason: number){
        this.PreloaderOn();
        this.http.post(this.EditUserMarkLink, "token="+this.Token+"&login="+login+"&id="+id+"&reason="+reason, { headers: this.PostHeaders }).subscribe((data: Response) => this.checkEditParticipantMark(data.json()));
    }

    private checkEditParticipantMark(data: any){
        if (data.code == 200) {
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