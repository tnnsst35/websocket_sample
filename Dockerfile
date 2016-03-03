# 基本イメージ
FROM centos:7

# 管理者
MAINTAINER tnnsst35

# メタデータ
LABEL version="0.0.1"

# おまじない
ENV LANG=ja_JP.utf8

RUN rm -f /etc/localtime

RUN ln -fs /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

RUN yum -y update

RUN yum -y install wget vim git

# Go言語のインストール
RUN cd /usr/local/src && wget https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf /usr/local/src/go1.6.linux-amd64.tar.gz

ENV GO_HOME=/usr/local/go

ENV PATH=$PATH:$GO_HOME/bin

RUN echo 'export GOPATH=$HOME/go:/workspace/go' >> /etc/profile.d/go.sh

RUN mkdir /etc/skel/go

RUN mkdir $HOME/go

# Node.jsのインストール
RUN git clone https://github.com/creationix/nvm.git /usr/local/nvm

RUN echo 'export NVM_HOME=/usr/local/nvm' >> /etc/profile.d/nvm.sh

RUN echo '. $NVM_HOME/nvm.sh' >> /etc/profile.d/nvm.sh

RUN . /etc/profile.d/nvm.sh && nvm install v4.3.2 && nvm use v4.3.2

# Nginx+PHPのインストール

# 開放ポート
EXPOSE 8080