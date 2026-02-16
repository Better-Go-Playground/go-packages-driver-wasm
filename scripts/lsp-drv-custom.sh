#!/usr/bin/env sh
DRV_ROOT="$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)"

if [ -z "$TEST_DIR" ]; then
  echo "Error: missing test repo dir. Set TEST_DIR env var"
  exit 1
fi

# Open neovim with gopls server using standard transport.
# Script used to run gopls with a custom driver and collect received data.
export GOPLS_BIN=gopls-devel
export GOPACKAGESDRIVER=pipetransport
export GOPACKAGESDRIVERADDR="unix:$DRV_ROOT/driver.sock"

export LSP_LOG_FILE="$DRV_ROOT/logs/gopls.log"
export LSP_RPC_TRACE="$DRV_ROOT/logs/rpc-got.trace.jsonl"

cd "$TEST_DIR" && nvim .
