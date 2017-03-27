import {Directive, AfterViewChecked} from '@angular/core';

declare let componentHandler: any;

@Directive({
    selector: '[mdl]'
})
export class MDL implements AfterViewChecked {

    ngAfterViewChecked() {
        if (componentHandler) {
            componentHandler.upgradeAllRegistered();
        }
    }

}