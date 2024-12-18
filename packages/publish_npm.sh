#!/bin/bash

set -e

# Mac
mkdir -p ./packages/npm-mcptee-darwin-x64/bin
cp dist/mcptee_darwin_amd64_v1/mcptee ./packages/npm-mcptee-darwin-x64/bin/mcptee
chmod +x ./packages/npm-mcptee-darwin-x64/bin/mcptee
mkdir -p ./packages/npm-mcptee-darwin-arm64/bin
cp dist/mcptee_darwin_arm64_v8.0/mcptee ./packages/npm-mcptee-darwin-arm64/bin/mcptee
chmod +x ./packages/npm-mcptee-darwin-arm64/bin/mcptee

# Linux
mkdir -p ./packages/npm-mcptee-linux-x64/bin
cp dist/mcptee_linux_amd64_v1/mcptee ./packages/npm-mcptee-linux-x64/bin/mcptee
chmod +x ./packages/npm-mcptee-linux-x64/bin/mcptee
mkdir -p ./packages/npm-mcptee-linux-arm64/bin
cp dist/mcptee_linux_arm64_v8.0/mcptee ./packages/npm-mcptee-linux-arm64/bin/mcptee
chmod +x ./packages/npm-mcptee-linux-arm64/bin/mcptee

# Windows
mkdir -p ./packages/npm-mcptee-win32-x64/bin
cp dist/mcptee_windows_amd64_v1/mcptee.exe ./packages/npm-mcptee-win32-x64/bin/mcptee.exe
mkdir -p ./packages/npm-mcptee-win32-arm64/bin
cp dist/mcptee_windows_arm64_v8.0/mcptee.exe ./packages/npm-mcptee-win32-arm64/bin/mcptee.exe

cd packages/npm-mcptee-darwin-x64
npm publish --access public

cd ../npm-mcptee-darwin-arm64
npm publish --access public

cd ../npm-mcptee-linux-x64
npm publish --access public

cd ../npm-mcptee-linux-arm64
npm publish --access public

cd ../npm-mcptee-win32-x64
npm publish --access public

cd ../npm-mcptee-win32-arm64
npm publish --access public

cd ../npm-mcptee
npm publish --access public

cd -