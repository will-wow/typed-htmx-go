#!/usr/bin/env bash
#
# Exit if tests fail
set -e

# Run tests and generate coverage report
go test -coverpkg ./... -coverprofile=coverage/coverage.out ./...

# Grep for the total value
COVERAGE=`go tool cover -func=coverage/coverage.out | grep 'total:' | grep -Eo '[0-9]+\.[0-9]+'`


# Set the color based on the coverage percentage
COLOR=orange
if (( $(echo "$COVERAGE <= 50" | bc -l) )) ; then
    COLOR=red
    elif (( $(echo "$COVERAGE > 80" | bc -l) )); then
    COLOR=green
fi

echo "Coverage is $COVERAGE%"

# Generate a badge to display in the README.
curl "https://img.shields.io/badge/coverage-$COVERAGE%25-$COLOR" > ./assets/badge.svg
