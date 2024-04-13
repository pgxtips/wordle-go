#!/bin/sh

# Stop existing Go app process
pkill app

# Install npm dependencies
# npm install

# Build go
# first cd into app then run go build
cd app
go build -o ../build/app cmd/main.go
cd ../

# copy static folder to build
cp -r app/static build/

# CSS
# Extract all css files and put them in static/css
cp app/internal/views/*.css build/static/css/

# Typescript transpile
# Will output .ts files to build/static/js
npx tsc

cd build

./app
