#!/bin/sh
go build main.go
rm Video.avi
./main $1| ppmtoy4m -F60:1 | ffmpeg -i pipe:0 -c:v libx264rgb Video.mp4
