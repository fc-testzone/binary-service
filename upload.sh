#!/bin/bash

GOOS=linux GOARCH=arm go build
mv binary-service ../binary-storage
cd ../binary-storage
git add -A
git commit --amend
git push -f