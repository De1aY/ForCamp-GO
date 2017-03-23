import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
import { AuthorizationModule } from "./components/authorization/authorization.module";
import { OrgPanelModule } from "./components/orgpanel/orgpanel.module";

const platform = platformBrowserDynamic();
let Loc: string = window.location.pathname;
if(Loc == "/" || Loc == "/index.html") {
    platform.bootstrapModule(AuthorizationModule);
}
if(Loc == "/main.html") {
    platform.bootstrapModule(OrgPanelModule);
}