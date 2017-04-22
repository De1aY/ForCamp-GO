import {Component, OnInit, ViewChild} from '@angular/core';
import {CookieService} from 'angular2-cookie/core';
import {CheckTokenService} from '../../src/checkToken.service';
import {UserService} from '../../src/user.service';
import {OrgSetService} from "../../src/orgset.service";
import {alert} from "notie";
import {DatatableComponent} from "@swimlane/ngx-datatable";

interface Category {
    id: number
    name: string
    negative_marks: boolean
}

@Component({
    selector: "org_set",
    templateUrl: "app/components/orgset/orgset.component.html",
    styleUrls: ["app/components/orgset/orgset.component.css"]
})
export class OrgSetComponent implements OnInit {
    private Token: string;
    private CategoryEdit: object = {};
    private TeamEdit: object = {};
    private ParticipantEdit: object = {};
    private EmployeeEdit: object = {};
    private ReasonEdit: object = {};
    private ParticipantsTable_FilterTeam:number = -1;

    constructor(private cookieService: CookieService,
                private checkTokenService: CheckTokenService,
                public userService: UserService,
                public orgSetService: OrgSetService,) {
    }

    ngOnInit() {
        this.TokenInit();
        this.ServiceInit();
    }

    private ServiceInit(){
        this.UserServiceInit();
        this.OrgSetServiceInit();
    }

    private OrgSetServiceInit(){
        if(this.orgSetService.Token == undefined){
            this.orgSetService.Token = this.Token;
        }
        this.orgSetService.GetData();
    }

    private UserServiceInit(){
        if(this.userService.Token == undefined) {
            this.userService.Token = this.Token;
        }
        this.userService.GetData();
    }

    private TokenInit(){
        this.Token = this.cookieService.get("token");
        if (this.Token == undefined) {
            window.location.href = "https://forcamp.ga/exit";
        }
        this.checkTokenService.CheckToken(this.Token);
    }

    public ChangeSelfMarksValue(self_marks: any){
        this.orgSetService.SetOrgSettingValue("self_marks", self_marks.checked);
    }

}