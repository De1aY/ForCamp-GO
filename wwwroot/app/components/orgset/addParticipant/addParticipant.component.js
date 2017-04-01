"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
Object.defineProperty(exports, "__esModule", { value: true });
var core_1 = require("@angular/core");
var orgset_service_1 = require("../../../src/orgset.service");
var AddParticipantComponent = (function () {
    function AddParticipantComponent(orgSetService) {
        this.orgSetService = orgSetService;
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
    AddParticipantComponent.prototype.AddParticipantSubmit = function () {
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
    };
    return AddParticipantComponent;
}());
AddParticipantComponent = __decorate([
    core_1.Component({
        selector: "add_participant",
        templateUrl: "app/components/orgset/addParticipant/addParticipant.component.html",
        styleUrls: ["app/components/orgset/addParticipant/addParticipant.component.css"]
    }),
    __metadata("design:paramtypes", [orgset_service_1.OrgSetService])
], AddParticipantComponent);
exports.AddParticipantComponent = AddParticipantComponent;
//# sourceMappingURL=addParticipant.component.js.map