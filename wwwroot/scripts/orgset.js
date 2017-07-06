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

function submitSettingEdit(textField, editButton, acceptButton, declineButton, content) {
    textField.unbind('keydown');
    textField.attr('contenteditable', 'false');
    textField.removeClass('mdl-card__body-row-text--editing');
    editButton.removeClass('mdl-card__body-row-button--off');
    acceptButton.addClass('mdl-card__body-row-button--off');
    declineButton.addClass('mdl-card__body-row-button--off');
    OrgSettings[content] = textField.text();
    editOrganizationSettings();
}

function declineSettingEdit(textField, editButton, acceptButton, declineButton, baseText) {
    textField.unbind('keydown');
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
            submitSettingEdit(textField, editButton, acceptButton, declineButton, content);
        }
    });
    acceptButton.click(function () {
        submitSettingEdit(textField, editButton, acceptButton, declineButton, content);
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