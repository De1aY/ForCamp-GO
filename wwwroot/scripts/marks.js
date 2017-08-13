GetTeams();
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
            render: function ( surname, type, row, meta ) {
                return '<a href="https://forcamp.ga/profile?login=' + row.login + '" ' +
                    'class="mdl-card__body-table-row__field" ' +
                    'id="mdl-card__body-table-participants--surname-'+row.login+'"' +
                    ' data-content="participant-'+row.login+'-surname">'+
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
            searchable: true,
            orderable: true,
            render: function ( team_id, type, row, meta ) {
                let dropdown = '<div class="mdl-card__body-table-row__dropdown mdl-card__body-table-row__field--capitalize" ' +
                    'id="mdl-card__body-table-participants--team-'+row.login+'"' +
                    ' data-content="participant-' + row.login + '-team-' + team_id + '">' +
                    '<div class="mdl-card__body-table-row__dropdown-ttl">' + GetTeamNameByID(team_id) +'</div><ul>'+
                    '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="participant-' + row.login + '-team-0">' +
                    '<span>отсутствует</span></li>';
                Teams.forEach(function (team) {
                    dropdown += '<li class="mdl-card__body-table-row__dropdown-field wave-effect" data-content="participant-' + row.login + '-team-' + team.id + '">' +
                        '<span>' + team.name + '</span></li>';
                });
                return dropdown + '</ul></label></div>';
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
                return '<span>' + marks[meta.col - 2].value + '</span>';
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
    },
});