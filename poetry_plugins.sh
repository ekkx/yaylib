#!/bin/bash

plugins=(
  "poetry-version-plugin"
)

for plugin in "${plugins[@]}"; do
  poetry self add "$plugin"
done
