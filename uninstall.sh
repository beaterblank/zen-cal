#!/bin/bash
set -euo pipefail

mv ~/.config/zen-cal/config.jsonc.zen-cal.bak ~/.config/waybar/config.jsonc

./purge.sh

omarchy-restart-waybar

echo "zen-cal uninstalled successfully."
