GetOrganizationSettings();
GetTeams();
GetReasons();

$(document).ready(function() {
    $(".form-control").after("<span></span>");
});

function reloadTables() {
    MarksTable.ajax.reload(null, false);
}

function onTableDraw() {
    componentHandler.upgradeDom();
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
            switch (editInfo[0]) {
                case "participant": {
                    let login = editInfo[1];
                    let category_id = editInfo[3];
                    let reason_id = editInfo[5];
                    editMark(login, category_id, reason_id);
                    break;
                }
                case "action": {
                    let name = $('#mdl-card__body-table-employees--name-' + editInfo[1]).text();
                    let surname = $('#mdl-card__body-table-employees--surname-' + editInfo[1]).text();
                    let middlename = $('#mdl-card__body-table-employees--middlename-' + editInfo[1]).text();
                    let post = $('#mdl-card__body-table-employees--post-' + editInfo[1]).text();
                    let sex = $('#mdl-card__body-table-employees--sex-' + editInfo[1]).data('content').split('-')[3];
                    let team = $('#mdl-card__body-table-employees--team-' + editInfo[1]).data('content').split('-')[3];
                    editEmployee(name, surname, middlename, post, sex, team, editInfo[1]);
                    break;
                }
            }
        });
    });
}

// Marks

function editMark(login, category_id, reason_id) {
    $.post(__EditMark, { token: Token,
        login: login,
        category_id: category_id,
        reason_id: reason_id}, function (resp) {
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            reloadTables();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

let MarksTable = $('#mdl-card__body-table-marks').DataTable({
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
            name: "full_name",
            className: 'mdl-data-table__cell--non-numeric',
            data: "surname",
            searchable: true,
            render: function ( surname, type, row, meta ) {
                return '<a href="https://forcamp.ga/profile?login=' + row.login + '" ' +
                    'class="mdl-card__body-table-row__field">'+
                    surname[0].toUpperCase() + surname.substring(1) + ' '
                    + row.name[0].toUpperCase() + row.name.substring(1) + ' '
                    + row.middlename[0].toUpperCase() + row.middlename.substring(1) + '</a>';
            },
        },
        {
            targets: 1,
            name: "team",
            className: 'mdl-data-table__cell--non-numeric',
            data: "team",
            orderable: true,
            render: function ( team_id, type, row, meta ) {
                let teamName = GetTeamNameByID(team_id);
                return '<span>' + teamName[0].toUpperCase() + teamName.substring(1) + '</span>';
            }
        },
        {
            targets: '_all',
            name: "mark",
            className: 'mdl-card__body-table-row_actions',
            data: "marks",
            searchable: false,
            orderable: true,
            render: function ( marks, type, row, meta ) {
                let mark = marks[meta.col - 2];
                let dropdown = '<div class="mdl-card__body-table-row__dropdown mdl-card__body-table-row__field--capitalize user-select--none" ' +
                    'id="mdl-card__body-table-participants--mark-'+row.login+'"' +
                    ' data-content="participant-' + row.login + '-mark-' + mark.id + '">' +
                    '<div class="mdl-card__body-table-row__dropdown-ttl">' + mark.value +'</div><ul>' +
                    '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="close-0"><span>Закрыть</span></li>';
                let reasons = Reasons.filter(function (reason) {
                    return reason.category_id === mark.id;
                });
                reasons.forEach(function (reason) {
                    dropdown += '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="participant-'
                        + row.login + '-mark-' + mark.id + '-reason-' + reason.id + '">' +
                        '<span>' + reason.text + '</span></li>';
                });
                return dropdown + '</ul></label></div>';
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
        $('#mdl-card__body-table-marks').css('width', '100%');
        onTableDraw();
    },
});