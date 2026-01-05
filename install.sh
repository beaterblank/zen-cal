#!/bin/bash
set -euo pipefail

BIN_DIR="$HOME/.local/bin"
BIN_DEST="$BIN_DIR/zen-cal"

HYPR_DIR="$HOME/.config/hypr"
HYPR_APPS_DIR="$HYPR_DIR/apps"
HYPR_APPS_CONF="$HYPR_DIR/apps.conf"
HYPR_MAIN_CONF="$HYPR_DIR/hyprland.conf"
HYPR_ZEN_CONF="$HYPR_APPS_DIR/zen-cal.conf"

ZEN_DIR="$HOME/.config/zen-cal"
ZEN_CONF_LIGHT="$ZEN_DIR/zen-cal-light.conf"
ZEN_CONF_DARK="$ZEN_DIR/zen-cal-dark.conf"

WAYBAR_DIR="$HOME/.config/waybar"
WAYBAR_CONFIG="$WAYBAR_DIR/config.jsonc"
WAYBAR_BACKUP="$ZEN_DIR/config.jsonc.zen-cal.bak"

TMP_WAYBAR_CLEAN="/tmp/waybar_clean.jsonc"
TMP_WAYBAR_MERGED="/tmp/waybar_merged.jsonc"

if [[ -e "$BIN_DEST" ]]; then
    echo "Error: $BIN_DEST already exists"
    exit 1
fi

mkdir -p "$BIN_DIR" "$HYPR_APPS_DIR" "$ZEN_DIR"

touch "$HYPR_APPS_CONF"

if ! grep -Fq "source = $HYPR_ZEN_CONF" "$HYPR_APPS_CONF"; then
    echo "source = $HYPR_ZEN_CONF" >> "$HYPR_APPS_CONF"
fi

cp ./assets/window-rule/zen-cal.conf "$HYPR_APPS_DIR/"

if ! grep -Fq "source = $HYPR_APPS_CONF" "$HYPR_MAIN_CONF"; then
    echo "source = $HYPR_APPS_CONF" >> "$HYPR_MAIN_CONF"
fi

if [[ -f "$WAYBAR_CONFIG" ]]; then
    cp "$WAYBAR_CONFIG" "$WAYBAR_BACKUP"

    sed -E 's|//.*||g; s/,([[:space:]]*[\]}])/\1/g' \
        "$WAYBAR_CONFIG" > "$TMP_WAYBAR_CLEAN"

    if ! jq -s '
        .[0] * .[1]
        | if (."modules-right" | contains(["custom/zen-cal"]))
          then .
          else .["modules-right"] += ["custom/zen-cal"]
          end
    ' "$TMP_WAYBAR_CLEAN" ./assets/waybar/waybar.json > "$TMP_WAYBAR_MERGED"; then
        echo "Error: Failed to merge waybar config"
        rm -f "$TMP_WAYBAR_CLEAN" "$TMP_WAYBAR_MERGED"
        exit 1
    fi

    mv "$TMP_WAYBAR_MERGED" "$WAYBAR_CONFIG"
    rm -f "$TMP_WAYBAR_CLEAN"
else
    echo "Warning: $WAYBAR_CONFIG not found, skipping waybar integration"
fi

cp ./assets/zen-cal/zen-cal-light.conf "$ZEN_CONF_LIGHT"
cp ./assets/zen-cal/zen-cal-dark.conf "$ZEN_CONF_DARK"


go mod tidy
go build -o zen-cal
cp zen-cal "$BIN_DEST"
rm zen-cal

if command -v omarchy-restart-waybar &> /dev/null; then
    omarchy-restart-waybar
elif pgrep -x waybar > /dev/null; then
    killall -USR2 waybar || echo "Note: Restart waybar manually to apply changes."
fi

echo "Installed successfully"
