#!/usr/bin/env sh
# This script is used in the README and https://www.mh.io/docs/#quick-start
set -e

os=$(uname | tr '[:upper:]' '[:lower:]')
arch=$(uname -m | tr '[:upper:]' '[:lower:]' | sed -e s/x86_64/amd64/)
name="mh_linux_$os_$arch"
echo "Downloading latest release from https://github.com/modulehub/mh/releases/latest/download/$name.tar.gz..."
curl -sL https://github.com/modulehub/mh/releases/latest/download/$name.tar.gz | tar xz -C /tmp
echo
echo "Moving /tmp/$name to /usr/local/bin/mh (you might be asked for your password due to sudo)"
sudo mv /tmp/$name /usr/local/bin/mh
echo
echo "Completed installing $(mh --version)"
