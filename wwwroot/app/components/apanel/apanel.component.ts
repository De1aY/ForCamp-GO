import {Component, OnInit} from "@angular/core"
import {CookieService} from "angular2-cookie/services/cookies.service";
import {ApanelService} from "../../src/apanel.service";
import {CheckTokenService} from "../../src/checkToken.service";

@Component({
    selector: "apanel",
    templateUrl: "app/components/apanel/apanel.component.html",
    styleUrls: ["app/components/apanel/apanel.component.css"]
})
export class ApanelComponent implements OnInit{
    private Token: string;
    private OrgName: string = "";

    constructor(private cookieService: CookieService,
                public ApanelService: ApanelService) {
    }

    ngOnInit() {
        this.TokenInit();
        this.ServiceInit();
    }

    private ServiceInit() {
        this.ApanelService.Token = this.Token;
        this.ApanelService.ServiceInit();
    }

    private TokenInit(){
        this.Token = this.cookieService.get("token");
        if (this.Token == undefined) {
            window.location.href = "https://forcamp.ga/exit";
        }
    }

    private CreateOrganization() {
        this.ApanelService.CreateOrganization(this.OrgName);
    }

}