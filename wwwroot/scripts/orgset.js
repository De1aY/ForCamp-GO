setInterval(updateSettings, 10000);

$(document).ready(function(){
    $(".form-control").after("<span></span>");
    $('.mdl-card__menu-button--add').mousedown(async function () {
        let button = $(this);
        let editInfo = button.data('content');
        switch (editInfo) {
            case "category": {
                let id = await addCategory("Новая категория", "false");
                if (id !== -1) {
                    CategoriesTable.row.add({id: id, name: "Новая категория", negative_marks: "false"}).draw();
                    $('.mdl-card__body-table-row__field#mdl-card__body-table-categories--name-'+id).dblclick();
                }
                break;
            }
        }
    });
});

// Settings

async function updateSettings() {
    await GetOrganizationSettings();
    $('.mdl-card__body-row-text[data-content=participant]').text(OrgSettings.participant);
    $('.mdl-card__body-row-text[data-content=period]').text(OrgSettings.period);
    $('.mdl-card__body-row-text[data-content=organization]').text(OrgSettings.organization);
    $('.mdl-card__body-row-text[data-content=team]').text(OrgSettings.team);
    $('#fc-orgset__main-self_marks').prop('checked', OrgSettings.self_marks);
}

function editOrganizationSettings() {
    $.post(__SetOrgSettingValueLink, { token: Token,
        participant: OrgSettings.participant,
        team: OrgSettings.team,
        organization: OrgSettings.organization,
        period: OrgSettings.period,
        self_marks: OrgSettings.self_marks
    }, function (resp) {
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены!", time: 2})
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 2});
        }
    });
}

$('.mdl-card__body-row-switch').mousedown(function () {
    let toggle = $(this);
    let content = toggle.data('content');
    OrgSettings[content] = $('#fc-orgset__main-'+content).prop('checked');
    editOrganizationSettings();
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
        editOrganizationSettings();
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
        $.post(__AddCategoryLink, { token: Token, name: name, negative_marks: negative_marks }, function (resp) {
            if(resp.code === 200) {
                resolve(resp.message.id);
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
                resolve(-1);
            }
        });
    });
}

function editCategory(name, negative_marks, id) {
    $.post(__EditCategoryLink, { token: Token, name: name, negative_marks: negative_marks, id: id}, function (resp) {
       if(resp.code === 200) {
           notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
       } else {
           notie.alert({type: 3, text: resp.message.ru, time: 2});
       }
    });
}

function deleteCategory(id, button) {
    $.post(__DeleteCategoryLink, { token: Token, id: id }, function (resp) {
        if(resp.code === 200) {
            notie.alert({type: 1, text: "Данные успешно изменены", time: 2});
            CategoriesTable.row(button.parents('tr')).remove().draw();
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
            return data.message.categories;
        },
    },
    columnDefs: [
        {
            targets: 0,
            name: "name",
            className: 'mdl-data-table__cell--non-numeric',
            data: "name",
            render: function ( name, type, full, meta ) {
                return '<div class="mdl-card__body-table-row__field" id="mdl-card__body-table-categories--name-'+full.id+'"' +
                    ' data-content="category-'+full.id+'-name">'+
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
                if(negative_marks === 'true') {
                    return '<label class="mdl-switch mdl-js-switch mdl-js-ripple-effect mdl-card__body-table-row_switch"' +
                        'for="mdl-card__body-table-categories--negative_marks-' + row.id + '">' +
                        '<input type="checkbox" id="mdl-card__body-table-categories--negative_marks-' + row.id + '" class="mdl-switch__input"> ' +
                        '<span class="mdl-switch__label"></span></label>';
                } else {
                    return '<label class="mdl-switch mdl-js-switch mdl-js-ripple-effect mdl-card__body-table-row_switch"' +
                        'for="mdl-card__body-table-categories--negative_marks-' + row.id + '">' +
                        '<input type="checkbox" id="mdl-card__body-table-categories--negative_marks-' + row.id + '" class="mdl-switch__input" checked> ' +
                        '<span class="mdl-switch__label"></span></label>';
                }
            },
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
                    ' data-content="category-'+id+'"> ' +
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
        componentHandler.upgradeDom();
        $('.mdl-card__body-table-row__field').unbind('dblclick').dblclick(function () {
            let textField = $(this);
            let baseText = textField.text();
            let editInfo = textField.data('content').split('-');
            textField.addClass('mdl-card__body-table-row__field--editing');
            textField.attr('contenteditable', 'true');
            textField.focus();
            textField.off('focusout keydown').on('focusout keydown',function (e) {
                if(e.keyCode === 13 || e.type === "focusout") {
                    textField.removeClass('mdl-card__body-table-row__field--editing');
                    textField.attr('contenteditable', 'false');
                    let text = textField.text();
                    if (text !== baseText) {
                        switch (editInfo[0]) {
                            case "category": {
                                let name = $('#mdl-card__body-table-categories--name-' + editInfo[1]).text();
                                let negative_marks = $('#mdl-card__body-table-categories--negative_marks-' + editInfo[1]).prop('checked');
                                editCategory(name, negative_marks, editInfo[1]);
                                break;
                            }
                        }
                    }
                }
            });
        });
        $('.mdl-card__body-table-row_switch').unbind('mousedown').mousedown( function () {
            let toggle = $(this);
            let editInfo = toggle.attr('for').split('--')[1].split('-');
            let editObject = toggle.attr('for').split('--')[0].split('-')[3];
            switch (editObject) {
                case "categories":
                    let name = $('#mdl-card__body-table-categories--name-' + editInfo[1]).text();
                    let negative_marks = $('#mdl-card__body-table-categories--negative_marks-' + editInfo[1]).prop('checked');
                    editCategory(name, negative_marks, editInfo[1]);
                    break;
            }
        });
        $('.mdl-card__body-table-row_actions--delete').unbind('mousedown').mousedown( function () {
           let button = $(this);
           let editInfo = button.data('content').split('-');
           switch (editInfo[0]){
               case "category": {
                   deleteCategory(editInfo[1], button);
                   break;
               }
           }
        });
    },
    responsive: true,
});
