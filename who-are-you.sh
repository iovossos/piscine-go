json_url="https://platform.zone01.gr/assets/superhero/all.json"

superhero_info=$(curl -s "$json_url" | jq -r '.[] | select(.id == 70) | {name: .name}')

name=$(jq -r '.name' <<< "$superhero_info")

printf "%s" "$name"