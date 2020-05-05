#!/usr/bin/env bash

# Usage:
# ./scripts/longpoll-wrapper/generator.sh <input-file> <output-file>

if [[ -n "$1"  && -n "$2" ]]; then
cat wrapper_header.text > $2
cat $1 | grep "//" | grep "struct" | sed 's/\/\/ / /' | awk -f generate_wrapper.awk >> $2
else
echo "Usage:"
echo "./scripts/longpoll-wrapper/generator.sh <input-file> <output-file>"
fi