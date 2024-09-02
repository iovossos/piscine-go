#!/bin/bash

# Find all .sh files, remove the .sh extension, sort them in descending order, and print the result
find . -type f -name "*.sh" | \
    sed 's/\.sh$//' | \
    awk '!seen[$0]++' | \
    sort -r
