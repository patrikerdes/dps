#!/usr/bin/env bash

if [[ $# -eq 0 ]] ; then
    echo 'please provide the version string'
    exit 1
fi

version=$1

echo "creating release files..."

set -x
cd osx 
zip ../dps_${version}_osx.zip *
cd ../windows
zip ../dps_${version}_windows.zip *
cd ../linux
zip ../dps_${version}_linux.zip *
