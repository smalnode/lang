#!/usr/bin/env bash

# shell cmd parameters list $* $@
# a b c -> $* -> "a b c" 
# a b c -> "$@" -> "a" "b" "c"
echo ${1} ${2} ${3} 

# op := exec when var is not set or is empty
echo ${NOSET_OR_EMPTY:=default}
# op = exec when var is not set
echo ${NOSET_OR_EMPTY=defaultb}
# op :- do not evaluate var, but set the expr to default
echo ${4:-default}
echo ${4} 

# ${var:?"ErrMsg"} <==> [ -z $var]
if [ -z $5 ]
then
    echo "5th param is not set or empty"
fi
# op :? check wheather var is not set or empty, exit if so
NOT_SET=${5:?"5th param is not set or empty"}

echo "Cmd exit staus:" $?

echo "Number of params:" $#

for PARAM in $*
do
    printf "%s\n" $PARAM
done

printf "\n\n"

for PARAM in "$@"
do
    printf "%s\n" "$PARAM"
done

for PARAM in *.sh
do
    # elimate sh and apend shell
    echo "${PARAM%sh}shell"
    # same as subsitution
    echo "${PARAM/.sh/.shell}"
done
