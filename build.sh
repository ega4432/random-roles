#!/bin/sh
rm clean_lab
GOOS="darwin" GOARCH="amd64" go build -o clean_lab
if [ -e clean_lab ]
then
echo "binary found."
chmod +x clean_lab
echo "grant execution permission."
else
echo "build binary not found."
fi
