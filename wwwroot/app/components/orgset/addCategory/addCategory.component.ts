import {Component} from '@angular/core';
import {OrgSetService} from "../../../src/orgset.service";
import {alert} from "notie";

@Component({
    selector: "add_category",
    templateUrl: "app/components/orgset/addCategory/addCategory.component.html",
    styleUrls: ["app/components/orgset/addCategory/addCategory.component.css"]
})
export class AddCategoryComponent{
    private CategoryName: string = '';

    constructor(public orgSetService: OrgSetService,) {

    }

    private AddCategorySubmit(){
        this.orgSetService.AddCategory_Active = false;
        this.orgSetService.AddCategory(this.CategoryName, false);
        this.CategoryName = '';
    }
}