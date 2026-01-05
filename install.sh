#!/bin/bash
set -euo pipefail

# Create directories
mkdir -p ~/.config/hypr/apps
mkdir -p ~/.config/zen-cal

# Create config files if they don't exist
touch ~/.config/hypr/apps.conf
touch ~/.config/hypr/apps/zen-cal.conf

# Link apps.conf to zen-cal.conf
if ! grep -q 'source = ~/.config/hypr/apps/zen-cal.conf' ~/.config/hypr/apps.conf; then
  echo "source = ~/.config/hypr/apps/zen-cal.conf" > ~/.config/hypr/apps.conf
fi

# Add window rules for zen-cal
cat <<EOF > ~/.config/hypr/apps/zen-cal.conf
windowrulev2 = float, class:^(TUI.zencal)\$
windowrulev2 = size 14% 16%, class:^(TUI.zencal)\$
windowrulev2 = move 44% 2%, class:^(TUI.zencal)\$
EOF

# Source apps.conf from hyprland.conf
if ! grep -q 'source = ~/.config/hypr/apps.conf' ~/.config/hypr/hyprland.conf; then
    echo "source = ~/.config/hypr/apps.conf" >> ~/.config/hypr/hyprland.conf
fi

# Add on-click action to waybar clock (removes any previously existing on-click)
sed -i '/"clock"[[:space:]]*:[[:space:]]*{/,/^[[:space:]]*}/{
  # Delete any existing on-click line inside the block
  /^[[:space:]]*"on-click"[[:space:]]*:/d
  # Insert the new on-click before the closing brace
  /^[[:space:]]*}/i\    "on-click": "hyprctl clients | grep -q \\"initialClass: TUI.zencal\\" && hyprctl dispatch focuswindow \\"initialClass: TUI.zencal\\" || xdg-terminal-exec --app-id=TUI.zencal -e zen-cal",
}' ~/.config/waybar/config.jsonc

# Copy zen-cal assets
cp ./assets/zen-cal.conf ~/.config/zen-cal/

# Build Go project
go mod tidy
go build

# place it on path
mkdir -p ~/.local/bin
cp zen-cal ~/.local/bin/

echo "installed successfully"
