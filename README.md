# Walldo 🖼️

**A lightning-fast, lightweight wallpaper changer written in Go.** Walldo is designed to make
changing wallpapers simple, fast, and hassle-free.
Whether you're on Windows, Linux, or macOS, Walldo delivers a seamless experience with minimal
bloat.

## ✨ Features

- **Simple, Fast and Lightweight**:
  Built for simplicity and speed.
- **Multi-Platform Support**:
  Works on Windows, Linux, and macOS.
- **Automatic Image Detection**:
  Scans folders recursivelly to search images.

## 🖼️ Supported Formats

- JPG
- PNG
- JPEG

## 🖥️ Supported Desktop Environments

Walldo supports a variety of desktop environments and tools:

| Environment       | Supported? | Dependency       |
|-------------------|------------|------------------|
| GNOME             | ✅         | Built-in         |
| KDE               | ✅         | Built-in         |
| XFCE              | ✅         | Built-in         |
| Cinnamon          | ✅         | Built-in         |
| LXDE / LXQT       | ✅         | Built-in         |
| MATE              | ✅         | Built-in         |
| Deepin            | ✅         | Built-in         |
| Non-DE (Feh)      | ✅         | [Feh](https://wiki.archlinux.org/title/Feh) |
| Wayland           | ✅         | [swaybg](https://github.com/swaywm/swaybg) |

## 🚀 Installation

### Pre-Compiled Binaries
We provide pre-compiled binaries for **Windows** and **Linux**.
Check out our [releases section](https://github.com/Elias-Gill/walldo/releases) to download the
latest version.

`Note`:
Windows may complain about the binary being unsigned.

### Manual Installation
To compile Walldo manually, you need to have [Go](https://go.dev/doc/install) installed.

#### Windows
1. Open a terminal session (`cmd`).
2. Run the following command:
   ```sh
   go install -ldflags -H=windowsgui github.com/elias-gill/walldo-in-go@latest
   ```
3. Walldo will be available in your search bar (`Win`+`S`) as "waldo-in-go".
4. (Optional) Create a shortcut and customize the icon.

#### Linux and macOS
1. Open a terminal session.
2. Run the following command:
   ```sh
   go install github.com/elias-gill/walldo-in-go@latest
   ```
3. You can now run Walldo by typing `walldo-in-go` in your terminal.
4. (Optional) Create a desktop shortcut for easier access.

### 🛠️ Uninstallation

To uninstall Walldo, simply run the following command in your terminal:

```sh 
walldo-in-go -uninstall
```

This will remove the executable and clean up the installation.

## 🎯 Goals

- **Replace Nitrogen or Feh on Windows**:
  Provide a lightweight alternative for Windows users.
- **Speed and Efficiency**:
  Be faster and lighter than other wallpaper-changing apps.
- **Simplicity and Focus**:
  Walldo is designed to do one thing and do it well—change your wallpaper as quickly as
  possible.
  No bloat, no unnecessary features, and no long startup times.
  Just a simple, fast, and reliable tool.

## 🙏 Mentions

- **ktr0731**:
  For the amazing fuzzy finder engine.
- **reujab**:
  For the original wallpaper-changing library, which was forked and adapted for Walldo.

## 📸 Showcase

https://github.com/user-attachments/assets/89c2ad3e-d5c8-4928-8bca-bddffefd0b12

## ❓ Why Walldo?

As someone who loves changing wallpapers frequently, I wanted an app that was simple, fast, and
bloat-free.
While Linux has many options, Windows often lacks lightweight tools for this purpose.
Walldo fills that gap, offering a no-nonsense solution for wallpaper enthusiasts.

## 🚧 Future Plans

Walldo is almost feature-complete, with future updates focusing on:
- Enhancing user experience.
- Adding security checks.
- Improving performance.

## 💬 Feedback

Love Walldo?
Found a bug?
Let us know by opening an issue on [GitHub](https://github.com/Elias-Gill/walldo/issues).

## ⭐ Star the Repo

If you find Walldo useful, please consider starring the repository to show your support!

[![Star on GitHub](https://img.shields.io/github/stars/Elias-Gill/walldo?style=social)](https://github.com/Elias-Gill/walldo)
