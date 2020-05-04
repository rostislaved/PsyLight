# PsyLight
Software to control Adalight LED strip.

I used to use [AmbiBox](http://www.ambibox.ru/ru/index.php/%D0%97%D0%B0%D0%B3%D0%BB%D0%B0%D0%B2%D0%BD%D0%B0%D1%8F_%D1%81%D1%82%D1%80%D0%B0%D0%BD%D0%B8%D1%86%D0%B0) on windows to control adalight, but it's not cross-platform, so I wrote this.

## Features:
* Capture screen and change LEDs color (Examples on pictures on the bottom of the page)
* Up to 100 fps on 2560x1440 screen with ~20% CPU (i7 6700k) load
* Offset from the top and bottom of the screen to watch movies with comfort
* Offset from left and right sides (disabled by default)
* Configuration from file (-c flag to set path. Current directory by default)
* Adjustable heigth of horizontal zones and width of vertical
* -v - verbose mod to print FPS

It probably never will support anything other than 21x36 LED, as that's all I have, but I left these numbers to be changes in config file, maybe it will work out of the box. But one thing which now is absolutelly not possible to do now - is to change point in which PsyLight suposes your LED strip to start and what direction it goes (clockwise or counterclockwise).

## Installation
* Just run psylight with sudo (sudo needed to connect to usb port)
* Or change code, compile and run your version


## Photo examples

<img
src="https://github.com/Rostislaved/PsyLight/blob/master/Photo%20examples/1.JPG"
raw=true
alt="Subject Pronouns"
/>


<img
src="https://github.com/Rostislaved/PsyLight/blob/master/Photo%20examples/2.JPG"
raw=true
alt="Subject Pronouns"
/>


<img
src="https://github.com/Rostislaved/PsyLight/blob/master/Photo%20examples/3.JPG"
raw=true
alt="Subject Pronouns"
/>

<img
src="https://github.com/Rostislaved/PsyLight/blob/master/Photo%20examples/4.JPG"
raw=true
alt="Subject Pronouns"
/>

<img
src="https://github.com/Rostislaved/PsyLight/blob/master/Photo%20examples/5.JPG"
raw=true
alt="Subject Pronouns"
/>
