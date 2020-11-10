#!/usr/bin/sh

%d9%85%d8%b1%d8%b2%d8%af%d8%a7%d8%b1%d8%a7%d9%86
curl -XGET -d'{"autocomp":"%d8%b3%d8%a7%d8%b2%d9%85%d8%a7%d9%86%20%d8%a7%d8%a8"}' localhost:8080/autocomp
curl -XGET -d'{"autocomp":"%d9%82%d8%a7%d8%a6%d9%85"}' localhost:8080/autocomp
curl -XGET -d'{"autocomp":"%d9%85%d8%b1%d8%b2%d8%af%d8%a7%d8%b1%d8%a7%d9%86"}' localhost:8080/autocomp
