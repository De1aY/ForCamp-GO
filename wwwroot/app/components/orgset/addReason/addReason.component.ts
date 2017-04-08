import {Component} from '@angular/core';
import {OrgSetService} from "../../../src/orgset.service";
import {alert} from "notie";

@Component({
    selector: "add_reason",
    templateUrl: "app/components/orgset/addReason/addReason.component.html",
    styleUrls: ["app/components/orgset/addReason/addReason.component.css"]
})
export class AddReasonComponent{
    private text = '';
    private change = 0;
    private catID = 0;

    constructor(public orgSetService: OrgSetService,) {

    }

    private AddReasonSubmit(){
        this.orgSetService.AddReason_Active = false;
        this.orgSetService.AddReason(this.catID, this.text, this.change);
    }
}