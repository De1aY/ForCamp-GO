let preloader = new $.materialPreloader({
    position: 'top',
    height: '5px',
    col_1: '#159756',
    col_2: '#da4733',
    col_3: '#3b78e7',
    col_4: '#fdba2c',
    fadeIn: 200,
    fadeOut: 200
});

$('#submit').click(function () {
    preloader.on();
    $.get(__GetUserTokenLink, {
        login: $('#login').val(),
        password: $('#password').val()
    }, function (data) {
        preloader.off();
        if (data.code === 200) {
            $.cookie('token', data.message.token.replace('%3D', '='), {expires: 366, path: '/', secure: true});
            window.location.href = "https://forcamp.nullteam.info";
        } else {
            notie.alert({type: 3, text: data.message.ru, time: 3});
        }
    }, "json");
});

$('.fc-auth__form-field').keydown(function (e) {
   if(e.keyCode === 13) {
       $('#submit').click();
   }
});

$('.fc-auth__overlay').click(function () {
    $('.fc-auth').animate({
        opacity: 0
    }, 700, function () {
        $(this).css('display', 'none');
    });
});

$('#fc-header__action-button--auth').click(function () {
    let authForm = $('.fc-auth');
    authForm.css('display', 'flex');
    authForm.animate({
       opacity: 1
    }, 700);
});

$(document).ready(function() {
    /*$('#fullpage').fullpage({
        //Navigation
        menu: '#menu',
        lockAnchors: false,
        anchors:['firstPage', 'secondPage'],
        navigation: false,
        navigationPosition: 'right',
        navigationTooltips: ['firstSlide', 'secondSlide'],
        showActiveTooltip: false,
        slidesNavigation: false,
        slidesNavPosition: 'bottom',

        //Scrolling
        css3: true,
        scrollingSpeed: 700,
        autoScrolling: true,
        fitToSection: true,
        fitToSectionDelay: 1000,
        scrollBar: false,
        easing: 'easeInOutCubic',
        easingcss3: 'ease',
        loopBottom: false,
        loopTop: false,
        loopHorizontal: true,
        continuousVertical: false,
        continuousHorizontal: false,
        scrollHorizontally: false,
        interlockedSlides: false,
        dragAndMove: false,
        offsetSections: false,
        resetSliders: false,
        fadingEffect: false,
        scrollOverflow: false,
        scrollOverflowReset: false,
        scrollOverflowOptions: null,
        touchSensitivity: 15,
        normalScrollElementTouchThreshold: 5,
        bigSectionsDestination: null,

        //Accessibility
        keyboardScrolling: true,
        animateAnchor: true,
        recordHistory: true,

        //Design
        controlArrows: true,
        verticalCentered: true,
        paddingTop: '0',
        paddingBottom: '10px',
        fixedElements: '#header, .footer',
        responsiveWidth: 0,
        responsiveHeight: 0,
        responsiveSlides: false,
        parallax: false,
        parallaxOptions: {type: 'reveal', percentage: 62, property: 'translate'},

        //Custom selectors
        sectionSelector: '.fc-content__section',
        slideSelector: '.fc-content__section-slide',

        lazyLoading: true,

        //Events
        onLeave: function(index, nextIndex, direction){},
        afterLoad: function(anchorLink, index){},
        afterRender: function(){},
        afterResize: function(){},
        afterResponsive: function(isResponsive){},
        afterSlideLoad: function(anchorLink, index, slideAnchor, slideIndex){},
        onSlideLeave: function(anchorLink, index, slideIndex, direction, nextSlideIndex){}
    });*/
});