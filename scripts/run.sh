#!/bin/bash

go build -o ~/adventofcode/2020/bin/$1 days/$1/*

if [ $? -eq 0 ]
then
    ~/adventofcode/2020/bin/$1;
fi
