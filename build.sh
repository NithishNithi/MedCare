#!/bin/bash

# This script builds and prepares the Go application for deployment on Elastic Beanstalk

# Set environment variables
export GOOS=linux
export GOARCH=amd64

# Install dependencies
go mod tidy

# Build the Go application
go build -o medcare

# Exit
exit 0
