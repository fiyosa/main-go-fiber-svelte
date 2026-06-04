#!/bin/sh
set -e

echo "Building SPA..."
cd src/web
npm ci
npm run build

echo "Building Go binary..."
cd ../..
go build -o app .

echo "Starting server..."
./app
