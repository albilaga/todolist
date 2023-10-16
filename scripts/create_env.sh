#!/bin/bash

# Function to get the value from GitHub secrets or use the default
get_secret_value() {
  local secret_name="$1"
  local default_value="$2"

  if [[ -n "${!secret_name}" ]]; then
    echo "${!secret_name}"
  else
    echo "$default_value"
  fi
}

# Read .env.sample and create .env
while IFS= read -r line; do
  # Check if the line contains a key-value pair
  if [[ "$line" =~ ^[A-Za-z_][A-Za-z_0-9]*= ]]; then
    # Extract key and value
    key="${line%%=*}"
    value="${line#*=}"

    secret_value=$(get_secret_value "key" "$value")

    # Write to .env
    echo "$key=$secret_value"
  else
    # Pass-through non-key-value lines
    echo "$line"
  fi
done < .env.sample > build/.env
