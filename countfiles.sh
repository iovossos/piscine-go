#!/bin/bash
c
# Count the number of regular files and directories in the current directory and its subdirectories
find . -type f -o -type d | wc -l
