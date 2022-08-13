# Walldo-in-Go
A simple wallpaper changer written in Go with the main goal of being light and fast.

Mainly inspired by Nitrogen (more like a knockoff  ¯\_( ͡❛͜ʖ ͡❛)_/¯ ), Walldo searches for your images, lists them all 
and displays them on a single screen.

Change your wallpaper as fast as you want. Nothing more and nothing less, simple and functional.

### Currently supported formats
- JPG
- PNG
- JPEG

https://user-images.githubusercontent.com/79729891/180588701-e58b76ca-60a5-4333-b573-171f988a3592.mp4

#### A wallpaper changer app ? Why ?... 
As a person who likes to change his wallpaper very often, I was looking for an app that would allow me to do it in a 
simple way and as fast as possible. In Linux there are many applications that can do that, 
but in Windows most of them are very slow and bloated.

## Instalation
### System Dependencies 
- [Go](https://go.dev/doc/install)

If you want to compile this yourself, you'll need to check out [fyne-cross](https://github.com/fyne-io/fyne-cross).

We don't provide precompiled builds for `macOS` because it's [not straightforward to cross-compose Fyne projects for it](https://github.com/fyne-io/fyne-cross#build-the-docker-image-for-osxdarwinapple-cross-compiling).

### Windows
After installing the dependencies, open a terminal session (`cmd`) and run:
````
go install -ldflags -H=windowsgui github.com/elias-gill/walldo-in-go@latest
````
Now Walldo should be available in your search bar (`Win`+`S`) as "waldo-in-go".

You can create a shortcut and change the icon if you want.

#### First run
 After the first run you have to setup your wallpapers folder. To do this open the config menu ( “⚙” button) and put your folders absolute path. 
You can put more than one folder separating them with commas
 
You can reload Walldo with the reload button ("⟳") in the bottom right.
##### *Path format example:*
```
  C:/Users/walldo/Desktop/fondos, C:/Users/walldo/Desktop/wallpapers2
```
*NOTE: do not use backslashes (\), instead use normal slashes (/)*


### Linux
After installing the dependencies, open a terminal session and run:
````
go install github.com/elias-gill/walldo-in-go@latest
````
You can now run "walldo-in-go" in your terminal. If you are using a Desktop environment, you can create a shortcut for the command.

#### Important
After running the application for the first time, you'll find a folder in `/home/your-user/.config/walldo`. Inside this directory there should be a `config.json` file; you can specify your wallpaper directories in here (Walldo supports recursive directories).

## Goals
- Be the replacement for Nitrogen or Feh on Windows Systems. 
- Be faster and lighter than similar apps.

### Comming features (maybe)
- A button to get a random wallpaper.
- More layouts like a list with captions.

## Mentions
- To [ktr0731](https://github.com/ktr07310). The fuzzy finder engine is his entire work.
- To [reujab](https://github.com/reujab/wallpaper). The library for changing wallapapers is a fork of his original module.
