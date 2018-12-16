updateSettings();

function reloadTables() {
    TeamsTable.ajax.reload(null, false);
    CategoriesTable.ajax.reload(null, false);
    ParticipantsTable.ajax.reload(null, false);
    EmployeesTable.ajax.reload(null ,false);
    ReasonsTable.ajax.reload(null, false);
    EventsTable.ajax.reload(null, false);
}

function getDefaultPermissions() {
    let permissions = [];
    Categories.forEach(category => {
       permissions.push({id: category.id, name: category.name, value: "true"});
    });
    return permissions;
}

$(document).ready(function(){
    $(".form-control").after("<span></span>");
    $('.mdl-card__menu-button--add').mouseup(async function () {
        let button = $(this);
        let editInfo = button.data('content');
        button.toggle();
        switch (editInfo) {
            case "category": {
                CategoriesTable.row.add({id: "000", name: "Новая категория", negative_marks: "false"}).search("Новая категория").draw();
                $('.mdl-card__body-table-row__field#mdl-card__body-table-categories--name-000').dblclick();
                break;
            }
            case "team": {
                TeamsTable.row.add({id: "000", name: "Новая команда", leader: {id: ""}, participants: []}).search("Новая команда").draw();
                $('.mdl-card__body-table-row__field#mdl-card__body-table-teams--name-000').dblclick();
                break;
            }
            case "participant": {
                ParticipantsTable.row.add({
                    id: "000",
                    name: "Имя",
                    surname: "Фамилия",
                    middlename: "Отчество",
                    sex: 0,
                    team: 0
                }).search("Фамилия Имя Отчество").draw();
                $('.mdl-card__body-table-row__field#mdl-card__body-table-participants--surname-000').dblclick();
                break;
            }
            case "employee": {
                EmployeesTable.row.add({
                    id: "000",
                    name: "Имя",
                    surname: "Фамилия",
                    middlename: "Отчество",
                    post: "должность",
                    sex: 0,
                    team: 0,
                    permissions: getDefaultPermissions(),
                }).search("Фамилия Имя Отчество должность").draw();
                $('.mdl-card__body-table-row__field#mdl-card__body-table-employees--surname-000').dblclick();
                break;
            }
            case "reason": {
                ReasonsTable.row.add({
                    id: "000",
                    category_id: "000",
                    text: "Новая причина",
                    change: 0
                }).search("Новая причина").draw();
                $('.mdl-card__body-table-row__field#mdl-card__body-table-reasons--text-000').dblclick();
            }
        }
    });
    $('.mdl-card__menu-button--download').click(function () {
        let button = $(this);
        let info = button.data('content');
        switch (info) {
            case "participants": {
                window.location.href = "https://api.forcamp.nullteam.info/orgset.participants.password.get?token="+Token;
                break;
            }
            case "employees": {
                window.location.href = "https://api.forcamp.nullteam.info/orgset.employees.password.get?token="+Token;
                break;
            }
        }
    });
});

function onTableDraw() {
    componentHandler.upgradeDom();
    $('.mdl-card__body-table-row__field:not(.mdl-card__body-table-row__field--noteditable)').unbind('dblclick').dblclick(function () {
        let textField = $(this);
        let baseText = textField.text();
        let editInfo = textField.data('content').split('-');
        textField.addClass('mdl-card__body-table-row__field--editing');
        textField.attr('contenteditable', 'true');
        textField.focus();
        textField.off('focusout keydown').on('focusout keydown',function (e) {
            if(e.keyCode === 13 || e.type === "focusout") {
                if (textField.attr('contenteditable') === "false") {
                    return
                }
                textField.removeClass('mdl-card__body-table-row__field--editing');
                textField.attr('contenteditable', 'false');
                let text = textField.text();
                if (text !== baseText && editInfo[1] !== "000") {
                    switch (editInfo[0]) {
                        case "category": {
                            let name = $('#mdl-card__body-table-categories--name-' + editInfo[1]).text();
                            let negative_marks = $('#mdl-card__body-table-categories--negative_marks-' + editInfo[1]).prop('checked');
                            editCategory(name, negative_marks, editInfo[1]);
                            break;
                        }
                        case "team": {
                            let name = $('#mdl-card__body-table-teams--name-'+editInfo[1]).text();
                            editTeam(name, editInfo[1]);
                            break;
                        }
                        case "participant": {
                            let name = $('#mdl-card__body-table-participants--name-'+editInfo[1]).text();
                            let surname = $('#mdl-card__body-table-participants--surname-'+editInfo[1]).text();
                            let middlename = $('#mdl-card__body-table-participants--middlename-'+editInfo[1]).text();
                            let sex = $('#mdl-card__body-table-participants--sex-'+editInfo[1]).data('content').split('-')[3];
                            let team = $('#mdl-card__body-table-participants--team-'+editInfo[1]).data('content').split('-')[3];
                            editParticipant(name, surname, middlename, sex, team, editInfo[1]);
                            break;
                        }
                        case "employee": {
                            let name = $('#mdl-card__body-table-employees--name-'+editInfo[1]).text();
                            let surname = $('#mdl-card__body-table-employees--surname-'+editInfo[1]).text();
                            let middlename = $('#mdl-card__body-table-employees--middlename-'+editInfo[1]).text();
                            let post = $('#mdl-card__body-table-employees--post-'+editInfo[1]).text();
                            let sex = $('#mdl-card__body-table-employees--sex-'+editInfo[1]).data('content').split('-')[3];
                            let team = $('#mdl-card__body-table-employees--team-'+editInfo[1]).data('content').split('-')[3];
                            editEmployee(name, surname, middlename, post, sex, team, editInfo[1]);
                            break;
                        }
                        case "reason": {
                            let text = $('#mdl-card__body-table-reasons--text-'+editInfo[1]).text();
                            let change = $('#mdl-card__body-table-reasons--change-'+editInfo[1]).text();
                            let category_id = $('#mdl-card__body-table-reasons--category-'+editInfo[1]).data('content').split('-')[3];
                            editReason(editInfo[1], text, change, category_id);
                            break;
                        }
                    }
                }
            }
        });
    });
    $('.mdl-card__body-table-row_switch').children('label').children('input').unbind('change').change( function () {
        let toggle = $(this).parents('label');
        let editInfo = toggle.attr('for').split('--')[1].split('-');
        let editObject = toggle.attr('for').split('--')[0].split('-')[3];
        if (editInfo[1] !== "000") {
            switch (editObject) {
                case "categories": {
                    let name = $('#mdl-card__body-table-categories--name-' + editInfo[1]).text();
                    let negative_marks = $('#mdl-card__body-table-categories--negative_marks-' + editInfo[1]).prop('checked').toString();
                    editCategory(name, negative_marks, editInfo[1]);
                    break;
                }
            }
        }
    });
    $('.mdl-card__body-table-row__dropdown').unbind('dblclick').dblclick( function () {
        let dropdownWrapper = $(this);
        dropdownWrapper.addClass('mdl-card__body-table-row__dropdown--editing');
        dropdownWrapper.children('ul').children('li').unbind('mousedown').mousedown(function () {
            dropdownWrapper.removeClass('mdl-card__body-table-row__dropdown--editing');
            let field = $(this);
            let editInfo = field.data('content').split('-');
            if (editInfo[0] === "close") {
                return
            }
            dropdownWrapper.children('.mdl-card__body-table-row__dropdown-ttl').text(field.text());
            dropdownWrapper.data('content', field.data('content'));
            if (editInfo[1] !== "000") {
                switch (editInfo[0]) {
                    case "participant": {
                        let name = $('#mdl-card__body-table-participants--name-' + editInfo[1]).text();
                        let surname = $('#mdl-card__body-table-participants--surname-' + editInfo[1]).text();
                        let middlename = $('#mdl-card__body-table-participants--middlename-' + editInfo[1]).text();
                        let sex = $('#mdl-card__body-table-participants--sex-' + editInfo[1]).data('content').split('-')[3];
                        let team = $('#mdl-card__body-table-participants--team-' + editInfo[1]).data('content').split('-')[3];
                        editParticipant(name, surname, middlename, sex, team, editInfo[1]);
                        break;
                    }
                    case "employee": {
                        let name = $('#mdl-card__body-table-employees--name-' + editInfo[1]).text();
                        let surname = $('#mdl-card__body-table-employees--surname-' + editInfo[1]).text();
                        let middlename = $('#mdl-card__body-table-employees--middlename-' + editInfo[1]).text();
                        let post = $('#mdl-card__body-table-employees--post-' + editInfo[1]).text();
                        let sex = $('#mdl-card__body-table-employees--sex-' + editInfo[1]).data('content').split('-')[3];
                        let team = $('#mdl-card__body-table-employees--team-' + editInfo[1]).data('content').split('-')[3];
                        editEmployee(name, surname, middlename, post, sex, team, editInfo[1]);
                        break;
                    }
                    case "reason": {
                        let text = $('#mdl-card__body-table-reasons--text-'+editInfo[1]).text();
                        let change = $('#mdl-card__body-table-reasons--change-'+editInfo[1]).text();
                        let category_id = $('#mdl-card__body-table-reasons--category-'+editInfo[1]).data('content').split('-')[3];
                        editReason(editInfo[1], text, change, category_id);
                        break;
                    }
                }
            }
        });
    });
    $('.mdl-card__body-table-row_actions--create').unbind('click').click(async function () {
       let button = $(this);
       let creationInfo = button.data('content').split('-');
       switch (creationInfo[0]) {
           case "team": {
               let name = $('#mdl-card__body-table-teams--name-'+creationInfo[1]).text();
               let id = await addTeam(name);
               if (id === -1) {
                   TeamsTable.row(button.parents('tr')).remove().search("").draw();
               } else {
                   TeamsTable.row(button.parents('tr')).remove();
                   TeamsTable.row.add({id: id, name: name, leader: {id: ""}, participants: []}).search("").draw();
               }
               break;
           }
           case "category": {
               let name = $('#mdl-card__body-table-categories--name-' + creationInfo[1]).text();
               let negative_marks = $('#mdl-card__body-table-categories--negative_marks-' + creationInfo[1]).prop('checked').toString();
               let id = await addCategory(name, negative_marks);
               if (id === -1) {
                   CategoriesTable.row(button.parents('tr')).remove().search("").draw();
               } else {
                   CategoriesTable.row(button.parents('tr')).remove();
                   CategoriesTable.row.add({id: id, name: name, negative_marks: negative_marks}).search("").draw();
               }
               break;
           }
           case "participant": {
               let name = $('#mdl-card__body-table-participants--name-'+creationInfo[1]).text();
               let surname = $('#mdl-card__body-table-participants--surname-'+creationInfo[1]).text();
               let middlename = $('#mdl-card__body-table-participants--middlename-'+creationInfo[1]).text();
               let sex = $('#mdl-card__body-table-participants--sex-'+creationInfo[1]).data('content').split('-')[3];
               let team = $('#mdl-card__body-table-participants--team-'+creationInfo[1]).data('content').split('-')[3];
               let id = await addParticipant(name, surname, middlename, sex, team);
               if (id === -1) {
                   ParticipantsTable.row(button.parents('tr')).remove().search("").draw();
               } else {
                   ParticipantsTable.row(button.parents('tr')).remove();
                   ParticipantsTable.row.add({
                       id: id,
                       name: name,
                       surname: surname,
                       middlename: middlename,
                       sex: sex,
                       team: team
                   }).search("").draw();
               }
               break;
           }
           case "employee": {
               let name = $('#mdl-card__body-table-employees--name-'+creationInfo[1]).text();
               let surname = $('#mdl-card__body-table-employees--surname-'+creationInfo[1]).text();
               let middlename = $('#mdl-card__body-table-employees--middlename-'+creationInfo[1]).text();
               let post = $('#mdl-card__body-table-employees--post-'+creationInfo[1]).text();
               let sex = $('#mdl-card__body-table-employees--sex-'+creationInfo[1]).data('content').split('-')[3];
               let team = $('#mdl-card__body-table-employees--team-'+creationInfo[1]).data('content').split('-')[3];
               let id = await addEmployee(name, surname, middlename, post,sex, team);
               if (id === -1) {
                   EmployeesTable.row(button.parents('tr')).remove().search("").draw();
               } else {
                   EmployeesTable.row(button.parents('tr')).remove();
                   EmployeesTable.row.add({
                       id: id,
                       name: name,
                       surname: surname,
                       middlename: middlename,
                       post: post,
                       sex: sex,
                       team: team
                   }).search("").draw();
               }
               break;
           }
           case "reason": {
               let text = $('#mdl-card__body-table-reasons--text-'+creationInfo[1]).text();
               let change = $('#mdl-card__body-table-reasons--change-'+creationInfo[1]).text();
               let category_id = $('#mdl-card__body-table-reasons--category-'+creationInfo[1]).data('content').split('-')[3];
               let id = await addReason(text, change, category_id);
               if (id === -1) {
                   ReasonsTable.row(button.parents('tr')).remove().search("").draw();
               } else {
                   ReasonsTable.row(button.parents('tr')).remove();
                   ReasonsTable.row.add({
                       id: id,
                       category_id: category_id,
                       text: text,
                       change: change
                   }).search("").draw();
               }
               break;
           }
       }
        $('.mdl-card__menu-button--add[data-content='+creationInfo[0]+']').toggle();
    });
    $('.mdl-card__body-table-row_actions--decline').unbind('click').click(async function () {
        let button = $(this);
        let creationInfo = button.data('content').split('-');
        switch (creationInfo[0]) {
            case "team": {
                TeamsTable.row(button.parents('tr')).remove().search("").draw();
                break;
            }
            case "category": {
                CategoriesTable.row(button.parents('tr')).remove().search("").draw();
                break;
            }
            case "participant": {
                ParticipantsTable.row(button.parents('tr')).remove().search("").draw();
                break;
            }
            case "employee": {
                EmployeesTable.row(button.parents('tr')).remove().search("").draw();
                break;
            }
            case "reason": {
                ReasonsTable.row(button.parents('tr')).remove().search("").draw();
                break;
            }
        }
        $('.mdl-card__menu-button--add[data-content='+creationInfo[0]+']').toggle();
    });
    $('.mdl-card__body-table-row_actions--delete').unbind('click').click( function () {
        let button = $(this);
        let editInfo = button.data('content').split('-');
        switch (editInfo[0]){
            case "category": {
                deleteCategory(editInfo[1], button);
                break;
            }
            case "team": {
                deleteTeam(editInfo[1], button);
                break;
            }
            case "participant": {
                deleteParticipant(editInfo[1], button);
                break;
            }
            case "employee": {
                deleteEmployee(editInfo[1], button);
                break;
            }
            case "reason": {
                deleteReason(editInfo[1], button);
                break;
            }
            case "event": {
                deleteEvent(editInfo[1], button);
                break;
            }
        }
    });
    $('.mdl-card__body-table-row_actions--reset_password').unbind('click').click( function () {
        let button = $(this);
        let editInfo = button.data('content').split('-');
        switch (editInfo[0]){
            case "employee": {
                resetEmployeePassword(editInfo[1]);
                break;
            }
            case "participant": {
                resetParticipantPassword(editInfo[1]);
                break;
            }
        }
    });
    $('.mdl-card__body-table-row_actions--profile').unbind('click').click( function () {
       let button = $(this);
       let id = button.data('content').split('-')[1];
       window.location.href = "https://forcamp.nullteam.info/profile?id=" + id;
    });
    ActivateWavesEffect();
}

// Teams

function addTeam(name) {
    return new Promise(resolve => {
        Preloader.on();
        $.post(__AddTeamLink, { token: Token, name: name }, function (resp) {
            Preloader.off();
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
                reloadTables();
                resolve(resp.message.id);
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
                resolve(-1);
            }
        });
    });
}

function editTeam(name, team_id) {
    Preloader.on();
    $.post(__EditTeamLink, { token: Token, name: name, team_id: team_id}, function (resp) {
        Preloader.off();
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            reloadTables();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

function deleteTeam(team_id, button) {
    Preloader.on();
    $.post(__DeleteTeamLink, { token: Token, team_id: team_id }, function (resp) {
        Preloader.off();
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            TeamsTable.row(button.parents('tr')).remove().draw();
            reloadTables();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

let TeamsTable = $('#mdl-card__body-table-teams').DataTable({
    "ajax": {
        "url": __GetTeamsLink,
        "type": "GET",
        "data": {
            "token": Token,
        },
        "dataSrc": function (data) {
            Teams = data.message.teams;
            return data.message.teams;
        },
    },
    columnDefs: [
        {
            targets: 0,
            name: "name",
            className: 'mdl-data-table__cell--non-numeric',
            data: "name",
            render: function ( name, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-teams--name-'+row.id+'"' +
                    ' data-content="team-'+row.id+'-name">'+
                    name[0].toUpperCase()+name.substring(1)+'</div>';
            },
        },
        {
            targets: 1,
            name: "leader",
            className: 'mdl-data-table__cell--non-numeric',
            data: "leader",
            searchable: false,
            orderable: false,
            render: function ( leader, type, row, meta ) {
                if (leader.id > 0) {
                    return '<a class="mdl-card__body-table-row__field mdl-card__body-table-row__field--noteditable mdl-card__body-table-row__field--capitalize" ' +
                        'id="mdl-card__body-table-teams--leader-' + row.id + '"' +
                        ' data-content="team-' + row.id + '-leader" href="https://forcamp.nullteam.info/profile?id=' + leader.id + '">' +
                        leader.surname + ' ' + leader.name + ' ' + leader.middlename + '</a>';
                } else {
                    return '<div class="mdl-card__body-table-row__field mdl-card__body-table-row__field--noteditable mdl-card__body-table-row__field--capitalize" ' +
                        'id="mdl-card__body-table-teams--leader-' + row.id + '"' +
                        ' data-content="team-' + row.id + '-leader">отсутствует</div>';
                }
            },
        },
        {
            targets: 2,
            name: "participants",
            className: '',
            data: "participants",
            searchable: false,
            orderable: true,
            render: function ( participants, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field mdl-card__body-table-row__field--noteditable mdl-card__body-table-row__field--capitalize" ' +
                    'id="mdl-card__body-table-teams--participants-'+row.id+'"' +
                    ' data-content="team-'+row.id+'-participants">'+participants.length+'</div>';
            }
        },
        {
            targets: -1,
            name: "actions",
            className: 'mdl-card__body-table-row_actions',
            data: "id",
            searchable: false,
            orderable: false,
            render: function ( id, type, row, meta ) {
                if (row.id === "000") {
                    return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--create"' +
                        ' data-content="team-' + id + '"> ' +
                        '<i class="material-icons">done</i></button>' +
                        '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--decline"' +
                        ' data-content="team-' + id + '"> ' +
                        '<i class="material-icons">close</i></button>';
                } else {
                    return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--delete"' +
                        ' data-content="team-' + id + '"> ' +
                        '<i class="material-icons">delete_forever</i></button>';
                }
            }
        },
    ],
    language: {
        "processing": "Подождите...",
        "search": "Поиск:",
        "lengthMenu": "Показать _MENU_ записей",
        "info": "",
        "infoEmpty": "",
        "infoFiltered": "",
        "infoPostFix": "",
        "loadingRecords": "Загрузка команд...",
        "zeroRecords": "Команды отсутствуют.",
        "emptyTable": "Команды отсутствуют",
        "paginate": {
            "first": "Первая",
            "previous": "Пред.",
            "next": "След.",
            "last": "Последняя"
        },
        "aria": {
            "sortAscending": ": отсортировать по возрастанию",
            "sortDescending": ": отсортировать по убыванию"
        }
    },
    "drawCallback": function () {
        $('#mdl-card__body-table-teams').css('width', '100%');
        onTableDraw();
    },
});

// Settings

async function updateSettings() {
    await GetOrganizationSettings();
    $('.mdl-card__body-row-text[data-content=participant]').text(OrgSettings.participant);
    $('.mdl-card__body-row-text[data-content=period]').text(OrgSettings.period);
    $('.mdl-card__body-row-text[data-content=organization]').text(OrgSettings.organization);
    $('.mdl-card__body-row-text[data-content=team]').text(OrgSettings.team);
    $('.mdl-card__body-row-text[data-content=emotional_mark_period]').text(OrgSettings.emotional_mark_period);
    $('#fc-orgset__main-self_marks').prop('checked', OrgSettings.self_marks);
}

function editOrganizationSetting(settingName) {
    Preloader.on();
    $.post(__SetOrgSettingValueLink, { token: Token,
        setting_value: OrgSettings[settingName],
        setting_name: settingName,
    }, function (resp) {
        Preloader.off();
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены!", time: 2})
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
            updateSettings();
        }
    });
}

$('.mdl-card__body-row-switch').children('label').children('input').change(function () {
    let toggle = $(this).parents('label').parents('.mdl-card__body-row-switch');
    let content = toggle.data('content');
    OrgSettings[content] = $('#fc-orgset__main-'+content).prop('checked');
    editOrganizationSetting(content);
});

function submitSettingEdit(textField, editButton, acceptButton, declineButton, content, baseText) {
    textField.unbind('keydown');
    acceptButton.unbind('click');
    declineButton.unbind('click');
    textField.attr('contenteditable', 'false');
    textField.removeClass('mdl-card__body-row-text--editing');
    editButton.removeClass('mdl-card__body-row-button--off');
    acceptButton.addClass('mdl-card__body-row-button--off');
    declineButton.addClass('mdl-card__body-row-button--off');
    OrgSettings[content] = textField.text();
    if (OrgSettings[content] !== baseText) {
        editOrganizationSetting(content);
    }
}

function declineSettingEdit(textField, editButton, acceptButton, declineButton, baseText) {
    textField.unbind('keydown');
    acceptButton.unbind('click');
    declineButton.unbind('click');
    textField.attr('contenteditable', 'false');
    textField.removeClass('mdl-card__body-row-text--editing');
    editButton.removeClass('mdl-card__body-row-button--off');
    acceptButton.addClass('mdl-card__body-row-button--off');
    declineButton.addClass('mdl-card__body-row-button--off');
    textField.text(baseText);
}

$('.mdl-card__body-row-button--edit').click(function () {
    let editButton = $(this);
    let content = editButton.data('content');
    let acceptButton = $('.mdl-card__body-row-button--accept[data-content='+content+']');
    let declineButton = $('.mdl-card__body-row-button--decline[data-content='+content+']');
    let textField = $('.mdl-card__body-row-text[data-content='+content+']');
    let baseText = textField.text();
    textField.keydown(function (e) {
        if((e.keyCode) === 13) {
            submitSettingEdit(textField, editButton, acceptButton, declineButton, content, baseText);
        }
    });
    acceptButton.click(function () {
        submitSettingEdit(textField, editButton, acceptButton, declineButton, content, baseText);
    });
    declineButton.click(function () {
        declineSettingEdit(textField, editButton, acceptButton, declineButton, baseText);
    });
    textField.addClass('mdl-card__body-row-text--editing');
    textField.attr('contenteditable', 'true');
    textField.focus();
    editButton.addClass('mdl-card__body-row-button--off');
    acceptButton.removeClass('mdl-card__body-row-button--off');
    declineButton.removeClass('mdl-card__body-row-button--off');
});

// Categories

function addCategory(name, negative_marks) {
    return new Promise(resolve => {
        Preloader.on();
        $.post(__AddCategoryLink, { token: Token, name: name, negative_marks: negative_marks }, function (resp) {
            Preloader.off();
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
                reloadTables();
                resolve(resp.message.id);
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
                resolve(-1);
            }
        });
    });
}

function editCategory(name, negative_marks, category_id) {
    Preloader.on();
    $.post(__EditCategoryLink, { token: Token, name: name,
        negative_marks: negative_marks, category_id: category_id}, function (resp) {
            Preloader.off();
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
                reloadTables();
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
       }
    });
}

function deleteCategory(category_id, button) {
    Preloader.on();
    $.post(__DeleteCategoryLink, { token: Token, category_id: category_id }, function (resp) {
        Preloader.off();
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            CategoriesTable.row(button.parents('td')).remove().draw();
            reloadTables();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

let CategoriesTable = $('#mdl-card__body-table-categories').DataTable({
    "ajax": {
        "url": __GetCategoriesLink,
        "type": "GET",
        "data": {
            "token": Token,
        },
        "dataSrc": function (data) {
            Categories = data.message.categories;
            return data.message.categories;
        },
    },
    columnDefs: [
        {
            targets: 0,
            name: "name",
            className: 'mdl-data-table__cell--non-numeric',
            data: "name",
            render: function ( name, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-categories--name-'+row.id+'"' +
                    ' data-content="category-'+row.id+'-name">'+
                    name[0].toUpperCase()+name.substring(1)+'</div>';
            },
        },
        {
            targets: 1,
            name: "negative_marks",
            className: 'mdl-data-table__cell--non-numeric',
            data: "negative_marks",
            searchable: false,
            orderable: false,
            render: function ( negative_marks, type, row, meta ) {
                if(negative_marks === "true") {
                    return '<div class="mdl-card__body-table-row_switch"><label class="mdl-switch mdl-js-switch mdl-js-ripple-effect"' +
                        'for="mdl-card__body-table-categories--negative_marks-' + row.id + '">' +
                        '<input type="checkbox" id="mdl-card__body-table-categories--negative_marks-' + row.id + '" class="mdl-switch__input" checked> ' +
                        '<span class="mdl-switch__label"></span></label></div>';
                } else {
                    return '<div class="mdl-card__body-table-row_switch"><label class="mdl-switch mdl-js-switch mdl-js-ripple-effect"' +
                        'for="mdl-card__body-table-categories--negative_marks-' + row.id + '">' +
                        '<input type="checkbox" id="mdl-card__body-table-categories--negative_marks-' + row.id + '" class="mdl-switch__input"> ' +
                        '<span class="mdl-switch__label"></span></label></div>';
                }
            },
        },
        {
            targets: -1,
            name: "actions",
            className: 'mdl-card__body-table-row_actions',
            data: "id",
            searchable: false,
            orderable: false,
            render: function ( id, type, row, meta ) {
                if (id === "000") {
                    return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--create"' +
                        ' data-content="category-' + id + '"> ' +
                        '<i class="material-icons">done</i></button>' +
                        '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--decline"' +
                        ' data-content="category-' + id + '"> ' +
                        '<i class="material-icons">close</i></button>';
                } else {
                    return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--delete"' +
                        ' data-content="category-' + id + '"> ' +
                        '<i class="material-icons">delete_forever</i></button>';
                }
            }
        },
    ],
    language: {
        "processing": "Подождите...",
        "search": "Поиск:",
        "lengthMenu": "Показать _MENU_ записей",
        "info": "",
        "infoEmpty": "",
        "infoFiltered": "",
        "infoPostFix": "",
        "loadingRecords": "Загрузка категорий...",
        "zeroRecords": "Категории отсутствуют.",
        "emptyTable": "Категории отсутствуют",
        "paginate": {
            "first": "Первая",
            "previous": "Пред.",
            "next": "След.",
            "last": "Последняя"
        },
        "aria": {
            "sortAscending": ": отсортировать по возрастанию",
            "sortDescending": ": отсортировать по убыванию"
        }
    },
    "drawCallback": function () {
        $('#mdl-card__body-table-categories').css('width', '100%');
        onTableDraw();
    },
});

// Participants

function addParticipant(name, surname, middlename, sex, team) {
    return new Promise(resolve => {
        Preloader.on();
        $.post(__AddParticipantLink, { token: Token,
            name: name,
            surname: surname,
            middlename: middlename,
            sex: sex,
            team: team}, function (resp) {
                Preloader.off();
                if(resp.code === 200) {
                    notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
                    reloadTables();
                    resolve(resp.message.id);
                } else {
                    notie.alert({type: 3, text: resp.message.ru, time: 2});
                    resolve(-1);
               }
        });
    });
}

function editParticipant(name, surname, middlename, sex, team, participant_id) {
    Preloader.on();
    $.post(__EditParticipantLink, { token: Token,
        name: name,
        surname: surname,
        middlename: middlename,
        sex: sex,
        team: team,
        participant_id: participant_id}, function (resp) {
            Preloader.off();
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
                reloadTables();
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
    });
}

function deleteParticipant(participant_id, button) {
    Preloader.on();
    $.post(__DeleteParticipantLink, { token: Token, participant_id: participant_id }, function (resp) {
        Preloader.off();
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            ParticipantsTable.row(button.parents('td')).remove().draw();
            reloadTables();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

function resetParticipantPassword(participant_id) {
    Preloader.on();
    $.post(__ResetParticipantPasswordLink, { token: Token, participant_id: participant_id }, function (resp) {
        Preloader.off();
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Новый пароль: " + resp.message.password, time: 0});
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

let ParticipantsTable = $('#mdl-card__body-table-participants').DataTable({
    "ajax": {
        "url": __GetParticipantsLink,
        "type": "GET",
        "data": {
            "token": Token,
        },
        "dataSrc": function (data) {
            return data.message.participants;
        },
    },
    columnDefs: [
        {
            targets: 0,
            name: "surname",
            className: 'mdl-data-table__cell--non-numeric',
            data: "surname",
            render: function ( surname, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-participants--surname-'+row.id+'"' +
                    ' data-content="participant-'+row.id+'-surname">'+
                    surname[0].toUpperCase()+surname.substring(1)+'</div>';
            },
        },
        {
            targets: 1,
            name: "name",
            className: 'mdl-data-table__cell--non-numeric',
            data: "name",
            render: function ( name, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-participants--name-'+row.id+'"' +
                    ' data-content="participant-'+row.id+'-name">'+
                    name[0].toUpperCase()+name.substring(1)+'</div>';
            },
        },
        {
            targets: 2,
            name: "middlename",
            className: 'mdl-data-table__cell--non-numeric',
            data: "middlename",
            render: function ( middlename, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-participants--middlename-'+row.id+'"' +
                    ' data-content="participant-'+row.id+'-middlename">'+
                    middlename[0].toUpperCase()+middlename.substring(1)+'</div>';
            },
        },
        {
            targets: 3,
            name: "sex",
            className: 'mdl-data-table__cell--non-numeric',
            data: "sex",
            searchable: false,
            orderable: false,
            render: function ( sex, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__dropdown mdl-card__body-table-row__field--capitalize user-select--none" ' +
                    'id="mdl-card__body-table-participants--sex-'+row.id+'"' +
                    ' data-content="participant-' + row.id + '-sex-' + sex + '">' +
                    '<div class="mdl-card__body-table-row__dropdown-ttl">'+GetSexByID(sex)+'</div><ul>'+
                    '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="close-0"><span>Закрыть</span></li>'+
                    '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="participant-' + row.id + '-sex-0"><span>мужской</span></li>'+
                    '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="participant-' + row.id + '-sex-1"><span>женский</span></li>'+
                    '</ul></label></div>';
            },
        },
        {
            targets: 4,
            name: "team",
            className: 'mdl-data-table__cell--non-numeric',
            data: "team",
            searchable: true,
            orderable: true,
            render: function ( team_id, type, row, meta ) {
                let dropdown = '<div class="mdl-card__body-table-row__dropdown mdl-card__body-table-row__field--capitalize user-select--none" ' +
                'id="mdl-card__body-table-participants--team-'+row.id+'"' +
                ' data-content="participant-' + row.id + '-team-' + team_id + '">' +
                '<div class="mdl-card__body-table-row__dropdown-ttl">' + GetTeamNameByID(team_id) +'</div><ul>'+
                '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="close-0"><span>Закрыть</span></li>'+
                '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="participant-' + row.id + '-team-0">' +
                    '<span>отсутствует</span></li>';
                Teams.forEach(function (team) {
                    dropdown += '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="participant-' + row.id + '-team-' + team.id + '">' +
                        '<span>' + team.name + '</span></li>';
                });
                return dropdown + '</ul></label></div>';
            }
        },
        {
            targets: -1,
            name: "actions",
            className: 'mdl-card__body-table-row_actions',
            data: "id",
            searchable: false,
            orderable: false,
            render: function ( id, type, row, meta ) {
                if (id === "000") {
                    return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--create"' +
                        ' data-content="participant-' + id + '"> ' +
                        '<i class="material-icons">done</i></button>' +
                        '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--decline"' +
                        ' data-content="participant-' + id + '"> ' +
                        '<i class="material-icons">close</i></button>';
                } else {
                    return '<button class="mdl-button mdl-button--primary mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--profile"' +
                        ' data-content="participant-' + id + '"> ' +
                        '<i class="material-icons">person</i></button>' +
                        '<button class="mdl-button mdl-button--primary mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--reset_password"' +
                        ' data-content="participant-' + id + '"> ' +
                        '<i class="material-icons">refresh</i></button>' +
                        '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--delete"' +
                        ' data-content="participant-' + id + '"> ' +
                        '<i class="material-icons">delete_forever</i></button>';
                }
            }
        },
    ],
    language: {
        "processing": "Подождите...",
        "search": "Поиск:",
        "lengthMenu": "Показать _MENU_ записей",
        "info": "",
        "infoEmpty": "",
        "infoFiltered": "",
        "infoPostFix": "",
        "loadingRecords": "Загрузка участников...",
        "zeroRecords": "Участники отсутствуют.",
        "emptyTable": "Участники отсутствуют",
        "paginate": {
            "first": "Первая",
            "previous": "Пред.",
            "next": "След.",
            "last": "Последняя"
        },
        "aria": {
            "sortAscending": ": отсортировать по возрастанию",
            "sortDescending": ": отсортировать по убыванию"
        }
    },
    "drawCallback": function () {
        $('#mdl-card__body-table-participants').css('width', '100%');
        onTableDraw();
    },
});

// Employees

function addEmployee(name, surname, middlename, post, sex, team) {
    return new Promise(resolve => {
        Preloader.on();
        $.post(__AddEmployeeLink, { token: Token,
            name: name,
            surname: surname,
            middlename: middlename,
            post: post,
            sex: sex,
            team: team}, function (resp) {
                Preloader.off();
                if(resp.code === 200) {
                    notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
                    reloadTables();
                    resolve(resp.message.id);
                } else {
                    notie.alert({type: 3, text: resp.message.ru, time: 2});
                    resolve(-1);
                }
        });
    });
}

function editEmployee(name, surname, middlename, post, sex, team, employee_id) {
    Preloader.on();
    $.post(__EditEmployeeLink, { token: Token,
        name: name,
        surname: surname,
        middlename: middlename,
        sex: sex,
        team: team,
        post: post,
        employee_id: employee_id}, function (resp) {
            Preloader.off();
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
                reloadTables();
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
    });
}

function deleteEmployee(employee_id, button) {
    Preloader.on();
    $.post(__DeleteEmployeeLink, { token: Token, employee_id: employee_id }, function (resp) {
        Preloader.off();
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            EmployeesTable.row(button.parents('td')).remove().draw();
            reloadTables();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

function editEmployeePermission(employee_id, category_id, value) {
    Preloader.on();
    $.post(__EditEmployeePermissionLink, { token: Token, employee_id: employee_id,
        category_id: category_id, value: value }, function (resp) {
            Preloader.off();
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
    });
}

function resetEmployeePassword(employee_id) {
    Preloader.on();
    $.post(__ResetEmployeePasswordLink, { token: Token, employee_id: employee_id }, function (resp) {
        Preloader.off();
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Новый пароль: " + resp.message.password, time: 0});
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

let EmployeesTable = $('#mdl-card__body-table-employees').DataTable({
    "ajax": {
        "url": __GetEmployeesLink,
        "type": "GET",
        "data": {
            "token": Token,
        },
        "dataSrc": function (data) {
            return data.message.employees;
        },
    },
    columnDefs: [
        {
            targets: 0,
            name: "surname",
            className: 'mdl-data-table__cell--non-numeric',
            data: "surname",
            render: function ( surname, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-employees--surname-'+row.id+'"' +
                    ' data-content="employee-'+row.id+'-surname">'+
                    surname[0].toUpperCase()+surname.substring(1)+'</div>';
            },
        },
        {
            targets: 1,
            name: "name",
            className: 'mdl-data-table__cell--non-numeric',
            data: "name",
            render: function ( name, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-employees--name-'+row.id+'"' +
                    ' data-content="employee-'+row.id+'-name">'+
                    name[0].toUpperCase()+name.substring(1)+'</div>';
            },
        },
        {
            targets: 2,
            name: "middlename",
            className: 'mdl-data-table__cell--non-numeric',
            data: "middlename",
            render: function ( middlename, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-employees--middlename-'+row.id+'"' +
                    ' data-content="employee-'+row.id+'-middlename">'+
                    middlename[0].toUpperCase()+middlename.substring(1)+'</div>';
            },
        },
        {
            targets: 3,
            name: "post",
            className: 'mdl-data-table__cell--non-numeric',
            data: "post",
            orderable: false,
            render: function ( post, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-employees--post-'+row.id+'"' +
                    ' data-content="employee-'+row.id+'-post">'+
                    post[0].toUpperCase()+post.substring(1)+'</div>';
            },
        },
        {
            targets: 4,
            name: "sex",
            className: 'mdl-data-table__cell--non-numeric',
            data: "sex",
            searchable: false,
            orderable: false,
            render: function ( sex, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__dropdown mdl-card__body-table-row__field--capitalize user-select--none" ' +
                    'id="mdl-card__body-table-employees--sex-'+row.id+'"' +
                    ' data-content="employee-' + row.id + '-sex-' + sex + '">' +
                    '<div class="mdl-card__body-table-row__dropdown-ttl">'+GetSexByID(sex)+'</div><ul>'+
                    '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="close-0"><span>Закрыть</span></li>'+
                    '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="employee-' + row.id + '-sex-0"><span>мужской</span></li>'+
                    '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="employee-' + row.id + '-sex-1"><span>женский</span></li>'+
                    '</ul></label></div>';
            },
        },
        {
            targets: 5,
            name: "team",
            className: 'mdl-data-table__cell--non-numeric',
            data: "team",
            searchable: true,
            orderable: true,
            render: function ( team_id, type, row, meta ) {
                let dropdown = '<div class="mdl-card__body-table-row__dropdown mdl-card__body-table-row__field--capitalize user-select--none" ' +
                    'id="mdl-card__body-table-employees--team-'+row.id+'"' +
                    ' data-content="employee-' + row.id + '-team-' + team_id + '">' +
                    '<div class="mdl-card__body-table-row__dropdown-ttl">' + GetTeamNameByID(team_id) +'</div><ul>'+
                    '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="close-0"><span>Закрыть</span></li>'+
                    '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="employee-' + row.id + '-team-0">' +
                    '<span>отсутствует</span></li>';
                Teams.forEach(function (team) {
                    dropdown += '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="employee-' + row.id + '-team-' + team.id + '">' +
                        '<span>' + team.name + '</span></li>';
                });
                return dropdown + '</ul></label></div>';
            }
        },
        {
            targets: -2,
            name: "permissions",
            className: 'mdl-card__body-table-cell--addition',
            data: "permissions",
            searchable: false,
            orderable: false,
            render: function ( permissions, type, row, meta ) {
                let additionalContent = '<div class="mdl-card__body-table-row__switch-dropdown"><div class="mdl-card__body-table-row__switch-dropdown-ttl"></div><ul>'+
                    '<li><div class="mdl-card__body-table-row__switch-dropdown-button">Закрыть</div></li>';
                if (permissions !== undefined) {
                    permissions.forEach(permission => {
                        if(permission.value === "true") {
                            additionalContent += '<li><div><label class="mdl-checkbox mdl-js-checkbox mdl-js-ripple-effect" ' +
                                'for="mdl-card__body-table-employees--permission-' + permission.id + '-' + row.id + '">' +
                                '<input type="checkbox" id="mdl-card__body-table-employees--permission-' + permission.id + '-' + row.id + '" class="mdl-checkbox__input" checked> ' +
                                '<span class="mdl-checkbox__label">' + permission.name[0].toUpperCase() + permission.name.substring(1) + '</span></label></div></li>';
                        } else {
                            additionalContent += '<li><div><label class="mdl-checkbox mdl-js-checkbox mdl-js-ripple-effect" ' +
                                'for="mdl-card__body-table-employees--permission-' + permission.id + '-' + row.id + '">' +
                                '<input type="checkbox" id="mdl-card__body-table-employees--permission-' + permission.id + '-' + row.id + '" class="mdl-checkbox__input"> ' +
                                '<span class="mdl-checkbox__label">' + permission.name[0].toUpperCase() + permission.name.substring(1) + '</span></label></div></li>';
                        }
                    });
                }
                return additionalContent + '</ul></div>'
            }
        },
        {
            targets: -1,
            name: "actions",
            className: 'mdl-card__body-table-row_actions',
            data: "id",
            searchable: false,
            orderable: false,
            render: function ( id, type, row, meta ) {
                if (id === "000") {
                    return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--create"' +
                        ' data-content="employee-' + id + '"> ' +
                        '<i class="material-icons">done</i></button>' +
                        '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--decline"' +
                        ' data-content="employee-' + id + '"> ' +
                        '<i class="material-icons">close</i></button>';
                } else {
                    return '<button class="mdl-button mdl-button--primary mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--profile"' +
                        ' data-content="employee-' + id + '"> ' +
                        '<i class="material-icons">person</i></button>' +
                        '<button class="mdl-button mdl-button--primary mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--permissions"' +
                        ' data-content="employee-' + id + '"> ' +
                        '<i class="material-icons">vpn_key</i></button>' +
                        '<button class="mdl-button mdl-button--primary mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--reset_password"' +
                        ' data-content="employee-' + id + '"> ' +
                        '<i class="material-icons">refresh</i></button>' +
                        '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--delete"' +
                        ' data-content="employee-' + id + '"> ' +
                        '<i class="material-icons">delete_forever</i></button>';
                }
            }
        },
    ],
    language: {
        "processing": "Подождите...",
        "search": "Поиск:",
        "lengthMenu": "Показать _MENU_ записей",
        "info": "",
        "infoEmpty": "",
        "infoFiltered": "",
        "infoPostFix": "",
        "loadingRecords": "Загрузка сотрудников...",
        "zeroRecords": "Сотрудники отсутствуют.",
        "emptyTable": "Сотрудники отсутствуют",
        "paginate": {
            "first": "Первая",
            "previous": "Пред.",
            "next": "След.",
            "last": "Последняя"
        },
        "aria": {
            "sortAscending": ": отсортировать по возрастанию",
            "sortDescending": ": отсортировать по убыванию"
        }
    },
    "drawCallback": function () {
        $('#mdl-card__body-table-employees').css('width', '100%');
        $('.mdl-card__body-table-row_actions--permissions').unbind('mousedown').mousedown( function () {
            let button = $(this);
            let additionContent = button.parents('tr').children('.mdl-card__body-table-cell--addition').children('.mdl-card__body-table-row__switch-dropdown');
            if (additionContent.hasClass('mdl-card__body-table-row__switch-dropdown--editing')) {
                additionContent.removeClass('mdl-card__body-table-row__switch-dropdown--editing');
            } else {
                additionContent.addClass('mdl-card__body-table-row__switch-dropdown--editing');
            }
        });
        $('.mdl-card__body-table-row__switch-dropdown ul li label input').unbind("change").change(function () {
            let checkbox = $(this);
            let editInfo = checkbox.attr('id').split('-');
            let value = checkbox.prop('checked');
            editEmployeePermission(editInfo[7], editInfo[6], value);
        });
        $('.mdl-card__body-table-row__switch-dropdown ul li .mdl-card__body-table-row__switch-dropdown-button').unbind('click').click(function () {
           let button = $(this);
           button.parents('.mdl-card__body-table-row__switch-dropdown').removeClass('mdl-card__body-table-row__switch-dropdown--editing')
        });
        onTableDraw();
    },
});

// Reasons

function addReason(text, change, category_id) {
    return new Promise(resolve => {
        Preloader.on();
        $.post(__AddReasonLink, { token: Token, text: text, change: change, category_id: category_id }, function (resp) {
            Preloader.off();
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
                reloadTables();
                resolve(resp.message.id);
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
                resolve(-1);
            }
        });
    });
}

function editReason(reason_id, text, change, category_id) {
    Preloader.on();
    $.post(__EditReasonLink, { token: Token, text: text,
        change: change, category_id: category_id, reason_id: reason_id }, function (resp) {
            Preloader.off();
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
                reloadTables();
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
    });
}

function deleteReason(reason_id, button) {
    Preloader.on();
    $.post(__DeleteReasonLink, { token: Token, reason_id: reason_id }, function (resp) {
        Preloader.off();
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            ReasonsTable.row(button.parents('td')).remove().draw();
            reloadTables();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

let ReasonsTable = $('#mdl-card__body-table-reasons').DataTable({
    "ajax": {
        "url": __GetReasonsLink,
        "type": "GET",
        "data": {
            "token": Token,
        },
        "dataSrc": function (data) {
            return data.message.reasons;
        },
    },
    columnDefs: [
        {
            targets: 0,
            name: "category",
            className: 'mdl-data-table__cell--non-numeric',
            data: "category_id",
            searchable: false,
            orderable: true,
            render: function ( category_id, type, row, meta ) {
                let dropdown = '<div class="mdl-card__body-table-row__dropdown user-select--none" ' +
                    'id="mdl-card__body-table-reasons--category-'+row.id+'"' +
                    ' data-content="reason-' + row.id + '-category-' + category_id + '">' +
                    '<div class="mdl-card__body-table-row__dropdown-ttl">' + GetCategoryNameByID(category_id) +'</div><ul>' +
                    '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="close-0"><span>Закрыть</span></li>';
                Categories.forEach(function (category) {
                    dropdown += '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="reason-' + row.id + '-category-' + category.id + '">' +
                        '<span>' + category.name[0].toUpperCase() + category.name.substring(1) + '</span></li>';
                });
                return dropdown + '</ul></label></div>';
            },
        },
        {
            targets: 1,
            name: "text",
            className: 'mdl-data-table__cell--non-numeric',
            data: "text",
            searchable: true,
            orderable: false,
            render: function ( text, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-reasons--text-'+row.id+'"' +
                    ' data-content="reason-'+row.id+'-text">'+
                    text[0].toUpperCase()+text.substring(1)+'</div>';
            },
        },
        {
            targets: 2,
            name: "change",
            data: "change",
            className: 'mdl-data-table__cell--non-numeric',
            searchable: true,
            orderable: false,
            render: function ( change, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-reasons--change-'+row.id+'"' +
                    ' data-content="reason-'+row.id+'-change">'+change+'</div>';
            },
        },
        {
            targets: -1,
            name: "actions",
            className: 'mdl-card__body-table-row_actions',
            data: "id",
            searchable: false,
            orderable: false,
            render: function ( id, type, row, meta ) {
                if (id === "000") {
                    return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--create"' +
                        ' data-content="reason-' + id + '"> ' +
                        '<i class="material-icons">done</i></button>' +
                        '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--decline"' +
                        ' data-content="reason-' + id + '"> ' +
                        '<i class="material-icons">close</i></button>';
                } else {
                    return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--delete"' +
                        ' data-content="reason-' + id + '"> ' +
                        '<i class="material-icons">delete_forever</i></button>';
                }
            }
        },
    ],
    language: {
        "processing": "Подождите...",
        "search": "Поиск:",
        "lengthMenu": "Показать _MENU_ записей",
        "info": "",
        "infoEmpty": "",
        "infoFiltered": "",
        "infoPostFix": "",
        "loadingRecords": "Загрузка причин...",
        "zeroRecords": "Причины отсутствуют.",
        "emptyTable": "Причины отсутствуют",
        "paginate": {
            "first": "Первая",
            "previous": "Пред.",
            "next": "След.",
            "last": "Последняя"
        },
        "aria": {
            "sortAscending": ": отсортировать по возрастанию",
            "sortDescending": ": отсортировать по убыванию"
        }
    },
    "drawCallback": function () {
        $('#mdl-card__body-table-reasons').css('width', '100%');
        onTableDraw();
    },
});

// Events

function deleteEvent(event_id, button) {
    Preloader.on();
    $.post(__DeleteEventLink, { token: Token, event_id: event_id }, function (resp) {
        Preloader.off();
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            EventsTable.row(button.parents('tr')).remove().draw();
            reloadTables();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

let EventsTable = $('#mdl-card__body-table-events').DataTable({
    "ordering": false,
    "ajax": {
        "url": __GetEventsLink,
        "type": "GET",
        "data": {
            "token": Token,
            "type": 0,
            "rows_per_page": -1,
            "user_id": 0,
        },
        "dataSrc": function (data) {
            return data.message.events;
        },
    },
    columnDefs: [
        {
            targets: 0,
            name: "event",
            className: 'mdl-data-table__cell--non-numeric',
            data: "event_data",
            searchable: true,
            render: function ( event_data, type, row, meta ) {
                switch (row.type){
                    case 1:
                        let resultStr = '<a href="https://forcamp.nullteam.info/profile?id=' + row.employee_id + '">' +
                            ToTitleCase(event_data.employee.surname) + " " +
                            ToTitleCase(event_data.employee.name) + '</a>';
                        if (event_data.employee.sex === 1) {
                            resultStr += " изменила балл ";
                        } else {
                            resultStr += " изменилл балл ";
                        }
                        resultStr += '<a href="https://forcamp.nullteam.info/profile?id=' + row.participant_id + '">' +
                            ToTitleCase(event_data.participant.surname) + " " +
                            ToTitleCase(event_data.participant.name) + '</a> на ';
                        resultStr += event_data.change + ' по причине "' + event_data.text + '"';
                        return resultStr;
                    case 2:
                        let resultStr_Type2 = '<a href="https://forcamp.nullteam.info/profile?id=' + row.participant_id + '">' +
                            ToTitleCase(event_data.participant.surname) + " " +
                            ToTitleCase(event_data.participant.name) + '</a>';
                        if (event_data.participant.sex === 1) {
                            resultStr_Type2 += " оценила своё настроение на ";
                        } else {
                            resultStr_Type2 += " оценил своё настроение на ";
                        }
                        resultStr_Type2 += '"' + event_data.value + '"';
                        return resultStr_Type2;
                    default:
                        return '<span>Произошла ошибка</span>' 
                }
            },
        },
        {
            targets: 1,
            name: "date",
            className: 'mdl-card__body-table-row_actions',
            data: "time",
            searchable: true,
            orderable: true,
            render: function ( time, type, row, meta ) {
                let date = new Date(time);
                return '<span>' + date.toLocaleDateString() + '</span>';
            }
        },
        {
            targets: 2,
            name: "actions",
            className: 'mdl-card__body-table-row_actions',
            data: "id",
            searchable: false,
            orderable: false,
            render: function ( id, type, row, meta ) {
                return '<button class="mdl-button mdl-js-button mdl-button--icon mdl-js-ripple-effect mdl-card__body-table-row_actions--delete"' +
                    ' data-content="event-' + id + '"> ' +
                    '<i class="material-icons">delete_forever</i></button>';
            }
        },
    ],
    language: {
        "processing": "Подождите...",
        "search": "Поиск:",
        "lengthMenu": "Показать _MENU_ записей",
        "info": "",
        "infoEmpty": "",
        "infoFiltered": "",
        "infoPostFix": "",
        "loadingRecords": "Загрузка изменений...",
        "zeroRecords": "События отсутствуют.",
        "emptyTable": "События отсутствуют",
        "paginate": {
            "first": "Первая",
            "previous": "Пред.",
            "next": "След.",
            "last": "Последняя"
        },
        "aria": {
            "sortAscending": ": отсортировать по возрастанию",
            "sortDescending": ": отсортировать по убыванию"
        }
    },
    "drawCallback": function () {
        $('#mdl-card__body-table-events').css('width', '100%');
        onTableDraw();
    },
});