#!/bin/bash

clear

if [ $# -gt 0 ] && { [ "$1" = "rebuild" ] || [ "$1" = "clean" ]; }; then
    set -v

    nix-shell --pure --run "go clean"

    set +v
fi

set -v

nix-shell --pure --run "go build -o ./out/echo-go cmd/main.go"
