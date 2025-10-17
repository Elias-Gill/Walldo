# Walldo 🖼️

**Walldo** is a minimal, cross-platform wallpaper changer written in Go.
It focuses on speed, simplicity, and low resource usage — no Electron, no daemons, no unnecessary dependencies.
Just a small binary that changes your wallpaper instantly.

Whether you're on Windows, Linux, or macOS, Walldo delivers a consistent experience with minimal system impact.

### Why another wallpaper changer ?

Most wallpaper changers I’ve tried are either heavy, slow, or tied to a specific OS or desktop environment. Many rely on frameworks like Electron or include unnecessary background services that constantly consume CPU and memory. I wanted a tool that behaved differently — something that would:

- Start instantly.
- Do its job and exit.
- Work the same across all major platforms.

Walldo was built for users who value minimalism and responsiveness, inspired by the simplicity and focus of classic Linux utilities like Nitrogen, where performance and purpose come before everything else.

## Features

* **Lightweight and fast**:
  written in Go, designed to start and run instantly.
* **Cross-platform**:
  works on Windows, Linux, and macOS.
* **Automatic image detection**:
  scans folders recursively to find images.
* **No background processes**:
  runs only when needed, then exits cleanly.

### Supported Formats

* JPG
* PNG
* JPEG

### Supported Desktop Environments (Linux)

| Environment  | Supported  | Dependency                                  |
| ------------ | ---------- | ------------------------------------------- |
| GNOME        | ✅         | Built-in                                    |
| KDE          | ✅         | Built-in                                    |
| XFCE         | ✅         | Built-in                                    |
| Cinnamon     | ✅         | Built-in                                    |
| LXDE / LXQT  | ✅         | Built-in                                    |
| MATE         | ✅         | Built-in                                    |
| Deepin       | ✅         | Built-in                                    |
| Non-DE (Feh) | ✅         | [Feh](https://wiki.archlinux.org/title/Feh) |
| Wayland      | ✅         | [swaybg](https://github.com/swaywm/swaybg)  |

## Installation

### Pre-compiled Binaries

Pre-built binaries are available for **Windows** and **Linux**.
Check the [releases section](https://github.com/Elias-Gill/walldo/releases) for the latest version.

There are currently no pre-compiled binaries for macOS due to the difficulty of cross-compiling and testing releases without direct access to Apple hardware.
In theory, Walldo should build and run correctly on macOS using the manual installation.

`Note:` Windows might warn you about the binary being unsigned.

### Manual Installation

To build Walldo manually, you need [Go](https://go.dev/doc/install) installed.

#### Windows

1. Open a terminal (`cmd`).
2. Run:

   ```sh
   go install -ldflags -H=windowsgui github.com/elias-gill/walldo-in-go@latest
   ```
3. Walldo will be available in your search bar (`Win` + `S`) as "walldo-in-go".
4. (Optional) Create a shortcut and set a custom icon.

#### Linux and macOS

1. Open a terminal.
2. Run:

   ```sh
   go install github.com/elias-gill/walldo-in-go@latest
   ```
3. You can now run Walldo by typing `walldo-in-go` in your terminal.
4. (Optional) Create a desktop shortcut for easier access.

### Uninstallation

To uninstall Walldo, run:

```sh
walldo-in-go -uninstall
```

This will remove the executable and clean up the installation.

## Screenshots

![image](https://github.com/user-attachments/assets/0563faa1-8430-42e2-92e7-22c807e8e236) <img width="1365" height="739" alt="image" src="https://github.com/user-attachments/assets/952bdefe-a55b-4dd8-b29a-7f6e2279b2d7" />
![image](https://github.com/user-attachments/assets/9d5fde99-9cda-47d4-ba09-9417f01df531)

## Goals

* Practical replacement for tools like Nitrogen (small, native and fast).
* Speed and efficiency over design.
* Focus and simplicity over features — _"do one thing and do it well"_.

Walldo will never support live wallpapers or dynamic backgrounds; it is intentionally limited to simple, static wallpaper changes.

## Future Plans

Walldo is mostly feature-complete. Future updates will focus on:

- Minor usability improvements and basic UI adjustments.
- Code cleanup, performance enhancements, stability checks, and security updates.
- Support for additional desktop environments.
- Bug fixes.
- Command line tool integrations.

## Feedback

If you find a bug or have suggestions, open an issue on [GitHub](https://github.com/Elias-Gill/walldo/issues).

## Mentions 🙏

* **ktr0731** — for the fuzzy finder engine.
* **reujab** — for the original wallpaper-changing library Walldo was based on.
