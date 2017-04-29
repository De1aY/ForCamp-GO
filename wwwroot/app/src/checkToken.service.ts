import {Injectable, Inject} from "@angular/core";
import {Http, Response} from "@angular/http";

@Injectable()
export class CheckTokenService {
    private Link:string = "https://api.forcamp.ga/token.verify";
    public AdminStatus: boolean = false;

    constructor(
      @Inject(Http) private http: Http,
    ){}

    public CheckToken(Token: string){
        this.http.get(this.Link+"?token="+Token).subscribe((data: Response) => this.checkToken(data.json()))
    }

    private checkToken(data: any){
          if(data.code != 200){
              window.location.href = "https://forcamp.ga/exit";
          } else {
              this.AdminStatus = data.admin_status;
          }
    }
}