#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

echo "==> Checking Go..."
go version

echo "==> Downloading dependencies..."
go mod download
go mod verify

echo "==> Building binary..."
mkdir -p bin
go build -o bin/go-ping-reddit .

if [[ ! -f reddit-account.agent ]]; then
  cp reddit-account.agent.example reddit-account.agent
  echo "==> Created reddit-account.agent from template."
  echo "    Edit it with your Reddit script app credentials for live mode."
else
  echo "==> reddit-account.agent already exists."
fi

echo
echo "Setup complete."
echo "  make demo   # offline demo (no credentials)"
echo "  make run    # live Reddit API (requires credentials)"
