# Walldo

**Walldo** is a small cross-platform wallpaper changer written in Go.

It exists because I got tired of using Windows wallpaper managers that felt unnecessarily heavy
for such a simple task.
All I wanted was something that could recursively list my wallpapers, let me search through
them, and change the desktop background.
Basically, something with the spirit of **Nitrogen** on Linux, but available on every major
desktop platform.

Walldo focuses on doing one thing well:
browse your wallpaper collection and change your desktop background.
No background services, no live wallpapers, no video playback, and no unnecessary extras.

## Why another wallpaper changer?

Because I couldn't find one that actually matched what I needed.

Some wallpaper applications try to become full desktop customization suites.
Others run permanently in the background, offering support for video wallpapers, online
galleries, scheduling, playlists, transitions, and countless other features. 

On top of that, many modern desktop applications rely heavily on web technologies, shipping an
entire browser engine just to display a simple interface.

That is perfectly fine if that is what you are looking for, but Walldo is not that.

The goal was simple:
browse a collection of wallpapers and set one as the desktop background as quickly as possible.
This project was built with lower-end Windows laptops in mind, where most existing wallpaper
managers feel much heavier than they need to be.

## Features

- Recursive wallpaper discovery.
- Native support for Windows, macOS and most Linux desktop environments.
- Small native binary, with no background services.
- Fast fuzzy search filter.
- Simple yet visualy pleasing UI.

## Supported Formats

- `.jpg` / `.jpeg`
- `.png`

## Supported Operating Systems

- Windows
- Linux
- macOS*

NOTE:
In theory, Walldo should compile and run on macOS without any major issues.
However, since I do not own any Apple hardware, I cannot guarantee that everything works as
expected. 

Cross-compiling for macOS is also considerably more painful than for Windows or Linux, even
with the tooling provided by Fyne.
Because of this, no pre-built binaries are available, and macOS users are currently required to
build the application from source.

Bug reports and pull requests from macOS users are always welcome to help improve support.

### Supported Desktop Environments (Linux)

Walldo changes wallpapers by talking directly to the tools or APIs provided by each **desktop
environment** whenever possible.

For standalone X11 **window managers** it falls back to `feh`, and for **Wayland** compositors
it uses `swaybg`.

If your preferred environment isn't supported yet, feel free to open an issue.

| Environment | Backend |
| --- | --- |
| **GNOME** | Native |
| **KDE Plasma** | Native |
| **XFCE** | Native |
| **Cinnamon** | Native |
| **LXDE / LXQT** | Native |
| **MATE** | Native |
| **Deepin** | Native |
| **Window Managers (X11)** | [feh](https://wiki.archlinux.org/title/Feh) |
| **Wayland Compositors** | [swaybg](https://github.com/swaywm/swaybg) |

## Screenshots

## Installation

### Pre-built binaries

Pre-built binaries are available for Windows and Linux from the releases page.

https://github.com/elias-gill/walldo-in-go/releases

> **Windows note:** Since Walldo is an unsigned open source executable, Windows SmartScreen may display a warning the first time you launch it.

macOS binaries are currently not provided.
Please build from source.

### Building from source

Make sure you have Go installed.

https://go.dev/doc/install

#### Windows

```sh
go install -ldflags -H=windowsgui github.com/elias-gill/walldo-in-go@latest
```

#### Linux & macOS

```sh
go install github.com/elias-gill/walldo-in-go@latest
```

Run it with:

```sh
walldo-in-go
```

## Uninstallation

Walldo includes a simple uninstall flag.

```sh
walldo-in-go -uninstall
```

This removes the installed executable, cache and config files without requiring any additional
cleanup.

## Roadmap

Walldo is very close to being feature complete.

Its goal has always been simple:
recursively list images and change the desktop wallpaper.

Future development will mostly focus on:

- Bug fixes.
- Supporting additional desktop environments when possible.
- General code cleanup and maintenance.
- Small quality-of-life improvements.

The following features are intentionally out of scope:

- Video wallpapers.
- Live wallpapers.
- Automatic wallpaper rotation.
- Scheduled wallpaper changes.
- Background daemons.

Those features require a completely different architecture built around services running
continuously in the background, which goes against the purpose of this project.

## Feedback & Contributions

Bug reports, suggestions, and pull requests are always welcome.

https://github.com/elias-gill/walldo-in-go/issues

## Acknowledgments

- **Fyne**, for providing a clean, simple, and enjoyable cross-platform GUI toolkit for Go.
  Walldo would have been a much larger project without it.
- **ktr0731**, for the fuzzy matching algorithms.
- **reujab**, for the wallpaper bindings that inspired Walldo's initial abstraction layer.
