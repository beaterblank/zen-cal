#!/bin/bash
set -euo pipefail

if [[ -e "$HOME/.local/bin/zen-cal" ]]; then
    echo "Error: ~/.local/bin/zen-cal already exists"
    exit 1
fi

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
cp ./assets/window-rule/zen-cal.conf ~/.config/hypr/apps/


# Source apps.conf from hyprland.conf
if ! grep -q 'source = ~/.config/hypr/apps.conf' ~/.config/hypr/hyprland.conf; then
    echo "source = ~/.config/hypr/apps.conf" >> ~/.config/hypr/hyprland.conf
fi

# Extend waybar
cp ~/.config/waybar/config.jsonc ~/.config/zen-cal/config.jsonc.zen-cal.bak
# Remove trailing commas into a temp file
sed -E 's/,([[:space:]]*[\]}])/\1/g' ~/.config/waybar/config.jsonc > /tmp/waybar_clean.jsonc
# Merge with waybar JSON into another temp file
jq -s '.[0] * .[1] | .["modules-right"] = (.["modules-right"] // []) + ["custom/zen-cal"]' \
  /tmp/waybar_clean.jsonc ./assets/waybar/waybar.json > /tmp/waybar_merged.jsonc
# Move the merged result to the actual config
mv /tmp/waybar_merged.jsonc ~/.config/waybar/config.jsonc
# Clean up temp file
rm /tmp/waybar_clean.jsonc

# Copy zen-cal assets
cp ./assets/zen-cal/zen-cal.conf ~/.config/zen-cal/

# Build Go project
go mod tidy
go build

# place it on path
mkdir -p ~/.local/bin/
cp zen-cal ~/.local/bin/
rm zen-cal

omarchy-restart-waybar

echo "installed successfully"
