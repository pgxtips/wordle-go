#!/bin/sh

# Stop existing Go app process
pkill app

# Remove existing build folder
rm -rf build

# Install npm dependencies
# npm install

# Typescript transpile
# Will output .ts files to build/static/js
npx tsc

# copy static folder to build
cp -r app/static build/

# Views 
# Extract all html templates from views folder and put them in build/views
mkdir -p build/internal/views
find app/internal/views -type f -name "*.html" -exec sh -c '
    src_path="$1"
    dest_dir="build/internal/views/${src_path#app/internal/views/}"
    dest_dir="${dest_dir%/*}" # Remove the filename from the path
    mkdir -p "$dest_dir"
    cp "$src_path" "$dest_dir"
' _ {} \;

# CSS
# Extract all css files and put them in static/css
mkdir -p build/static/css
find app/internal/views -type f -name "*.css" -exec sh -c '
    src_path="$1"
    dest_dir="build/static/css/${src_path#app/internal/views/}"
    dest_dir="${dest_dir%/*}" # Remove the filename from the path
    mkdir -p "$dest_dir"
    cp "$src_path" "$dest_dir"
' _ {} \;

# JS Libraries 
# Jquery
mkdir -p build/static/js-libs/jquery
cp node_modules/jquery/dist/jquery.min.js build/static/js-libs/jquery/

# Build go
# first cd into app then run go build
cd app
go build -o ../build/app cmd/main.go
cd ../

cd build

./app
