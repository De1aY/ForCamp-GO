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