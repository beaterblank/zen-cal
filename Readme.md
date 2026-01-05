# Zen-Cal

A minimal, interactive terminal-based calendar built in **Go**, with month and year navigation using vim-style or arrow keys.

<img src="assets/screenshot.png" alt="zen-cal" />

## Controls

| Key                  | Action         |
| -------------------- | -------------- |
| `h`, `←`             | Previous month |
| `l`, `→`             | Next month     |
| `k`, `↑`             | Previous year  |
| `j`, `↓`             | Next year      |
| `r`, `↵`             | Reset to today |
| `q`, `Ctrl+C`, `esc` | Quit           |

## Requirements

* Go 1.20 or higher
* UTF-8 compatible terminal

## Installation

```bash
git clone https://github.com/beaterblank/zen-cal.git
cd zen-cal
chmod +x ./install.sh
./install.sh
```

## Uninstallation

```bash
chmod +x ./uninstall.sh ./purge.sh
./uninstall.sh
# optionally, remove all files using purge.sh
```

## Notes

* The installer modifies your Waybar configuration in-place, but creates a backup before making changes, backup your config.

## Dependencies

* Omarchy
* `github.com/charmbracelet/bubbletea`
* `github.com/charmbracelet/lipgloss`
