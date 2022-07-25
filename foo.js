// create a simple tcp socket server
var net = require('net');

function createServer(port) {
    var server = net.createServer(function (socket) {
        socket.on('data', function (data) {
            console.log(port, "==>", data.toString());
        });
    }
    ).listen(port)
    console.log('Server listening on port ' + port);

}


createServer(9001)
createServer(9002)