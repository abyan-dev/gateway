#!/bin/bash

mkdir -p reports
echo "mode: set" > reports/coverage.out

for dir in $(find . -type d ! -path "./cmd/api*"); do
    ls "$dir"/*_test.go &> /dev/null && go test -coverprofile=tmp_coverage.out "$dir" && tail -n +2 tmp_coverage.out >> reports/coverage.out && rm tmp_coverage.out
done