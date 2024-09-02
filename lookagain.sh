#!/bin/bash

# Find all .sh files, remove the .sh extension and the leading ./, sort them in descending order, and print the result
find . -type f -name "*.sh" | \
    sed -e 's/^\.\///' -e 's/\.sh$//' | \
    sort -r
