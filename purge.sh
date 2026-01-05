#!/bin/bash
set -euo pipefail

ZEN_BIN="$HOME/.local/bin/zen-cal"
ZEN_DIR="$HOME/.config/zen-cal"
ZEN_CONF="$HOME/.config/hypr/apps/zen-cal.conf"
APPS_CONF="$HOME/.config/hypr/apps.conf"
HYPR_CONF="$HOME/.config/hypr/hyprland.conf"

ZEN_LINE="source = $ZEN_CONF"
APPS_LINE="source = $APPS_CONF"

# Remove zen-cal binary
rm -f "$ZEN_BIN"
# Remove zen-cal assets
rm -rf "$ZEN_DIR"
# Remove zen-cal app config
rm -f "$ZEN_CONF"
# Remove zen-cal source line from apps.conf
if [[ -f "$APPS_CONF" ]] && grep -Fxq "$ZEN_LINE" "$APPS_CONF"; then
    sed -i "\|^$ZEN_LINE$|d" "$APPS_CONF"
fi
# Remove apps.conf source from hyprland.conf if apps.conf has no source entries
if [[ -f "$APPS_CONF" && -f "$HYPR_CONF" ]] && ! grep -q '^source[[:space:]]*=' "$APPS_CONF"; then
    sed -i "\|^$APPS_LINE$|d" "$HYPR_CONF"
fi

echo "Uninstalled successfully"
