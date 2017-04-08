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
var marks_service_1 = require("../../../src/marks.service");
var EditMarkComponent = (function () {
    function EditMarkComponent(orgSetService, marksService) {
        this.orgSetService = orgSetService;
        this.marksService = marksService;
        this.Reason = "";
        this.Value = 0;
        this.CategoryName = '';
    }
    EditMarkComponent.prototype.EditMarkSubmit = function () {
        this.marksService.MarkEdit_Active = false;
        this.marksService.EditParticipantMark(this.login, this.categoryID, this.Value, this.Reason);
        this.CategoryName = '';
    };
    return EditMarkComponent;
}());
__decorate([
    core_1.Input(),
    __metadata("design:type", String)
], EditMarkComponent.prototype, "login", void 0);
__decorate([
    core_1.Input(),
    __metadata("design:type", Number)
], EditMarkComponent.prototype, "categoryID", void 0);
EditMarkComponent = __decorate([
    core_1.Component({
        selector: "edit_mark",
        templateUrl: "app/components/marks/marksEdit/marksEdit.component.html",
        styleUrls: ["app/components/marks/marksEdit/marksEdit.component.css"]
    }),
    __metadata("design:paramtypes", [orgset_service_1.OrgSetService,
        marks_service_1.MarksService])
], EditMarkComponent);
exports.EditMarkComponent = EditMarkComponent;
//# sourceMappingURL=marksEdit.component.js.map