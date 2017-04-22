import {Component} from '@angular/core';
import {OrgSetService} from "../../../src/orgset.service";
import {alert} from "notie";

interface Mark{
    id: number
    value: number
}

interface Participant{
    login: string
    name: string
    surname: string
    middlename: string
    sex: number
    team: number
    marks: Mark[]
}

@Component({
    selector: "add_participant",
    templateUrl: "app/components/orgset/addParticipant/addParticipant.component.html",
    styleUrls: ["app/components/orgset/addParticipant/addParticipant.component.css"]
})
export class AddParticipantComponent{
    private participant: Participant = {
        name: "",
        surname: "",
        middlename: "",
        sex: 0,
        team: 0,
        login: "",
        marks: []
    };

    constructor(public orgSetService: OrgSetService,) {

    }

    private AddParticipantSubmit(){
        this.orgSetService.AddParticipant_Active = false;
        this.orgSetService.AddParticipant(this.participant);
        this.participant = {
            name: "",
            surname: "",
            middlename: "",
            sex: 0,
            team: 0,
            login: "",
            marks: []
        };
    }
}