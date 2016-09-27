var WebSocket = require('faye-websocket'),
    ws = new WebSocket.Client('ws://localhost:1323/ws?user_id=1');

ws.on('open', function (event) {
    console.log('open');
    setInterval(function () {
        ws.send(JSON.stringify({
            user_id: '1',
            content: '123123123'
        }));
    }, 1000);
});

ws.on('message', function (event) {
    console.log('message', event.data);
});

var timer = setInterval(function () {
    ws.ping();
}, 10000);

timer.unref();

ws.on('close', function (event) {
    console.log('close', event.code, event.reason);
    clearInterval(timer);
    ws = null;
});
