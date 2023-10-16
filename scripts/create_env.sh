#!/bin/bash

# Read each line in .env.sample file
while IFS= read -r line
do
  # Check if line isn't empty or a comment
  if [[ ! -z "$line" && ! "$line" =~ ^# ]]; then
    # remove any existing instances of the variable in .env
    sed -i "/^$line/d" ../.env
    # append variable to .env
    echo "$line" >> ../.env
  fi
done < "../.env.sample"