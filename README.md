# thread-art

## Brug af program
Hjælp om programmet fås ved
```
./thread-art-go -h
```

Man kan på en p5.js sketch progressivt se, hvordan det ser ud, når trådene sættes på. Den er ikke helt færdig, men man kan se det [her](https://editor.p5js.org/NikolajK-HTX/sketches/q3gxY4B9H). Det virker ved at sætte punkterne fra tekstfilen ind i tekstfeltet.
- En anden udgave: https://editor.p5js.org/NikolajK-HTX/sketches/8PC0UoFKX

## Ressourcer
Selve algoritmen
- http://artof01.com/vrellis/works/knit.html
- https://sim-on.github.io/2017/07/26/hula/
  - https://github.com/sim-on/aNewWayToKnit/blob/master/knit.py
  - https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
- https://www.youtube.com/watch?v=UsbBSttaJos
- https://eheitzresearch.wordpress.com/implementation-and-realization-of-petros-vrellis-knitting/
- https://github.com/alyyousuf7/Weaver
- https://github.com/i-make-robots/weaving_algorithm
- https://halfmonty.github.io/StringArtGenerator/

ID til specifikke billede
- https://github.com/dchest/uniuri

Billedet, der medfølger i mappen er taget af [Megan Bagshaw](https://unsplash.com/@megbagshaw) og kan findes på følgende link: https://unsplash.com/photos/zYDISXBOWmA ([web.archive.org](https://web.archive.org/web/20201203024840/https://unsplash.com/photos/zYDISXBOWmA) og [direkte link](https://images.unsplash.com/photo-1592124549776-a7f0cc973b24?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=1080&fit=max&ixid=eyJhcHBfaWQiOjEyMDd9)).

Upload af billeder håndteres med
- https://pqina.nl/filepond/
- https://github.com/pqina/svelte-filepond
- Alternativt https://www.dropzone.dev/js/

## Bygge instruktioner
Hent go fra [go.dev](https://go.dev/doc/install).

Programmet kompileres med
```
go build
```
og køres med `./thread-art-go` (Linux) eller `.\thread-art-go.exe` (Windows).

Man kan også nøjes med
```
go run .
```
Det kører programmet, men man får ikke en eksekverbar fil.

## Bibliotek ydelse
Ydelsen af programmet kommer selvfølgelig an på processoren ud over det anvendte bibliotek. Windows maskinen kører med en `i5-6600k@3.5GHz`, Linux med en `i5-8250u@3.4GHz` og Linux Arm med `Cortex-A53@1.4GHz` (der er tale om en Raspberry Pi 3 B+).

Bibliotek     | Windows | Linux | Linux Arm
--------------|---------|-------|----------
ImageSharp    | 4,6s    | 5,49s | 38,24s
System.Drawing| 0,67s   | 2,22s | NA
SkiaSharp     | tbd     | tbd   | tbd
Halide        | tbd     | tbd   | tbd
Go stdlib     | 0,53s   | 0,51s | tbd

Go stdlib er kørt med andre RAM på Windows. "Linux" har samme hardware, men kører en nyere linux kernel.

På linux kræves `libgdiplus`, som kan installeres med `sudo dnf install libgdiplus`. I stedet for `dnf` kan `apt` eller (formentlig) andre package managers bruges. Det virker dog desværre på `Linux Arm` eller `Linux Arm 64` og giver følgende fejl:
```
Unhandled exception. System.ArgumentException: Parameter is not valid.
   at System.Drawing.SafeNativeMethods.Gdip.CheckStatus(Int32 status)
   at System.Drawing.Bitmap..ctor(String filename, Boolean useIcm)
   at threadArtApplication.Program.Main(String[] args) in C:\Users\nikol\Documents\GitHub\thread-art\threadArtApplication\Program.cs:line 336
```

## Tak til
- https://pqina.nl/filepond/
- https://github.com/dchest/uniuri
- https://github.com/sim-on/aNewWayToKnit
- https://github.com/fputs/bresenham
- https://unsplash.com/@megbagshaw
