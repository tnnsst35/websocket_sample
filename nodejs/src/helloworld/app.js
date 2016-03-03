// httpモジュールを読み込む。
var http = require('http');

// createServerメソッドで、サーバーオブジェクトを生成する。
var server = http.createServer();

// requestイベントを定義する。
server.on('request', function (req, res) {
  res.write('Hello World');
  res.end();
});

// 8080ポートで待ち受け開始。
server.listen(8080);