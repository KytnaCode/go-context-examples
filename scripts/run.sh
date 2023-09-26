#!/usr/bin/bash

heading_color="\033[1;33m"
normal="\033[0m"

BASE_DIR=$(dirname "$0")

get_help() {
  echo "Usage: $0 <options> <example_name>"
  echo "example_name:"
  echo "  - name of the example to run"
  echo "  - run '$0 -l' to list available examples"
  echo "  - if not provided, all examples will be run"
  echo "Options:"
  echo "  -l: list available examples"
  echo "  -h: show this help"
  exit 0
}

print_examples() {
  echo "Available examples:"
  for i in "$BASE_DIR"/../examples/*;
  do
    IFS="/" read -ra string <<< "$i"
    echo "  - ${string[-1]}"
  done
}

execute_example() {
  folder_name="${1,,}"

  if [[ ! -d "$BASE_DIR/../examples/$folder_name" ]]; then
    echo "Example '$1' does not exist"
    exit 1
  fi

  IFS=- read -ra example_name <<< "$1"

  heading="Example - ${example_name[*]^}"
  echo -e "$heading_color$heading$normal"
  go run "$BASE_DIR/../examples/$1/main.go"
}

execute_all_examples() {
  for i in "$BASE_DIR"/../examples/*;
  do
    IFS="/" read -ra string <<< "$i"
    heading="Example - ${string[-1]^}"
    echo -e "$heading_color$heading$normal"
    go run "$i/main.go"
  done
}

if [[ "$1" == "-h" ]]; then
  get_help
elif [[ "$1" == "-l" ]]; then
  print_examples
elif [[ ! -z "$1" ]]; then
  execute_example "$1"
elif [[ -z "$1" ]]; then
  execute_all_examples
fi