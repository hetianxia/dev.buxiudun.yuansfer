#!/bin/bash

pid=`(ps -ef | grep buxiudun | grep -v "grep") | awk '{print $2}'`
if [ "$pid" != "" ];then
	kill -9 $pid
	echo "buxiudun is stopped."
else
	echo "buxiudun does not exist."
fi

nohup ./buxiudun >> nohup.out 2>&1 &