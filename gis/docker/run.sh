#!/usr/bin/env sh
docker run -p 8080:8080 -it --entrypoint /root/gis $1
