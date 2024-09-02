
json_url="https://platform.zone01.gr/assets/superhero/all.json"

superhero_info=$(curl -s "$json_url" | jq -r '.[] | select(.id == 170) | {name: .name, power: .power, gender: .gender}')

name=$(echo "$superhero_info" | jq -r '.name')
power=$(echo "$superhero_info" | jq -r '.power')
gender=$(echo "$superhero_info" | jq -r '.gender')

printf "Name: %s\nPower: %s\nGender: %s\n" "$name" "$power" "$gender"