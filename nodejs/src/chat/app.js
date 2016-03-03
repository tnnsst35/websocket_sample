var server = require('http').createServer();
server.listen(8080);

// http://host:8080のアクセスはHTMLを返す
server.on('request', function (req, res) {
  require('fs').readFile('./index.html', 'utf-8', function (error, data) {
    res.write(data);
    res.end();
  });
});

var io = require('socket.io').listen(server);

// Socket.IOの処理
io.on('connection', function (socket) {     // 接続されるとconnectionイベントが発生する
  socket.on('hello_everyone', function (msg) {  // hello_everyoneイベントを定義する
    socket.broadcast.emit('message', msg);  // messageイベントを発生させる
  });
  socket.on('disconnect', function () {     // 切断されるとdisconnectイベントが発生する
    console.log('byebye');
  });
});