#!/bin/sh

# Detect system architecture
ARCH=$(uname -m)

# Select the appropriate binary
if [ "$ARCH" = "x86_64" ]; then
    EXEC="./main_amd64"
elif [ "$ARCH" = "aarch64" ]; then
    EXEC="./main_arm64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

# Run the selected binary with all passed arguments
exec $EXEC "$@"
