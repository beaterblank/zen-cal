#!/bin/bash
set -euo pipefail
# Remove zen-cal binary
rm -f "$HOME/.local/bin/zen-cal"
# Remove zen-cal assets
rm -rf "$HOME/.config/zen-cal"
# Remove zen-cal app config
rm -f "$HOME/.config/hypr/apps/zen-cal.conf"
# Remove source line from apps.conf if it exists
if grep -Fq "source = $HOME/.config/hypr/apps/zen-cal.conf" "$HOME/.config/hypr/apps.conf" 2>/dev/null; then
    sed -i "s|source = $HOME/.config/hypr/apps/zen-cal.conf||d" "$HOME/.config/hypr/apps.conf"
fi
# Remove apps.conf source from hyprland.conf if apps.conf is now empty
if ! grep -q "source =" "$HOME/.config/hypr/apps.conf" 2>/dev/null; then
    sed -i "s|source = $HOME/.config/hypr/apps.conf||d" "$HOME/.config/hypr/hyprland.conf"
fi
echo "Uninstalled successfully"
