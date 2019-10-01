window.onload = function() {
    // Only use connect once because its like different connections
    const templateAudioContext = window.AudioContext || window.webkitAudioContext;
    var audioContext = new templateAudioContext();
    var analyser = audioContext.createAnalyser();

    const audioElement = document.querySelector('audio');
    audioElement.play().then(response =>{
        console.log(response);
    }).catch(e => {
        console.log(e);
    });

    const track = audioContext.createMediaElementSource(audioElement);

    // select our play button
    const playButton = document.querySelector('button');

    playButton.addEventListener('click', function() {
        // check if context is in suspended state (autoplay policy)
        if (audioContext.state === 'suspended') {
            audioContext.resume();
        }

        // play or pause track depending on state
        if (this.dataset.playing === 'false') {
            audioElement.play();
            this.dataset.playing = 'true';
        } else if (this.dataset.playing === 'true') {
            audioElement.pause();
            this.dataset.playing = 'false';
        }

    }, false);

    // ended audio listener
    audioElement.addEventListener('ended', () => {
        playButton.dataset.playing = 'false';
    }, false);

    // add volume control 
    const gainNode = audioContext.createGain();
    const volumeControl = document.querySelector('#volume');
    volumeControl.addEventListener('input', function() {
        gainNode.gain.value = this.value;
    }, false);

    // panning 
    const pannerOptions = { pan: 0 };
    const panner = new StereoPannerNode(audioContext, pannerOptions);
    const pannerControl = document.querySelector('#panner');
    pannerControl.addEventListener('input', function() {
        panner.pan.value = this.value;
    }, false);

    // analyzer code 
    analyser.fftSize = 2048;
    var bufferLength = analyser.frequencyBinCount;
    var dataArray = new Uint8Array(bufferLength);
    analyser.getByteTimeDomainData(dataArray);
    analyser.minDecibels = -120;
    analyser.maxDecibels = 10;
}