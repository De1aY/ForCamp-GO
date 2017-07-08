Token = $.cookie("token");
let OrgSettings = {
    participant: "",
    period: "",
    organization: "",
    self_marks: "",
    team: ""
};
let Categories = {};
let Teams = {};

function GetOrganizationSettings() {
    return new Promise ( resolve => {
        $.get(__GetOrgSettingsLink, {token: Token}, function (resp) {
            if (resp.code === 200) {
                OrgSettings.participant = resp.message.settings.participant;
                OrgSettings.period = resp.message.settings.period;
                OrgSettings.organization = resp.message.settings.organization;
                OrgSettings.self_marks = resp.message.settings.self_marks;
                OrgSettings.team = resp.message.settings.team;
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
            resolve();
        });
    });
}

function GetCategories() {
    return new Promise ( resolve => {
        $.get(__GetCategoriesLink, {token: Token}, function (resp) {
            if (resp.code === 200) {
                Categories = resp.message.categories;
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
            resolve();
        });
    });
}

function GetTeams() {
    return new Promise ( resolve => {
        $.get(__GetTeamsLink, {token: Token}, function (resp) {
            if (resp.code === 200) {
                Categories = resp.message.teams;
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
            resolve();
        });
    });
}

function GetTeamNameByID(team_id) {
    if (team_id === 0) {
        return "отсутствует";
    } else {
        return Teams.filter(team => {
            return team.id === team_id;
        })[0].name;
    }
}

function GetSexByID(id) {
    if (id === 0) {
        return "мужской";
    } else {
        return "женский";
    }
}