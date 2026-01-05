#!/bin/bash
set -euo pipefail

# Use absolute path for consistency
BIN_DEST="$HOME/.local/bin/zen-cal"

if [[ -e "$BIN_DEST" ]]; then
    echo "Error: $BIN_DEST already exists"
    exit 1
fi

# Create directories
mkdir -p "$HOME/.config/hypr/apps"
mkdir -p "$HOME/.config/zen-cal"
mkdir -p "$HOME/.local/bin"

# Create config files if they don't exist
touch "$HOME/.config/hypr/apps.conf"

# Link apps.conf to zen-cal.conf
if ! grep -Fq "source = $HOME/.config/hypr/apps/zen-cal.conf" "$HOME/.config/hypr/apps.conf"; then
    echo "source = $HOME/.config/hypr/apps/zen-cal.conf" >> "$HOME/.config/hypr/apps.conf"
fi

# Add window rules for zen-cal
cp ./assets/window-rule/zen-cal.conf "$HOME/.config/hypr/apps/"

# Source apps.conf from hyprland.conf
if ! grep -Fq "source = $HOME/.config/hypr/apps.conf" "$HOME/.config/hypr/hyprland.conf"; then
    echo "source = $HOME/.config/hypr/apps.conf" >> "$HOME/.config/hypr/hyprland.conf"
fi

# Waybar Integration
WAYBAR_CONFIG="$HOME/.config/waybar/config.jsonc"
if [[ -f "$WAYBAR_CONFIG" ]]; then
    cp "$WAYBAR_CONFIG" "$HOME/.config/zen-cal/config.jsonc.zen-cal.bak"
    
    # This strips comments and trailing commas for JQ processing
    sed -E 's|//.*||g; s/,([[:space:]]*[\]}])/\1/g' "$WAYBAR_CONFIG" > /tmp/waybar_clean.jsonc
    
    # check to see if module is already in the list then merge
    if ! jq -s '.[0] * .[1] | if (."modules-right" | contains(["custom/zen-cal"])) then . else .["modules-right"] += ["custom/zen-cal"] end' \
      /tmp/waybar_clean.jsonc ./assets/waybar/waybar.json > /tmp/waybar_merged.jsonc; then
        echo "Error: Failed to merge waybar config"
        rm -f /tmp/waybar_clean.jsonc /tmp/waybar_merged.jsonc
        exit 1
    fi

    mv /tmp/waybar_merged.jsonc "$WAYBAR_CONFIG"
    rm -f /tmp/waybar_clean.jsonc
else
    echo "Warning: $WAYBAR_CONFIG not found, skipping waybar integration"
fi

# Copy zen-cal assets
cp ./assets/zen-cal/zen-cal.conf "$HOME/.config/zen-cal/"

# Build Go project
go mod tidy
go build -o zen-cal

# Place it on path
cp zen-cal "$BIN_DEST"
rm zen-cal

# Restart waybar
if command -v omarchy-restart-waybar &> /dev/null; then
    omarchy-restart-waybar
elif pgrep -x waybar > /dev/null; then
    killall -USR2 waybar || echo "Note: Restart waybar manually to apply changes."
fi

echo "Installed successfully"
