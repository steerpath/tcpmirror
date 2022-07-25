const net = require('net');
// a function that send a simple text to a tcp socket server
function send(port, text) {
    var client = net.connect(port, '127.0.0.1', function () {
        client.write(text);
    });
}

send(8080, "Hello World");