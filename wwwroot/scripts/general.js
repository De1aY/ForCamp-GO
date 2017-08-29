GetTeams();

function reloadTables() {
    Top5Table.ajax.reload(null, false);
}

$(document).ready(function() {
    $(".form-control").after("<span></span>");
});

function onTableDraw() {
    componentHandler.upgradeDom();
}

// General - Top5

let Top5Table = $('#mdl-card__body-table-top5').DataTable({
    "lengthChange": false,
    "paging": false,
    "searching": false,
    "ordering": false,
    "ajax": {
        "url": __GetParticipantsLink,
        "type": "GET",
        "data": {
            "token": Token,
        },
        "dataSrc": function (data) {
            data.message.participants.forEach( participant => {
                let sum = 0;
                participant.marks.forEach( mark => {
                    sum += mark.value;
                });
                participant['sum'] = sum
            });
            Participants = data.message.participants;
            Participants = Participants.sort( (a, b) => {
               return b.sum - a.sum
            });
            return Participants.slice(0, 6);
        },
    },
    columnDefs: [
        {
            targets: 0,
            name: "full_name",
            className: 'mdl-data-table__cell--non-numeric',
            data: "surname",
            orderable: false,
            render: function ( surname, type, row, meta ) {
                return '<a href="https://forcamp.ga/profile?id=' + row.id + '" ' +
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
            orderable: false,
            render: function ( team_id, type, row, meta ) {
                let teamName = GetTeamNameByID(team_id);
                return '<span>' + teamName[0].toUpperCase() + teamName.substring(1) + '</span>';
            }
        },
        {
            targets: 2,
            name: "sum",
            data: "sum",
            type: "num",
            render: function ( sum, type, row, meta ) {
                return sum
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
    "initComplete": function() {
        drawCharts();
    }
});

// General - BubbleChart

let general_chart_bubble_ctx = document.getElementById('general-chart--top5').getContext('2d');

async function drawCharts() {
    let general_chart_top5_data = getGeneralTop5ChartData();
    let general_chart_top5 = new Chart(general_chart_bubble_ctx, {
        type: 'bar',
        data: {
            labels: general_chart_top5_data.labels,
            datasets: general_chart_top5_data.datasets
        },
        options: {
            responsive: true,
            legend: {
                display: false,
            },
        }
    });
}

function getGeneralTop5ChartData() {
    let chartData = {
        datasets: [],
        labels: [],
    };
    Participants[0].marks.forEach( mark => {
       chartData.labels.push(mark.name[0].toUpperCase()+mark.name.substring(1));
    });
    Participants.slice(0, 6).forEach( participant => {
        let dataset = {
            label:  participant.surname[0].toUpperCase() + participant.surname.substring(1) + ' '
            + participant.name[0].toUpperCase() + participant.name.substring(1) + ' '
            + participant.middlename[0].toUpperCase() + participant.middlename.substring(1),
            data: [],
            borderColor: "",
            backgroundColor: "",
        };
        participant.marks.forEach( mark => {
            dataset.data.push(mark.value);
        });
        let rColor = randomColor({
            luminosity: 'bright',
        });
        dataset.borderColor = rColor;
        dataset.backgroundColor = rColor;
        chartData.datasets.push(dataset);
    });
    return chartData
}