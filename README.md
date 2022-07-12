# Walldo-in-Go
A simple wallpaper changer written in Go with the main goal of being light and fast.
Mainly inspired by Nitrogen (more like a knockoff  ¯\_( ͡❛͜ʖ ͡❛)_/¯ ), Walldo searches for your images, lists them all 
and displays them on a single screen.
Change your wallpaper as fast as you want. Nothing more and nothing less, simple and functional.

#### A wallpaper changer app ? Why ?... 
As a person who likes to change his wallpaper very often, I was looking for an app that would allow me to do it in a 
simple way and as fast as possible. In Linux there are many applications that can do that, 
but in Windows most of them are very slow and bloated.

## Instalation
##### Dependendies 
Walldo is written in pure Go so of course you have to install [Golang](https://go.dev/doc/install).

### Windows
After you have installed Go, simply open a new terminal (cmd) and run:
````
go install -ldflags -H=windowsgui github.com/elias-gill/walldo-in-go@latest
````
Now Walldo must be available in your search bar (```Win```+```S```) as "waldo-in-go".
You can create a direct access and change the icon if you want.

### Linux
Note: Linux is not the main target, I recommend looking for Feh or Nitrogen.
````
go install github.com/elias-gill/walldo-in-go@latest
````
Now you can run "walldo-in-go" on your terminal. If you are using some Desktop Enviroment you can create a direct 
acces for the command.

## Goals
- Be the replace for Nitrogen or Feh on Windows Systems. 
- Be faster and lighter than other similar apps.

### Comming features (maybe)
- A button to get a random wallpaper from the internet.
