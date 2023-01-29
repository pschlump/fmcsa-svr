#!/bin/bash

# Copyright (c) Philip Schlump, 2023.
# This file is MIT licensed, see ../LICENSE.mit

mkdir -p x
cd x

curl -o tgo_1st.out \
	--header "X-Authentication: FMCSA.l4BUEAzcBvqu8_C1JFxi8vMFr5g" \
	 'http://127.0.0.1:10042/api/v1/mc-number-data?mc=MC-53467'

cat tgo_1st.out
echo ""

cp tgo_1st.out ../ref/tgo_1st.ref 

if diff tgo_1st.out ../ref/tgo_1st.ref ; then
	:
else
	echo "Failed" | color-cat -c red
	exit 1
fi 

echo "PASS" | color-cat -c green

