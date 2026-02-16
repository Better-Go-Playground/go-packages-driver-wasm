#!/usr/bin/env sh
DRV_ROOT="$PWD"

if [ -z "$TEST_DIR" ]; then
  echo "Error: missing test repo dir. Set TEST_DIR env var"
  exit 1
fi

# Open neovim with gopls server using standard transport.
# Script used to collect "expected" data.
export GOPLS_BIN=gopls-devel

export LSP_LOG_FILE="$DRV_ROOT/logs/lsp-standard.log"
export LSP_RPC_TRACE="$DRV_ROOT/logs/rpc-trace.expected.jsonl"
