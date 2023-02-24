#!/bin/bash

if [ $# -ne 4 ]; then
  echo "Usage: $0 <namespace> <verb> <resource> <api-group>"
  exit 1
fi

NAMESPACE=$1
VERB=$2
RESOURCE=$3
API_GROUP=$4

for EACH in $(kubectl -n "$NAMESPACE" get clusterroles,roles --no-headers | awk '{ print $1 }');
do
  k get "$EACH" -o yaml | python3 rulematching.py "$API_GROUP" "$RESOURCE" "$VERB"
done