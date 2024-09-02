
json_url="https://platform.zone01.gr/assets/superhero/all.json"

superhero_info=$(curl -s "$json_url" | jq -r '.[] | select(.id == 170) | {name: .name, power: .power, gender: .gender}')

name=$(jq -r '.name' <<< "$superhero_info")
power=$(jq -r '.power' <<< "$superhero_info")
gender=$(jq -r '.gender' <<< "$superhero_info")

printf "Name: %s\nPower: %s\nGender: %s\n" "$name" "$power" "$gender"