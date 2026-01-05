#!/bin/bash
set -euo pipefail

# Remove zen-cal binary
rm -f ~/.local/bin/zen-cal

# Remove zen-cal assets
rm -rf ~/.config/zen-cal

# Remove zen-cal app config
rm -f ~/.config/hypr/apps/zen-cal.conf

# Remove source line from apps.conf if it points to zen-cal.conf
if grep -q 'source = ~/.config/hypr/apps/zen-cal.conf' ~/.config/hypr/apps.conf; then
    sed -i '/source = ~\/.config\/hypr\/apps\/zen-cal.conf/d' ~/.config/hypr/apps.conf
fi

# Remove apps.conf source from hyprland.conf if it points to apps.conf (optional: only if installed by this script)
# Check if other apps configs exist; if not, remove source
if ! grep -q 'source = ~/.config/hypr/apps/' ~/.config/hypr/apps.conf; then
    sed -i '/source = ~\/.config\/hypr\/apps.conf/d' ~/.config/hypr/hyprland.conf
fi

# Remove on-click from waybar clock for zen-cal
sed -i '/"clock"[[:space:]]*:[[:space:]]*{/,/^[[:space:]]*}/{
  # Delete on-click lines containing zen-cal
  /^[[:space:]]*"on-click"[[:space:]]*:.*zen-cal/d
  # Remove trailing comma before closing brace
  /^[[:space:]]*},?$/s/,[[:space:]]*}/}/
}' ~/.config/waybar/config.jsonc

echo "zen-cal uninstalled successfully."
