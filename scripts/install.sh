#!/bin/bash
currentpath=`pwd`
dest="/usr/local/pipe"

mkdir -p ${dest}
mkdir -p ${dest}/conf
mkdir -p ${dest}/bin
cp -f ${currentpath}/../bin/pipe-server -p ${dest}/bin/
cp -f ${currentpath}/server/config.yaml -p ${dest}/conf/