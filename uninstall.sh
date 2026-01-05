#!/bin/bash
set -euo pipefail

ZEN_DIR="$HOME/.config/zen-cal"
WAYBAR_DIR="$HOME/.config/waybar"

WAYBAR_CONFIG="$WAYBAR_DIR/config.jsonc"
WAYBAR_BACKUP="$ZEN_DIR/config.jsonc.zen-cal.bak"

if [[ -f "$WAYBAR_BACKUP" ]]; then
    mv "$WAYBAR_BACKUP" "$WAYBAR_CONFIG"
else
    echo "Warning: waybar config backup not found, skipping restore"
fi

./purge.sh

if command -v omarchy-restart-waybar &> /dev/null; then
    omarchy-restart-waybar
elif command -v killall &> /dev/null && killall -0 waybar 2>/dev/null; then
    echo "Note: omarchy-restart-waybar not found. Please restart waybar manually if needed."
fi

