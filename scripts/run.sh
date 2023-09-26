#!/usr/bin/bash

heading_color="\033[1;33m"
normal="\033[0m"

BASE_DIR=$(dirname "$0")

if [[ ! -z "$1" ]]; then
  if [[ ! -d "$BASE_DIR/../examples/$1" ]]; then
    echo "Example '$1' does not exist"
    exit 1
  fi

  heading="Example - ${$1^}"
  echo -e "$heading_color$heading$normal"
  go run "$BASE_DIR/../examples/$1/main.go"
  exit 0
fi

for i in "$BASE_DIR"/../examples/*;
do
  IFS="/" read -ra string <<< "$i"
  heading="Example - ${string[-1]^}"
  echo -e "$heading_color$heading$normal"
  go run "$i/main.go"
done