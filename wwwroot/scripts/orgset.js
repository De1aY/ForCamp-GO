Token = $.cookie("token");
let OrgSettings = {
    participant: "",
    period: "",
    organization: "",
    self_marks: "",
    team: ""
};
getOrganizationSettings();
setInterval(getOrganizationSettings, 10000);

function getOrganizationSettings() {
    $.get(__GetOrgSettingsLink, { token: Token }, function(resp) {
        if(resp.code === 200) {
            OrgSettings.organization = resp.message.settings.organization;
            OrgSettings.period = resp.message.settings.period;
            OrgSettings.organization = resp.message.settings.organization;
            OrgSettings.self_marks = resp.message.settings.self_marks;
            OrgSettings.team = resp.message.settings.team;
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 3});
        }
    });
}