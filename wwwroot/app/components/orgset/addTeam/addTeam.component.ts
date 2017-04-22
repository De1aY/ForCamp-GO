import {Component} from '@angular/core';
import {OrgSetService} from "../../../src/orgset.service";
import {alert} from "notie";

@Component({
    selector: "add_team",
    templateUrl: "app/components/orgset/addTeam/addTeam.component.html",
    styleUrls: ["app/components/orgset/addTeam/addTeam.component.css"]
})
export class AddTeamComponent{
    private TeamName: string = '';

    constructor(public orgSetService: OrgSetService,) {

    }

    private AddTeamSubmit(){
        this.orgSetService.AddTeam_Active = false;
        this.orgSetService.AddTeam(this.TeamName);
        this.TeamName = '';
    }
}