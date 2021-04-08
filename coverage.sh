#!/bin/bash

go test ./... -coverprofile cover.out
curr_coverage=$(go tool cover -func cover.out | grep total | awk '{print $3}' | tr -d '%')
trgt_coverage="80"

echo "Total: $curr_coverage%"

if [ "$(echo $curr_coverage | cut -d'.' -f1)" -ge "$trgt_coverage" ]; then
  echo "Unit tests coverage is OK!"
  exit 0
else
  echo "Unit tests do not pass $trgt_coverage% coverage"
  exit 1
fi