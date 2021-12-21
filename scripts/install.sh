#!/usr/bin/env sh
# This script is used in the README and https://www.mh.io/docs/#quick-start
set -e

os=$(uname | tr '[:upper:]' '[:lower:]')
arch=$(uname -m | tr '[:upper:]' '[:lower:]' | sed -e s/x86_64/amd64/)
name="mh_${os}_${arch}"
url="https://github.com/modulehub/mh/releases/latest/download/$name.tar.gz"
echo "Downloading latest release from ${url}..."
curl -Lo ./mh.tar.gz ${url} || exit 1
tar -xzf mh.tar.gz || exit 1
chmod +x mh || exit 1
mv mh /usr/local/bin/mh || exit 1
echo "Completed installing $(mh --version)"
