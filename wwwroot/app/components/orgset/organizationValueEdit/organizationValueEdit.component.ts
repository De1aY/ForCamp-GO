import {Component} from '@angular/core';
import {OrgSetService} from "../../../src/orgset.service";
import {alert} from "notie";

@Component({
    selector: "organization_value_edit",
    templateUrl: "app/components/orgset/organizationValueEdit/organizationValueEdit.component.html",
    styleUrls: ["app/components/orgset/organizationValueEdit/organizationValueEdit.component.css"]
})
export class OrganizationValueEditComponent{
    private Organization_Value: string = '';

    constructor(public orgSetService: OrgSetService,) {

    }

    private ChangeOrganizationValue(){
        this.orgSetService.OrganizationValueEdit_Active = false;
        this.orgSetService.SetOrgSettingValue("organization", this.Organization_Value);
        this.Organization_Value = '';
    }
}