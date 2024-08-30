#!/bin/bash

echo "Please enter your Hugging Face API key:"
read HUGGINGFACE_API_TOKEN
export HUGGINGFACE_API_TOKEN

go build # run the build command

HUGGINGFACE_API_TOKEN=$HUGGINGFACE_API_TOKEN ./dashlab server &
SERVER_PID=$!

./dashlab client

echo "Successfully finished running script"

kill $SERVER_PID
