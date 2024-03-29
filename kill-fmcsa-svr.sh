#!/bin/bash

if [ -f log/pid_file ] ; then
	if ps -ef | awk '{print $2}' | grep $( cat ./log/pid_file ) >/dev/null ; then
		kill " $( cat ./log/pid_file ) "
	fi
fi

if ps -ef | grep -v grep | grep -v kill- | grep fmcsa-svr >/dev/null ; then
	X=$(  ps -ef |  grep -v grep | grep -v kill- | grep fmcsa-svr | awk '{print $2}' )
	echo "->$X<-"
	if [ -z "$X" ] ; then
		:
	else 
		kill $X
	fi
fi

