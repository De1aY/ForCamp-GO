let SelfData = {
    Login: "",
    Data: ""
};
let UserData = {
    Login: "",
    Data: ""
};
showProfile_Status = false;
getSelfLogin();

function getSelfLogin() {
    $.get(__GetUserLoginLink, {token: Token}, function (resp) {
        if (resp.code === 200) {
            SelfData.Login = resp.message.login;
            getSelfData();
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 3});
        }
    }, 'json');
}

function getSelfData() {
    $.get(__GetUserDataLink, {token: Token, login: SelfData.Login}, function (resp) {
        if (resp.code === 200) {
            SelfData.Data = resp.message.data;
        } else {
            notie.alert({type: 3, text: resp.message.ru, time: 3});
        }
    }, 'json');
}

function setProfile(userData) {
    setFullName(userData.Data.name, userData.Data.surname, userData.Data.middlename);
    setAvatar(userData.Data.avatar);
}

function setFullName(name, surname, middlename) {
    fullName = surname + ' ' + name[0] + '.' + middlename[0] + '.';
    $('.fc-profile__header-fullname').text(fullName);
}

function setAvatar(src) {
    $('.fc-profile__header-avatar').attr('src', __ImagesLink + src);
}

function showProfile(userData) {
    setProfile(userData);
    $('.fc-profile').animate({
        left: "+=353px"
    }, 1000);
}

function hideProfile() {
    $('.fc-profile').animate({
        left: "-=353px"
    }, 1000);
}

function toggleProfile(userData, button) {
    if(showProfile_Status) {
        $(button).removeClass('fc-header__menu-section--active');
        hideProfile();
    } else {
        $(button).addClass('fc-header__menu-section--active');
        showProfile(userData);
    }
    showProfile_Status = !showProfile_Status;
}

$(document).ready(function () {
    $('.fc-header__menu .fc-header__menu-section:first-child').click(function () {
        toggleProfile(SelfData, this);
    });
});