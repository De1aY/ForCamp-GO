$(document).ready(function() {
    var options =
    {
        thumbBox: '.fc-imageBox__thumbBox',
        spinner: '.fc-imageBox__spinner',
        imgSrc: ''
    }
    var cropper = $('.imageBox').cropbox(options);
    $('#fc-avatar__file--new').on('change', function(){
        var reader = new FileReader();
        reader.onload = function(e) {
            options.imgSrc = e.target.result;
            cropper = $('.fc-imageBox').cropbox(options);
            $('.fc-avatar').css('display', 'flex');
        }
        reader.readAsDataURL(this.files[0]);
    })
    $('.fc-avatar__save').on('click', function(){
        let rawImg = cropper.getDataURL();
        $('.fc-avatar').hide();
        let img = rawImg.replace(/^data:image\/(png|jpg);base64,/, "");
        UploadFile({url: __ChangeUserAvatar, method: "POST", params: {token: Token}, file: img,
        callback: function(resp) {
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Аватар успешно изменён", time: 2});
                $('.mdl-card__title-photo--round').attr('src', rawImg)
            } else {
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
        }});
    })
    $('.fc-avatar__crop--inc').on('click', function(){
        cropper.zoomIn();
    })
    $('.fc-avatar__crop--dec').on('click', function(){
        cropper.zoomOut();
    })
});

$('.fc-password__save').click(function(){
    let currentPasswordField = $('#fc-password_old');
    let newPasswordField = $('#fc-password_new');
    let newPasswordRepeatField = $('#fc-password_new_repeated');
    let currentPassword = currentPasswordField.val();
    let newPassword = newPasswordField.val();
    let newPasswordRepeat = newPasswordRepeatField.val();
    if (currentPassword.length < 6 || newPassword.length < 6 || newPasswordRepeat.length < 6) {
        notie.alert({type: 3, text: "Минимальная длина пароля - 6 символов", time: 2});
        return
    }
    if (newPassword !== newPasswordRepeat) {
        newPasswordField.parents('div').addClass('is-invalid');
        newPasswordRepeatField.parents('div').addClass('is-invalid');
        notie.alert({type: 3, text: "Введённые пароли не совпадают", time: 2});
        return
    }
    changePassword(currentPassword, newPassword,
    currentPasswordField, newPasswordField, newPasswordRepeatField);
});

function changePassword(oldPassword, newPassword,
    currentPasswordField, newPasswordField, newPasswordRepeatField) {
    Preloader.on();
    $.post(__ChangeUserPassword,
        {token: window.global.Token, password_new: newPassword, password_current: oldPassword},
        function(resp){
            Preloader.off();
            if(resp.code === 200) {
                notie.alert({type: 1, text: "Пароль изменён", time: 2});
                currentPasswordField.val("");
                newPasswordField.val("");
                newPasswordRepeatField.val("");
            } else {
                if (resp.code === 630) {
                    currentPasswordField.parents('div').addClass('is-invalid');
                }
                notie.alert({type: 3, text: resp.message.ru, time: 2});
            }
        });
}

$('.mdl-button--modal').click(function(){
    ShowModal($('.fc-modal'), 'flex');
});

$('.mdl-modal__button--close').click(function(){
    HideModal($('.fc-modal'));
});
