#!/usr/bin/env bash

if [[ $# -ne 1 ]]; then
    echo usage: "create.sh <file name>"
    exit 1
fi

name=$1

# AldS15C -> alds15c
name=$(echo $name | gsed -E "s/.*/\L\0/g")

cp ./_tmp_solve.go "${name}.go"
cp ./_tmp_solve_test.go "${name}_test.go"

# alds15c -> Alds15c
name=$(echo $name | gsed -E "s/^([a-z])/\U\1/g")

#v="ho ge"; echo ${v^}→Ho geo
gsed -i -e "s/AldsTmp/${name}/g" "${1}.go"
gsed -i -e "s/TestTmp/Test${name}/g" "${1}_test.go"
