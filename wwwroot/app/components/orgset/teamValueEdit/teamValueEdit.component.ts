import {Component} from '@angular/core';
import {OrgSetService} from "../../../src/orgset.service";
import {alert} from "notie";

@Component({
    selector: "team_value_edit",
    templateUrl: "app/components/orgset/teamValueEdit/teamValueEdit.component.html",
    styleUrls: ["app/components/orgset/teamValueEdit/teamValueEdit.component.css"]
})
export class TeamValueEditComponent{
    private Team_Value: string = '';

    constructor(public orgSetService: OrgSetService,) {

    }

    private ChangeTeamValue(){
        this.orgSetService.TeamValueEdit_Active = false;
        this.orgSetService.SetOrgSettingValue("team", this.Team_Value);
        this.Team_Value = '';
    }
}