#!/bin/sh
go build main.go
rm VideoName.avi
./main | ppmtoy4m -F60:1 | ffmpeg -i pipe:0 -c:v libx264rgb VideoName.avi

