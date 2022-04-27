#!/usr/bin/env bash

GH_REPO="axllent/golp"
TIMEOUT=90

set -e

VERSION=$(curl --silent --location --max-time "${TIMEOUT}" "https://api.github.com/repos/${GH_REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
if [ $? -ne 0 ]; then
  echo -ne "\nThere was an error trying to check what is the latest version of ssbak.\nPlease try again later.\n"
  exit 1
fi

# detect the platform
OS="$(uname)"
case $OS in
Linux)
  OS='linux'
  ;;
FreeBSD)
  OS='freebsd'
  echo 'OS not supported'
  exit 2
  ;;
NetBSD)
  OS='netbsd'
  echo 'OS not supported'
  exit 2
  ;;
OpenBSD)
  OS='openbsd'
  echo 'OS not supported'
  exit 2
  ;;
Darwin)
  OS='darwin'
  ;;
SunOS)
  OS='solaris'
  echo 'OS not supported'
  exit 2
  ;;
*)
  echo 'OS not supported'
  exit 2
  ;;
esac

# detect the arch
OS_type="$(uname -m)"
case "$OS_type" in
x86_64 | amd64)
  OS_type='amd64'
  ;;
i?86 | x86)
  OS_type='386'
  ;;
aarch64 | arm64)
  OS_type='arm64'
  ;;
*)
  echo 'OS type not supported'
  exit 2
  ;;
esac

GH_REPO_BIN="golp_${OS}_${OS_type}.tar.gz"

#create tmp directory and move to it with macOS compatibility fallback
tmp_dir=$(mktemp -d 2>/dev/null || mktemp -d -t 'golp-install.XXXXXXXXXX')
cd "$tmp_dir"

echo "Downloading golp $VERSION"
LINK="https://github.com/${GH_REPO}/releases/download/${VERSION}/${GH_REPO_BIN}"

curl --silent --location --max-time "${TIMEOUT}" "${LINK}" | tar zxf - || {
  echo "Error downloading"
  exit 2
}

echo "Installing"
mkdir -p /usr/local/bin || exit 2
cp golp /usr/local/bin/ || exit 2
chown root:root /usr/local/bin/golp || exit 2
chmod 755 /usr/local/bin/golp || exit 2

rm -rf "$tmp_dir"
echo "Installed successfully to /usr/local/bin/golp"