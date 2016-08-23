#!/usr/bin/env bash
MYSELF=run_myself.sh
if [ -f $MYSELF -a $MYSELF ]
then
    printf "running %s ...\n" $MYSELF 
    $MYSELF
fi
