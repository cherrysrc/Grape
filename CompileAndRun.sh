#!/bin/sh
go build main.go
rm VideoName.avi
if [ $# -eq 0 ]
then
    ./main
fi

if [ $# -eq 1 ]
then
    ./main $1| ppmtoy4m -F60:1 | ffmpeg -i pipe:0 -c:v libx264rgb VideoName.avi
fi

