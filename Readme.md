# Zen-Cal
Terminal UI calendar with navigation of months and years with vim-style and arrow keys
A minimal, interactive terminal-based calendar built in **Go** using
[Bubble Tea](https://github.com/charmbracelet/bubbletea) and
[Lip Gloss](https://github.com/charmbracelet/lipgloss).

<p align="center">
  <img src="assets/screenshot.png" alt="zen-cal" width="300"/>
</p>

## Controls

| Key             | Action         |
| ----------------| -------------- |
| `h`, `←`            | Previous month |
| `l`, `→`            | Next month     |
| `k`, `↑`            | Previous year  |
| `j`, `↓`            | Next year      |
| `r`, `↵`            | Reset to today |
| `q`, `Ctrl+C`, `esc`  | Quit           |

## Requirements
* Go 1.20+
* Terminal with UTF-8 support

## Note
This tries to modify your waybar config in-place
although it backs it up and restores it 
back it up somewhere

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
# optionally just remove the files with purge.sh
```
```
```


## Dependencies
* Omarchy
* `github.com/charmbracelet/bubbletea`
* `github.com/charmbracelet/lipgloss`
