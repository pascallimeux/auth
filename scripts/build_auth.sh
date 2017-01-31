#!/bin/bash
echo "Build AUTHGO"
go build -ldflags "-s" $GOPATH/src/github.com/pascallimeux/auth/auth.go
if [ ! -d "$GOPATH/src/github.com/pascallimeux/auth/dist" ]; then
  mkdir $GOPATH/src/github.com/pascallimeux/auth/dist
fi
mv auth $GOPATH/src/github.com/pascallimeux/auth/dist/auth
cp $GOPATH/src/github.com/pascallimeux/auth/settings.toml $GOPATH/src/github.com/pascallimeux/auth/dist/settings.toml
cp *.sh $GOPATH/src/github.com/pascallimeux/auth/dist
chmod u+x $GOPATH/src/github.com/pascallimeux/auth/dist/*.sh


