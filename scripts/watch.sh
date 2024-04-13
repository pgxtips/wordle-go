#!/bin/sh

# Define the build script command
BUILD_SCRIPT="./scripts/build.sh"

# Specify the folder to watch
WATCH_FOLDER="./app"

# Run nodemon to watch for changes in all files within the specified folder(s)
# and re-run the build script
npx nodemon --watch "$WATCH_FOLDER" --ext '*' --exec "$BUILD_SCRIPT"
