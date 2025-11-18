#!/usr/bin/env bash
# Simple wrapper to execute mysql-backup.sh with the same arguments.
SCRIPT_PATH=$(readlink -f "${BASH_SOURCE[0]}" 2>/dev/null || realpath "${BASH_SOURCE[0]}")
SCRIPT_DIR=$(dirname "$SCRIPT_PATH")
exec "$SCRIPT_DIR/mysql-backup.sh" "$@"
