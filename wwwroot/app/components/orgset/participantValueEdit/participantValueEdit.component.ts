import {Component} from '@angular/core';
import {OrgSetService} from "../../../src/orgset.service";
import {alert} from "notie";

@Component({
    selector: "participant_value_edit",
    templateUrl: "app/components/orgset/participantValueEdit/participantValueEdit.component.html",
    styleUrls: ["app/components/orgset/participantValueEdit/participantValueEdit.component.css"]
})
export class ParticipantValueEditComponent{
    private Participant_Value: string = '';

    constructor(public orgSetService: OrgSetService,) {

    }

    private ChangeParticipantValue(){
        this.orgSetService.ParticipantValueEdit_Active = false;
        this.orgSetService.SetOrgSettingValue("participant", this.Participant_Value);
        this.Participant_Value = '';
    }
}