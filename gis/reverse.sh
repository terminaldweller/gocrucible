#!/usr/bin/sh

curl -H "Content-Type: application/json" -XPOST -d'{"lon":-58.70521,"lat":-34.4407, "len":-34.4407}' localhost:8080/reverse
