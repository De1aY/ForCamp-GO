function ActivateWavesEffect(){
    Waves.attach('.wave-effect');
    Waves.attach('.wave-effect-circle', ['waves-circle', 'waves-float']);
    Waves.init();
}

$(document).ready(function(){
    ActivateWavesEffect();
});