#!/usr/bin/sh

docker run -p 8080:8080 -it --entrypoint gis $1
