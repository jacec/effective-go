#!/bin/sh -e

#check that godep is installed, if not, do it!
command -v $GOPATH/bin/godep >/dev/null 2>&1 || {
  echo "godep not found, so installing.." >&2;
  go get github.com/tools/godep
}
