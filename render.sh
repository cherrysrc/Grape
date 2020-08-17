#!/bin/sh
./main | ffmpeg -i pipe:0 -c:v libx264rgb VideoName.avi
