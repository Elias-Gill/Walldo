# Walldo
A simple wallpaper changer written in Go with the main goal of being light and fast.

Mainly inspired by Nitrogen, Walldo searches for your images, lists them all 
and displays them on a single screen.

Change your wallpaper as fast as you want. Nothing more and nothing less, simple and functional.

### Currently supported formats
+ JPG
+ PNG
+ JPEG

https://user-images.githubusercontent.com/79729891/180588701-e58b76ca-60a5-4333-b573-171f988a3592.mp4

#### A wallpaper changer app ? Why ?... 
As a person who likes to change his wallpaper very often, I was looking for an app that would allow me to do it in a 
simple way and as fast as possible. In Linux there are many applications that can do that, 
but in Windows most of them are very slow and bloated.

## Instalation
We have pre-compiled binaries for windows and linux. Check out our [releases](https://github.com/Elias-Gill/walldo/releases) section (windows may complain because the binary is unisgned).

We don't provide precompiled builds for `macOS` because it's [not straightforward to cross-compose Fyne projects for it](https://github.com/fyne-io/fyne-cross#build-the-docker-image-for-osxdarwinapple-cross-compiling).

## Manual installation
To compile Walldo manually you need to have [Go](https://go.dev/doc/install) installed.

### Windows
After installing the dependencies, open a terminal session (`cmd`) and run:
```sh
go install -ldflags -H=windowsgui github.com/elias-gill/walldo-in-go@latest
```
Now Walldo should be available in your search bar (`Win`+`S`) as "waldo-in-go".

You can create a shortcut and change the icon if you want.

### Linux and macOS
After installing the dependencies, open a terminal session and run:
```sh
go install github.com/elias-gill/walldo-in-go@latest
```
You can now run "walldo-in-go" in your terminal. If you are using a Desktop environment, you can create a shortcut for the command.

## First run
 After the first run you have to setup your wallpapers folder. To do this open the config menu ( “⚙” button) and put your folders absolute path. 
You can put more than one folder separating them with commas (you can add line jumps between folders too).
 
You can reload Walldo with the reload button ("⟳") in the bottom right.
##### *Path format example:*
```
  C:/Users/walldo/Desktop/fondos,
  C:/Users/walldo/Desktop/wallpapers2
```
*NOTE: do not use backslashes, instead use normal slashes (/)*

## Important (Linux users)
Walldo supports a variaty of desktop enviroments:
- Gnome
- LXDE
- XFCE
- Cinnamon 
- Mate 
- KDE
- Deepin

For non desktop enviroments Walldo relies on [Feh](https://wiki.archlinux.org/title/Feh).
Wayland is supported via [swaybg](https://github.com/swaywm/swaybg).

## Goals
+ Be the replacement for Nitrogen or Feh on Windows Systems. 
+ Be faster and lighter than similar apps.

## Mentions
+ To [ktr0731](https://github.com/ktr07310). The fuzzy finder engine is his entire work.
+ To [reujab](https://github.com/reujab/wallpaper). The library for changing wallapapers is a fork of his original module.
