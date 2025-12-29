#!/bin/sh
set -eu

TARGET_APPLICATION=${1:-api}

if [ ! -f "/app/${TARGET_APPLICATION}" ]; then
    echo "Error: /app/${TARGET_APPLICATION} not found"
    exit 1
fi

exec "/app/${TARGET_APPLICATION}"
