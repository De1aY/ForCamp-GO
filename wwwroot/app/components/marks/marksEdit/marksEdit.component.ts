import {Input, Component} from '@angular/core';
import {OrgSetService} from "../../../src/orgset.service";
import {alert} from "notie";
import {MarksService} from "../../../src/marks.service";

@Component({
    selector: "edit_mark",
    templateUrl: "app/components/marks/marksEdit/marksEdit.component.html",
    styleUrls: ["app/components/marks/marksEdit/marksEdit.component.css"]
})
export class EditMarkComponent{
    @Input() login: string;
    @Input() categoryID: number;
    private Reason: string = "";
    private Value: number = 0;
    private CategoryName: string = '';

    constructor(
        public orgSetService: OrgSetService,
        public marksService: MarksService
    ) {}

    private EditMarkSubmit(){
        this.marksService.MarkEdit_Active = false;
        this.marksService.EditParticipantMark(this.login, this.categoryID, this.Value, this.Reason);
        this.CategoryName = '';
    }
}