#!/bin/bash

go build -o ~/go/src/adventofcode/2020/bin/$1 days/$1/*

if [ $? -eq 0 ]
then
    ~/go/src/adventofcode/2020/bin/$1;
fi