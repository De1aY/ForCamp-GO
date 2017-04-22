import {Component} from '@angular/core';
import {OrgSetService} from "../../../src/orgset.service";
import {alert} from "notie";

@Component({
    selector: "period_value_edit",
    templateUrl: "app/components/orgset/periodValueEdit/periodValueEdit.component.html",
    styleUrls: ["app/components/orgset/periodValueEdit/periodValueEdit.component.css"]
})
export class PeriodValueEditComponent{
    private Period_Value: string = '';

    constructor(public orgSetService: OrgSetService,) {

    }

    private ChangePeriodValue(){
        this.orgSetService.PeriodValueEdit_Active = false;
        this.orgSetService.SetOrgSettingValue("period", this.Period_Value);
        this.Period_Value = '';
    }
}