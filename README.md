# websocket_sample

## Prepare
Operation on the host machine.

|Confirmed|Unconfirmed|
|---------|-----------|
|Windows7 |MacOS      |
|         |Linux      |

1. $ git clone https://github.com/tnnsst35/websocket_sample.git
2. $ cd websocket_sample
3. $ sh run.sh

## golang
Operation on the guest machine.
Docker container.

4. \# cd /workspace/go/src/chat
5. \# sh init.sh
6. \# go build chat ./
7. \# ./chat
9. Please open with web browser. http://localhost:8080