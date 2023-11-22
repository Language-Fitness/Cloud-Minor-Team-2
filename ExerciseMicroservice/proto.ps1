# Set your GitHub repository details

$githubRepo = "https://github.com/BramTerlouw/Cloud-Minor-Team-2"
$filePathOnGitHub = "proto/exerciseProto.proto"

# Set the local directory to store the proto file
$localDirectory = "$PSScriptRoot/proto"
$localFilePath = "$localDirectory/exerciseProto.proto"

Write-Host "Deleting proto files in $PSScriptRoot/proto"
# Remove existing files in the proto directory
Remove-Item "$localDirectory\*" -Force -Recurse

Write-Host "downloading proto files from $githubRepo/raw/main/$filePathOnGitHub"
# Download the raw contents of the proto file from GitHub
Invoke-WebRequest -Uri "$githubRepo/raw/main/$filePathOnGitHub" -OutFile $localFilePath

Write-Host "Generating Go code using protoc"
# Generate Go code using protoc
protoc --proto_path=.\proto --go_out=.\proto --go-grpc_out=.\proto .\proto\exerciseProto.proto

