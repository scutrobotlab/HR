#!/bin/bash

export DB_HOST=172.24.64.1
export DB_PORT=5432
export DB_USER=postgres
export DB_PASS=1290
export DB_NAME=HR

TARGET_DIR="ultimate"

PROJECT_DIR=$(dirname "$0")
GENERATE_DIR="$PROJECT_DIR/cmd/$TARGET_DIR"

cd "$GENERATE_DIR" || exit

echo "Start Generating"
go run .
