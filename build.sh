#! /usr/bin/env bash


clear
echo "STARTING."
echo "Building hello-world"
echo "Formatting..."
go fmt ./...
echo "Cleaning..."
go clean
echo "Getting..."
go get -d
echo "Building..."
GOOS=linux CGO_ENABLED=0 go build -a  # binary for mac-> GOOS=darwin CGO_ENABLED=0 go build -a
echo "Done building hello-world"
echo "Moving to bin folder..."
rm -drf bin
mkdir bin
cp go-hello-world bin/go-hello-world
rm go-hello-world
echo "DONE."
