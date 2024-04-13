#!/bin/sh

# Install npm dependencies
# npm install

# Build typescript
npx tsc

# Build go

# first cd into app then run go build
cd app
go build -o ../build/app cmd/main.go
cd ../

# copy static folder to build
cp -r app/static build/
cd build

./app
