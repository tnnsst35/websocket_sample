#!/bin/sh
export GOPATH=$GOPATH
go get github.com/gorilla/websocket
rpm -Uvh http://dl.fedoraproject.org/pub/epel/5/i386/epel-release-5-4.noarch.rpm
yum -y install bzr
go get github.com/stretchr/gomniauth/...