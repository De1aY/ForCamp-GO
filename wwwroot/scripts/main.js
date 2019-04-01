window.global = {};
window.global.Token = $.cookie("token");
window.global.OrgSettings = {
    participant: "",
    period: "",
    organization: "",
    self_marks: "",
    team: "",
    emotional_mark_period: ""
};
window.global.Participants = [];
window.global.Categories = [];
window.global.Teams = [];
window.global.Reasons = [];
Preloader = new $.materialPreloader({
    position: 'top',
    height: '5px',
    col_1: '#159756',
    col_2: '#da4733',
    col_3: '#3b78e7',
    col_4: '#fdba2c',
    fadeIn: 200,
    fadeOut: 200
});

function GetOrganizationSettings() {
    return new Promise ( resolve => {
        $.get(__GetOrgSettingsLink, {token: window.global.Token}, function (resp) {
            if (resp.code === 200) {
                window.global.OrgSettings.participant = resp.message.settings.participant;
                window.global.OrgSettings.period = resp.message.settings.period;
                window.global.OrgSettings.organization = resp.message.settings.organization;
                window.global.OrgSettings.self_marks = resp.message.settings.self_marks;
                window.global.OrgSettings.team = resp.message.settings.team;
                window.global.OrgSettings.emotional_mark_period = resp.message.settings.emotional_mark_period;
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
            resolve();
        });
    });
}

function GetParticipants() {
    return new Promise ( resolve => {
        $.get(__GetParticipantsLink, {token: window.global.Token}, function (resp) {
            if (resp.code === 200) {
                resp.message.participants.forEach( participant => {
                    let sum = 0;
                    participant.marks.forEach( mark => {
                        sum += mark.value;
                    });
                    participant['sum'] = sum
                });
                window.global.Participants = resp.message.participants;
                window.global.Participants = window.global.Participants.sort( (a, b) => {
                    return b.sum - a.sum
                });
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
            resolve();
        });
    });
}

function GetCategories() {
    return new Promise ( resolve => {
        $.get(__GetCategoriesLink, {token: window.global.Token}, function (resp) {
            if (resp.code === 200) {
                window.global.Categories = resp.message.categories;
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
            resolve();
        });
    });
}

function GetLastMarkChanges(participant_id = "") {
    return new Promise( resolve => {
       $.get(__GetEventsLink, {token: window.global.Token,
           user_id: participant_id,
           rows_per_page: 20,
           event_type: 1,}, function (resp) {
          if (resp.code === 200) {
              resolve(resp.message.events);
          } else {
              resolve(null);
          }
       });
    });
}

function GetReasons() {
    return new Promise ( resolve => {
        $.get(__GetReasonsLink, {token: window.global.Token}, function (resp) {
            if (resp.code === 200) {
                window.global.Reasons = resp.message.reasons;
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
            resolve();
        });
    });
}

function GetTeams() {
    return new Promise ( resolve => {
        $.get(__GetTeamsLink, {token: window.global.Token}, function (resp) {
            if (resp.code === 200) {
                window.global.Teams = resp.message.teams;
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
            resolve();
        });
    });
}

function GetTeamNameByID(team_id) {
    try {
        return window.global.Teams.filter(team => {
            return team.id === team_id;
        })[0].name;
    } catch (e) {
        return "отсутствует";
    }
}

function GetSexByID(id) {
    if (id === 0) {
        return "мужской";
    } else {
        return "женский";
    }
}

function GetCategoryNameByID(category_id) {
    if (category_id === "000") {
        return "Выберите категорию";
    }
    try {
        let name = window.global.Categories.filter(category => {
            return category.id === category_id;
        })[0].name;
        return name[0].toUpperCase() + name.substring(1);
    } catch (e) {
        return "Произошла ошибка";
    }
}

function GetUserData(user_id = "") {
    return new Promise( resolve => {
        $.get(__GetUserDataLink, { token: window.global.Token, user_id: user_id}, function(resp) {
           if (resp.code === 200) {
               resolve(resp.message.data);
           } else {
               notie.alert({type: 3, text: resp.message.ru, time: 2});
               resolve(null);
           }
        });
    });
}

function SetTeamsMarks() {
    window.global.Teams.forEach(team => {
        team.marks = [];
        team.mark = 0;
        window.global.Participants[0].marks.forEach( mark => {
            team.marks.push(0)
        });
    });
    window.global.Participants.forEach( participant => {
        if (participant.team !== 0) {
            let team = window.global.Teams.filter(team => {
                return team.id === participant.team;
            })[0];
            if (team !== undefined) {
                participant.marks.forEach((mark, i) => {
                    team.mark += mark.value;
                    team.marks[i] += mark.value;
                });
            }
        }
    });
    window.global.Teams = window.global.Teams.sort( (a, b) => {
        return b.mark - a.mark
    });
}

function GetURLParameter(name) {
    return decodeURIComponent((new RegExp('[?|&]' + name + '=' + '([^&;]+?)(&|#|;|$)').exec(location.search) || [null, ''])[1].replace(/\+/g, '%20')) || null;
}

function ToTitleCase(str) {
    return str[0].toUpperCase() + str.substring(1);
}

function UploadFile() {
    try{
      let xml = new XMLHttpRequest();
      let args = arguments;
      let context = this;
      let multipart = "";
      xml.open(args[0].method,args[0].url,true);
      if(args[0].method.search(/post/i)!=-1){
        let boundary=Math.random().toString().substr(2);
        xml.setRequestHeader("content-type",
                    "multipart/form-data; charset=utf-8; boundary=" + boundary);
        for(let key in args[0].params){
          multipart += "--" + boundary
                     + "\r\nContent-Disposition: form-data; name=" + key
                     + "\r\nContent-type: application/octet-stream"
                     + "\r\n\r\n" + args[0].params[key] + "\r\n";
        }
        multipart += "--" + boundary
                    + '\r\nContent-Disposition: form-data; name="file"; filename="file"'
                    + "\r\nContent-type: image/png"
                    + "\r\n\r\n" + args[0].file + "\r\n";
        multipart += "--"+boundary+"--\r\n";
      }
      xml.onreadystatechange=function(){
        try{
          if(xml.readyState==4){
            context.txt = xml.responseText;
            context.xml = xml.responseXML;
            args[0].callback(JSON.parse(context.txt));
          }
        }
        catch(e){}
      }
      xml.send(multipart);
    }
    catch(e){}
}

function ShowModal(modal, display) {
    modal.fadeIn();
    modal.css('display', display);
}

function HideModal(modal) {
    modal.fadeOut();
}
