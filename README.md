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
##### Dependendies 
Walldo is written in pure Go so of course you have to install [Golang](https://go.dev/doc/install).

### Windows
After you have installed Go, simply open a new terminal (cmd) and run:
````
go install -ldflags -H=windowsgui github.com/elias-gill/walldo-in-go@latest
````
Now Walldo must be available in your search bar (```Win```+```S```) as "waldo-in-go".
You can create a direct access and change the icon if you want.

#### Important
After the first run you would find a folder in C:/Users/<your user>/AppData/Local/walldo , there is a file called config.json, in that file you have to put your backgrounds folders (if a folder contains other folders you only have to put the main folder, Walldo will search through all of them). If you have more than one main folder you can put them to.

**Example: **

{
  "Path": ["C:/Users/walldo/Desktop/fondos", "C:/Users\walldo/Desktop/wallpapers2"]
}

PLEASE DON'T put "\" instead put "/" as a separator.
You can simply reload Walldo with the reload button that is on the bottom right.

### Linux
For Linux is the same that is on Windows.
````
go install github.com/elias-gill/walldo-in-go@latest
````
Now you can run "walldo-in-go" on your terminal. If you are using some Desktop Enviroment you can create a direct 
acces for the command.
The same rule aplies here, after the first run you have to add the folders into the config.json file which is on /home/<your user>/.config/walldo.

## Goals
- Be the replace for Nitrogen or Feh on Windows Systems. 
- Be faster and lighter than other similar apps.

### Comming features (maybe)
- A button to get a random wallpaper.
- More layouts like a list with captions.

## Mentions
To [ktr0731](https://github.com/ktr07310). The fuzzy finder engine is his entire work.
To [reujab](https://github.com/reujab/wallpaper). The library for changing wallapapers is a fork of his original module.
