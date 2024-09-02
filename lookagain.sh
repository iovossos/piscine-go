#!/bin/bash

# Find all .sh files, remove the directory path, remove the .sh extension, sort them in descending order, and print the result
find . -type f -name "*.sh" | \
    xargs -n 1 basename | \
    sed 's/\.sh$//' | \
    sort -r
