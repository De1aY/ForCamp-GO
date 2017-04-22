import {Component} from '@angular/core';
import {OrgSetService} from "../../../src/orgset.service";
import {alert} from "notie";

interface MarkPermission{
    id: number
    value: boolean
}

interface Employee{
    login: string
    name: string
    surname: string
    middlename: string
    sex: number
    team: number
    post: string
    permissions: MarkPermission[]
}

@Component({
    selector: "add_employee",
    templateUrl: "app/components/orgset/addEmployee/addEmployee.component.html",
    styleUrls: ["app/components/orgset/addEmployee/addEmployee.component.css"]
})
export class AddEmployeeComponent{
    private employee: Employee = {
        name: "",
        surname: "",
        middlename: "",
        sex: 0,
        team: 0,
        login: "",
        post: "",
        permissions: []
    };

    constructor(public orgSetService: OrgSetService,) {

    }

    private AddEmployeeSubmit(){
        this.orgSetService.AddEmployee_Active = false;
        this.orgSetService.AddEmployee(this.employee);
        this.employee = {
            name: "",
            surname: "",
            middlename: "",
            sex: 0,
            team: 0,
            login: "",
            post: "",
            permissions: []
        };
    }
}