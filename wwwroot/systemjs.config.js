(function (global) {
    System.config({
        paths: {
            'npm:': 'node_modules/'
        },
        map: {
            app: 'app',
            // пакеты angular
            '@angular/core': 'npm:@angular/core/bundles/core.umd.js',
            '@angular/common': 'npm:@angular/common/bundles/common.umd.js',
            '@angular/compiler': 'npm:@angular/compiler/bundles/compiler.umd.js',
            '@angular/platform-browser': 'npm:@angular/platform-browser/bundles/platform-browser.umd.js',
            '@angular/platform-browser-dynamic': 'npm:@angular/platform-browser-dynamic/bundles/platform-browser-dynamic.umd.js',
            '@angular/http': 'npm:@angular/http/bundles/http.umd.js',
            '@angular/router': 'npm:@angular/router/bundles/router.umd.js',
            '@angular/forms': 'npm:@angular/forms/bundles/forms.umd.js',
            '@angular/material':         'npm:@angular/material/bundles/material.umd.js',
            '@swimlane/ngx-datatable': 'npm:@swimlane/ngx-datatable/release/index.js',
            // остальные пакеты
            'rxjs':                      'npm:rxjs',
            'angular-in-memory-web-api': 'npm:angular-in-memory-web-api/bundles/in-memory-web-api.umd.js',
            'angular2-cookie':           'npm:angular2-cookie',
            'notie':                     'npm:notie/dist/notie.js',
            'hammerjs':                  'npm:hammerjs/hammer.js',
        },
        packages: {
            app: {
                main: './main.js',
                defaultExtension: 'js'
            },
            rxjs: {
                defaultExtension: 'js'
            },
            notie: {
                defaultExtension: 'js'
            },
            'angular2-in-memory-web-api': {
                main: './index.js',
                defaultExtension: 'js'
            },
            'angular2-cookie': {
                main: './core.js',
                defaultExtension: 'js'
            },
        },
    });
})(this);