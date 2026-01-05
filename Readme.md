# Zen-Cal

A minimal, interactive terminal-based calendar built in **Go**, with month and year navigation using vim-style or arrow keys.

<div style="display: flex; gap: 12px;">
  <img src="assets/screenshot-dark.png" alt="zen-cal dark" style="width: 50%;" />
  <img src="assets/screenshot-light.png" alt="zen-cal light" style="width: 50%;" />
</div>


## Configuration

Zen-Cal can be customized to match your theme. The following defaults work well with most dark themes:

```toml
today    = #f38ba8
headings = #cba6f7
text     = #cdd6f4
weekends = #f9e2af
```

You can adjust these values to match your preferred color scheme in `~/.config/zen-cal/zen-cal[ dark / light ].config`.

`~/.config/hypr/app/zen-cal.conf` file defines the window rules; adjust them to position the calendar anywhere on the monitor.

It should look like this:

```
windowrulev2 = float, class:^(TUI.zencal)$
windowrulev2 = size 14% 18%, class:^(TUI.zencal)$
windowrulev2 = move 86% 2.5%, class:^(TUI.zencal)$
```

* `move` controls the window’s position on the monitor.
* `size` controls the window’s dimensions.

By default, the module appears in the right-most corner. If you want it centered, update your Waybar configuration aswell:

* Add `custom/zen-cal` to `modules-center`.
* Remove `custom/zen-cal` from the `modules-right`.

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
> **Note:** The installer updates your Waybar configuration directly but first creates a backup to restore on uninstall, make sure to keep a copy of your config in case recovery is needed.

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



## Dependencies

* hyprland
* Waybar
* go lang
* jq (installation dependency)
* `github.com/charmbracelet/bubbletea`
* `github.com/charmbracelet/lipgloss`
