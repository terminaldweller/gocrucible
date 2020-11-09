#!/usr/bin/sh

curl -H "Content-Type: application/json" -XGET -d'{"lon":51.52103,"lat":35.7988357, "len":-34.4407}' localhost:8080/reverse
