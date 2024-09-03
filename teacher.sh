#!/bin/bash

# Define the relative path to the interviews folder
INTERVIEWS_FOLDER="interviews"

# Verify that the interviews folder exists
if [[ ! -d "$INTERVIEWS_FOLDER" ]]; then
    exit 1
fi

# Locate the Interview File and Extract the Number
INTERVIEW_PATH=$(find "$INTERVIEWS_FOLDER" -type f -name "interview-*.txt" 2>/dev/null | head -n 1)

if [[ -z "$INTERVIEW_PATH" ]]; then
    exit 1
fi

# Extract the interview number from the file name
INTERVIEW_NUMBER=$(basename "$INTERVIEW_PATH" | grep -oP '\d+(?=\.txt)')

if [[ -z "$INTERVIEW_NUMBER" ]]; then
    exit 1
fi

# Export the interview number to an environment variable
export INTERVIEW_NUMBER

# Print the results
echo "$INTERVIEW_NUMBER"
cat "$INTERVIEW_PATH"
echo "$MAIN_SUSPECT"
