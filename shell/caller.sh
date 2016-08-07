#!/usr/bin/env bash
export GREETING="HELLO, WORLD!"
./callee.sh

echo "IN CALLER, GREETING = " $GREETING
printf "\n"
