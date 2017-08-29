// Emotional mark

$('document').ready(function() {
    $('.mdl-card__body-row-emotional_mark').mousedown(function () {
        let emotionalMark_value = $(this).data('content');
        let sentiment = $(this).text();
        $.post(__SetEmotionalMarkLink, {token: Token, value: emotionalMark_value}, function (resp) {
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Успешно", time: 2});
                $('#card-emotional_mark').fadeOut();
                let oldEmotionalMark = $('.mdl-card__title-photo-emotional_mark');
                oldEmotionalMark.text(sentiment);
                oldEmotionalMark.removeClass('emotional_mark--very_dissatisfied emotional_mark--dissatisfied' +
                    'emotional_mark--neutral emotional_mark--satisfied emotional_mark--very_satisfied');
                switch (sentiment) {
                    case "mood":
                        oldEmotionalMark.addClass('emotional_mark--very_satisfied');
                        break;
                    case "sentiment_satisfied":
                        oldEmotionalMark.addClass('emotional_mark--satisfied');
                        break;
                    case "sentiment_neutral":
                        oldEmotionalMark.addClass('emotional_mark--neutral');
                        break;
                    case "sentiment_dissatisfied":
                        oldEmotionalMark.addClass('emotional_mark--dissatisfied');
                        break;
                    case "mood_bad":
                        oldEmotionalMark.addClass('emotional_mark--very_dissatisfied');
                        break;
                    default:
                        oldEmotionalMark.addClass('emotional_mark--very_satisfied');
                        break;
                }
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
        });
    });
});

// Charts

let marksChanges_lineChart_ctx = document.getElementById('marks_changes-chart--line').getContext('2d');

async function drawCharts() {
    let marksChanges_lineChart_data = await getMarksChangesLineChartData();
    let marksChanges_lineChart = new Chart(marksChanges_lineChart_ctx, {
        type: 'line',
        data: {
            labels: marksChanges_lineChart_data.labels,
            datasets: [{
                borderColor: "#3f51b5",
                backgroundColor: "#3f51ff",
                fill: false,
                data: marksChanges_lineChart_data.data,
            }],
        },
        options: {
            responsive: true,
            legend: {
                display: false,
            },
            scales: {
                xAxes: [{
                    display: false,
                }]
            }
        }
    });
}

async function getMarksChangesLineChartData() {
    return new Promise(async function (resolve) {
        let requestUserData = await GetUserData(RequestUserID);
        let marksSum = 0;
        requestUserData.additional_data.marks.forEach(mark => {
            marksSum += mark.value;
        });
        let events = await GetLastMarkChanges(RequestUserID);
        let chartData = {
            labels: [],
            data: [],
        };
        events.forEach(event => {
           chartData.labels.push(event.time);
           chartData.data.push(marksSum);
           marksSum -= event.event_data.change;
        });
        chartData.labels.reverse();
        chartData.data.reverse();
        resolve(chartData);
    });
}

drawCharts();