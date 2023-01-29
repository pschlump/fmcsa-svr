#!/bin/bash

mkdir -p x
cd x

MC="$1"
if [ -z "${MC}" ] ; then
	echo "Missing MC number as parameter"
	exit 1
fi

if [ -z "${FMCSA_WebKey}" ] ; then
	echo "Must source config file to setup FMCSA key"
	exit 1
fi

TOKEN=FMCSA.l4BUEAzcBvqu8_C1JFxi8vMFr5g

wget -o mc-out.err -O mc-out.out \
	--header="X-Authentication: ${TOKEN}" \
	"http://127.0.0.1:10042/api/v1/mc-number-data?mc=${MC}"

cat mc-out.err
echo ""
cat mc-out.out
echo ""

