#!/bin/bash

mkdir -p reports
echo "mode: set" > reports/coverage.out

for dir in $(find . -type d ! -path "./cmd/api*"); do
    if ls "$dir"/*_test.go &> /dev/null; then
        go test -coverprofile=tmp_coverage.out "$dir" && tail -n +2 tmp_coverage.out >> reports/coverage.out
        rm -f tmp_coverage.out
    fi
done