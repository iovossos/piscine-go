
# Check if HERO_ID environment variable is set
if [ -z "$HERO_ID" ]; then
    echo "Error: HERO_ID environment variable is not set."
    exit 1
fi

# Ensure HERO_ID is treated as a number
# Convert HERO_ID to an integer to ensure correct type
HERO_ID_INT=$(printf "%d" "$HERO_ID")

# Fetch the JSON data
json_data=$(curl -s https://platform.zone01.gr/assets/superhero/all.json)

# Check if curl succeeded
if [ $? -ne 0 ]; then
    echo "Error: Failed to fetch data."
    exit 1
fi

# Extract the family information from the connections tab for the given HERO_ID
family_info=$(echo "$json_data" | jq -r --argjson id "$HERO_ID_INT" '
  .[] | select(.id == $id) | .connections.relatives
')

# Check if the HERO_ID was found and has relatives data
if [ -z "$family_info" ]; then
    echo "Error: No data found for HERO_ID '$HERO_ID'."
    exit 1
fi

# Display the family information
printf "\"%s\"\n" "$family_info"
