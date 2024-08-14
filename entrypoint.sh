#!/bin/bash

/opt/chainbase-cli run --config /opt/avs.toml

if [ $? -eq 0 ]; then
    echo "Signature verification passed. Starting Flink..."
    # Start Flink
    /opt/flink/bin/start-cluster.sh
    
    # Keep the container running
    tail -f /dev/null
else
    echo "Signature verification failed. Exiting..."
    exit 1
fi