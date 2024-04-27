#!/bin/sh
param=$1 || 0
curl -X GET "http://localhost:8080/fibonacci?n=$param" | jq
