#!/usr/bin/env bash
# runs goxc in each product directory
set -e

echo "Building static linux binary for gosieve..."
mkdir -p pkg
buildcmd='CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o pkg/gosieve'
docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" \
  -w /app golang:1.6 sh -c "$buildcmd"

echo "Building docker image for gosieve..."
docker build --no-cache=true --tag benton/gosieve .
rm -f pkg/gosieve

exit 0
