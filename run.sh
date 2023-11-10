#!/bin/sh

env=""

case $(uname -a) in
    *Microsoft*|*microsoft*) env="GOOS=windows";;
esac

sh -c "$env go run ."