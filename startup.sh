#!/bin/bash

if [ -z "$HUGGINGFACE_API_TOKEN" ]; then
    echo "HUGGINGFACE_API_TOKEN is not set. Please set it when running the container."
    exit 1
fi

./dashlab server &
SERVER_PID=$!

./dashlab client

echo "Successfully finished running script"

kill $SERVER_PID

