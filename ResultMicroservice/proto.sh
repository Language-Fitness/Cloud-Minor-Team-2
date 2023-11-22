#!/bin/bash

# Set your GitHub repository details
githubRepo="https://github.com/BramTerlouw/Cloud-Minor-Team-2"
filePathOnGitHub="proto/resultProto.proto"

# Set the local directory to store the proto file
localDirectory="$(pwd)/proto"
localFilePath="$localDirectory/resultProto.proto"

echo "Deleting proto files in $localDirectory"
# Remove existing files in the proto directory
rm -f "$localDirectory/*"

echo "Downloading proto files from $githubRepo/raw/main/$filePathOnGitHub"
# Download the raw contents of the proto file from GitHub and save it locally
curl -LJO "$githubRepo/raw/main/$filePathOnGitHub" -o "$localFilePath"

echo "Generating Go code using protoc"
# Generate Go code using protoc
protoc --proto_path=./proto --go_out=./proto --go-grpc_out=./proto ./proto/resultProto.proto
