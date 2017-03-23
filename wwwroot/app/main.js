"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const platform_browser_dynamic_1 = require("@angular/platform-browser-dynamic");
const authorization_module_1 = require("./components/authorization/authorization.module");
const orgpanel_module_1 = require("./components/orgpanel/orgpanel.module");
const platform = platform_browser_dynamic_1.platformBrowserDynamic();
let Loc = window.location.pathname;
if (Loc == "/" || Loc == "/index.html") {
    platform.bootstrapModule(authorization_module_1.AuthorizationModule);
}
if (Loc == "/main.html") {
    platform.bootstrapModule(orgpanel_module_1.OrgPanelModule);
}
//# sourceMappingURL=main.js.map