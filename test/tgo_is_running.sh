#!/bin/bash

mkdir -p x
cd x

wget -o is-running.err -O is-running.out 'http://127.0.0.1:10042/api/v1/status'

if grep "200 OK" is-running.err >/dev/null ; then
	echo "Success. Output:"
	cat is-running.out
	echo ""
else
	echo "Failed!"
	cat is-running.err
	echo ""
	exit 1
fi


