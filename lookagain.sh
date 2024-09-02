#!/bin/bash

# Find all .sh files, extract just the file name without any directory, remove the .sh extension, and sort them in descending order
find . -type f -name "*.sh" | \
    sed 's/^.*\///' | \
    sed 's/\.sh$//' | \
    sort -r
