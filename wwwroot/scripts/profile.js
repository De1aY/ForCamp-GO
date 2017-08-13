let MarksTable = $('#mdl-card__body-table-marks').DataTable({
    "searching": false,
    "paging": false,
    "lengthChange": false,
    "ajax": {
        "url": __GetUserDataLink,
        "type": "GET",
        "data": {
            "token": Token,
            "login": "participant_11",
        },
        "dataSrc": function (data) {
            return data.message.data.additional_data;
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
            name: "value",
            className: 'mdl-data-table__cell--non-numeric',
            data: "value",
            render: function ( name, type, row, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-teams--name-'+row.id+'"' +
                    ' data-content="team-'+row.id+'-name">'+
                    row.value+'</div>';
            },
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
        "loadingRecords": "Загрузка баллов...",
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
        $('#mdl-card__body-table-marks').css('width', '100%');
    },
});