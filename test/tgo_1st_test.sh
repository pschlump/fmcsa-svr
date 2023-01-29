#!/bin/bash

# Copyright (c) Philip Schlump, 2023.
# This file is MIT licensed, see ../LICENSE.mit

curl  \
	--header "X-Authentication: FMCSA.l4BUEAzcBvqu8_C1JFxi8vMFr5g" \
	 'http://127.0.0.1:10042/api/v1/mc-number-data?mc=MC-53467'

