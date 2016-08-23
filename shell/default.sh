#!/usr/bin/env bash

# shell cmd default parameter operator :-
echo ${1:-"default"}

# default env variable value
echo ${USER:="wdeqin"}
echo ${HOME:="/home"}
