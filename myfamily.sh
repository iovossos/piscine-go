#!/bin/bash

# Check if HERO_ID environment variable is set
if [ -z "$HERO_ID" ]; then
    echo "Error: HERO_ID environment variable is not set."
    exit 1
fi

# Ensure HERO_ID is treated as a number
HERO_ID_INT=$(printf "%d" "$HERO_ID")

# Fetch the JSON data
json_data=$(curl -s https://platform.zone01.gr/assets/superhero/all.json)

# Check if curl succeeded
if [ $? -ne 0 ]; then
    echo "Error: Failed to fetch data."
    exit 1
fi

# Convert JSON data to a single-line string
json_data_single_line=$(echo "$json_data" | tr -d '\n')

# Extract the family information from the connections tab for the given HERO_ID
family_info=$(echo "$json_data_single_line" | jq -r --argjson id "$HERO_ID_INT" '
  .[] | select(.id == $id) | .connections.relatives
')

# Check if the HERO_ID was found and has relatives data
if [ -z "$family_info" ]; then
    echo "Error: No data found for HERO_ID '$HERO_ID'."
    exit 1
fi

# Remove any leading and trailing quotes from the output
family_info_clean=$(echo "$family_info" | sed 's/^"\(.*\)"$/\1/')

# Output the data without extra quotes, preserving formatting
printf "%s\n" "$family_info_clean"
