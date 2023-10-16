#!/bin/bash

# Read the env.sample file
while IFS= read -r line; do
  # Extract the secret name from the placeholder, assuming placeholders are like {{SECRET_NAME}}
  secret_name="${line#*{{}"
  secret_name="${secret_name%%}}*}"

  if [[ -n "$secret_name" ]]; then
    # Replace the placeholder with the corresponding GitHub secret
    secret_value="${{ secrets.$secret_name }}"
    line="${line/\{\{${secret_name}\}\}/$secret_value}"
  fi

  echo "$line"
done < .env.sample > .env
