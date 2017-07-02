function getMousePos(xRef, yRef) {
    let panelRect = parallax_container.getBoundingClientRect();
    return {
        x: Math.floor(xRef - panelRect.left) /
        (panelRect.right - panelRect.left) * parallax_container.offsetWidth,
        y: Math.floor(yRef - panelRect.top) /
        (panelRect.bottom - panelRect.top) * parallax_container.offsetHeight
    };

}

// Parallax
$('document').ready(function () {
    const boxer = parallax_container.querySelector(".fc-content-background");
    maxMove = parallax_container.offsetWidth / 30;
    boxerCenterX = boxer.offsetLeft + (boxer.offsetWidth / 2);
    boxerCenterY = boxer.offsetTop + (boxer.offsetHeight / 2);

    document.body.addEventListener("mousemove", function (e) {
        let mousePos = getMousePos(e.clientX, e.clientY);
        distX = mousePos.x - boxerCenterX;
        distY = mousePos.y - boxerCenterY;
        if (Math.abs(distX) < 5000 && distY < 200) {
            boxer.style.transform = "translate(" + (-1 * distX) / 12 + "px," + (-1 * distY) / 12 + "px)";
        }
    })
});

$('#submit').click(function () {
    $.get(__GetUserTokenLink, {
        login: $('#login').val(),
        password: $('#password').val()
    }, function (data) {
        if (data.code === 200) {
            $.cookie('token', data.message.token, {expires: 366, path: '/', secure: true});
            timeout = setTimeout('window.location.href = "https://forcamp.ga/general"', 2000);
            notie.alert({type: 1, text: "Вход успешно выполнен!", time: 3});
        } else {
            notie.alert({type: 3, text: data.message.ru, time: 3});
        }
    }, "json");
});