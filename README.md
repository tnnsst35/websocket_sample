# websocket_sample

## Prepare
Operation on the host machine.

|Confirmed|Unconfirmed|
|---------|-----------|
|Windows7 |MacOS      |
|         |Linux      |

1. $ git clone https://github.com/tnnsst35/websocket_sample.git
2. $ cd websocket_sample
3. $ sh build.sh
4. $ sh run.sh

## golang
Operation on the guest machine.
Docker container.

1. \# cd /workspace/go/src/chat
2. \# sh init.sh
3. \# go build chat ./
4. \# ./chat
5. Please open with web browser. http://localhost:8080

## Node.js
Operation on the guest machine.
Docker container.

1. \# cd /workspace/nodejs/src/chat
2. \# sh init.sh
3. \# node app.js
4. Please open with web browser. http://localhost:8080