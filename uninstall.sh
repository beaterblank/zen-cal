#!/bin/bash
set -euo pipefail

# Restore waybar config backup if it exists
if [[ -f $HOME/.config/zen-cal/config.jsonc.zen-cal.bak ]]; then
    mv $HOME/.config/zen-cal/config.jsonc.zen-cal.bak $HOME/.config/waybar/config.jsonc
else
    echo "Warning: waybar config backup not found, skipping restore"
fi

./purge.sh

# Restart waybar if command exists
if command -v omarchy-restart-waybar &> /dev/null; then
    omarchy-restart-waybar
elif command -v killall &> /dev/null && killall -0 waybar 2>/dev/null; then
    echo "Note: omarchy-restart-waybar not found. Please restart waybar manually if needed."
fi

echo "zen-cal uninstalled successfully."
