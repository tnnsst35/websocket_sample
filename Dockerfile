###
# このDockerfileを適当なディレクトリに置きます
# 置いたディレクトリで以下のコマンドでビルドし、起動します。
# docker build -t tnnsst35/websocket-sample .
# docker images
# docker run -i -t -v /c/Users/ユーザー名/Workspace:/workspace tnnsst35/websocket-sample /bin/bash
# 起動しているDockerプロセスの確認、破棄は以下のコマンドで。
# docker ps -a
# docker rm `docker ps -a -q`
###

# 基本イメージ
FROM centos:7

# 管理者
MAINTAINER tnnsst35

# メタデータ
LABEL version="0.0.1"

# おまじない
RUN yum -y update

RUN yum -y install wget vim git

# Go言語のインストール
RUN cd /usr/local/src && wget https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf /usr/local/src/go1.6.linux-amd64.tar.gz

ENV GO_HOME=/usr/local/go

ENV PATH=$PATH:$GO_HOME/bin

RUN echo 'GOPATH=$HOME/go' >> /etc/profile.d/go.sh

RUN mkdir /etc/skel/go

RUN mkdir $HOME/go

# Nginx+PHPのインストール

# Node.jsのインストール

# 開放ポート
# EXPOSE 80 8000 8080 8800