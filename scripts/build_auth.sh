#!/bin/bash
echo "Build AUTHGO"
go build -ldflags "-s" $GOPATH/src/github.com/pascallimeux/auth-go/src/main.go
if [ ! -d "$GOPATH/src/github.com/pascallimeux/ocms/dist" ]; then
  mkdir $GOPATH/src/github.com/pascallimeux/auth-go/dist
fi
mv main $GOPATH/src/github.com/pascallimeux/auth-go/dist/auth-go
cp $GOPATH/src/github.com/pascallimeux/auth-go/config/config.json $GOPATH/src/github.com/pascallimeux/auth-go/dist/config.json
cp *.sh $GOPATH/src/github.com/pascallimeux/auth-go/dist
chmod u+x $GOPATH/src/github.com/pascallimeux/auth-go/dist/*.sh


