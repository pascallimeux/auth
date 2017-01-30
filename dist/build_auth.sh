#!/bin/bash
echo "Build AUTHGO"
go build -ldflags "-s" $GOPATH/src/github.com/pascallimeux/auth/src/main.go
if [ ! -d "$GOPATH/src/github.com/pascallimeux/ocms/dist" ]; then
  mkdir $GOPATH/src/github.com/pascallimeux/auth/dist
fi
mv main $GOPATH/src/github.com/pascallimeux/auth/dist/auth-go
cp $GOPATH/src/github.com/pascallimeux/auth/config/config.json $GOPATH/src/github.com/pascallimeux/auth/dist/config.json
cp *.sh $GOPATH/src/github.com/pascallimeux/auth/dist
chmod u+x $GOPATH/src/github.com/pascallimeux/auth/dist/*.sh


