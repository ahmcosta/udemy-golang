#!/bin/bash


cd $GOPATH/src/github/ahmcosta/udemy-golang

# Enable coverage analysis 
go test -cover 

# Write a coverage profile to the file after all tests have passed
go test -coverprofile=resultado.out ./... 

# Given a coverage profile produced by 'go test'
# Display coverage percentages to stdout for each function
go tool cover -func=resultado.out
# Open a web browser displaying annotated source code
go tool cover -html=resultado.out


