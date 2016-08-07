#!/usr/bin/env bash
for SCRIPT in /d/ws/lang/shell/*
do 
    if [ -f $SCRIPT -a $SCRIPT ]
    then
        $SCRIPT
    fi
done
