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

$('.mdl-card__body-row-button--edit').click(function () {
    let editButton = $(this);
    let content = editButton.data('content');
    let acceptButton = $('.mdl-card__body-row-button--accept[data-content='+content+']');
    let declineButton = $('.mdl-card__body-row-button--decline[data-content='+content+']');
    let textField = $('.mdl-card__body-row-text[data-content='+content+']');
    textField.addClass('mdl-card__body-row-text--editing');
    textField.attr('contenteditable', 'true');
    textField.focus();
    textField.keydown(function (e) {
        if((e.keyCode) === 13) {
            textField.attr('contenteditable', 'false');
            textField.removeClass('mdl-card__body-row-text--editing');
            editButton.removeClass('mdl-card__body-row-button--off');
            acceptButton.addClass('mdl-card__body-row-button--off');
            declineButton.addClass('mdl-card__body-row-button--off');
        }
    });
    editButton.addClass('mdl-card__body-row-button--off');
    acceptButton.removeClass('mdl-card__body-row-button--off');
    declineButton.removeClass('mdl-card__body-row-button--off');
});