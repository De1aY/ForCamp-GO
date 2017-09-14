Token = $.cookie("token");
let OrgSettings = {
    participant: "",
    period: "",
    organization: "",
    self_marks: "",
    team: "",
    emotional_mark_period: ""
};
let Participants = [];
let Categories = [];
let Teams = [];
let Reasons = [];
let Preloader = new $.materialPreloader({
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
        $.get(__GetOrgSettingsLink, {token: Token}, function (resp) {
            if (resp.code === 200) {
                OrgSettings.participant = resp.message.settings.participant;
                OrgSettings.period = resp.message.settings.period;
                OrgSettings.organization = resp.message.settings.organization;
                OrgSettings.self_marks = resp.message.settings.self_marks;
                OrgSettings.team = resp.message.settings.team;
                OrgSettings.emotional_mark_period = resp.message.settings.emotional_mark_period;
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
            resolve();
        });
    });
}

function GetParticipants() {
    return new Promise ( resolve => {
        $.get(__GetParticipantsLink, {token: Token}, function (resp) {
            if (resp.code === 200) {
                resp.message.participants.forEach( participant => {
                    let sum = 0;
                    participant.marks.forEach( mark => {
                        sum += mark.value;
                    });
                    participant['sum'] = sum
                });
                Participants = resp.message.participants;
                Participants = Participants.sort( (a, b) => {
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

function GetLastMarkChanges(participant_id = "") {
    return new Promise( resolve => {
       $.get(__GetEventsLink, {token: Token,
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
        $.get(__GetReasonsLink, {token: Token}, function (resp) {
            if (resp.code === 200) {
                Reasons = resp.message.reasons;
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
                Teams = resp.message.teams;
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
            resolve();
        });
    });
}

function GetTeamNameByID(team_id) {
    try {
        return Teams.filter(team => {
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
        let name = Categories.filter(category => {
            return category.id === category_id;
        })[0].name;
        return name[0].toUpperCase() + name.substring(1);
    } catch (e) {
        return "Произошла ошибка";
    }
}

function GetUserData(user_id = "") {
    return new Promise( resolve => {
        $.get(__GetUserDataLink, { token: Token, user_id: user_id}, function(resp) {
           if (resp.code === 200) {
               resolve(resp.message.data);
           } else {
               notie.alert({type: 3, text: resp.message.ru, time: 2});
               resolve(null);
           }
        });
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
      var xml = new XMLHttpRequest();
      var args = arguments;
      var context = this;
      var multipart = "";
      xml.open(args[0].method,args[0].url,true);
      if(args[0].method.search(/post/i)!=-1){
        var boundary=Math.random().toString().substr(2);
        xml.setRequestHeader("content-type",
                    "multipart/form-data; charset=utf-8; boundary=" + boundary);
        for(var key in args[0].params){
          multipart += "--" + boundary
                     + "\r\nContent-Disposition: form-data; name=" + key
                     + "\r\nContent-type: application/octet-stream"
                     + "\r\n\r\n" + args[0].params[key] + "\r\n";
        }
        multipart += "--" + boundary
                    + "\r\nContent-Disposition: form-data; name=file filename=file"
                    + "\r\nContent-type: image/png"
                    + "\r\n\r\n" + args[0].file + "\r\n";
        multipart += "--"+boundary+"--\r\n";
      }
      xml.onreadystatechange=function(){
        try{
          if(xml.readyState==4){
            context.txt=xml.responseText;
            context.xml=xml.responseXML;
            args[0].callback();
          }
        }
        catch(e){}
      }
      xml.send(multipart);
    }
    catch(e){}
}