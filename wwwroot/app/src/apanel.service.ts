import {Injectable, Inject} from "@angular/core";
import {Http, Response, Headers} from "@angular/http";

@Injectable()
export class ApanelService {
    // Links
    private CreateOrganizationLink:string = "https://api.forcamp.ga/apanel.organization.add";
    private TokenVerifyLink:string = "https://api.forcamp.ga/token.verify";
    // Vars
    private PostHeaders: Headers = new Headers();
    public Token: string;


    constructor(@Inject(Http) private http: Http,){
        this.PostHeaders.append('Content-Type', 'application/x-www-form-urlencoded');
    }

    public ServiceInit() {
        this.checkToken();
    }

    public CreateOrganization(orgname: string) {
        this.http.post(this.CreateOrganizationLink, "token="+this.Token+"&orgname="+orgname, { headers: this.PostHeaders}).subscribe((data: Response) => this.checkCreateOrganizationResponse(data.json));
    }

    private checkToken(){
        this.http.get(this.TokenVerifyLink+"?token="+this.Token).subscribe((data: Response) => this.checkTokenResponse(data.json()))
    }

    private checkTokenResponse(data: any){
        if(data.code != 200){
            window.location.href = "https://forcamp.ga/exit";
        } else {
            if (data.admin_status != true) {
                window.location.href = "https://forcamp.ga";
            }
        }
    }

    private checkCreateOrganizationResponse(data: any) {
        if (data.code == 200) {
            alert(data.login+" "+data.password);
        } else {
            alert(data.code);
        }
    }
}