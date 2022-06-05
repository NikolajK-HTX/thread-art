# thread-art

## Brug af program
Man har følgende indstillinger, når man kører programmet.
```
./threadArtApplication -i <input_image> -n <number_of_pins> -s <outputimage_size> -t <number_of_threads> -m <minimum_difference> -o <output-image-path> -p <image_id>
```

Man kan på en p5.js sketch progressivt se, hvordan det ser ud, når trådene sættes på. Den er ikke helt færdig, men man kan se det [her](https://editor.p5js.org/NikolajK-HTX/sketches/q3gxY4B9H). Det virker ved at sætte punkterne fra tekstfilen ind i tekstfeltet.

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
- https://docs.microsoft.com/en-us/dotnet/api/system.guid?view=net-5.0
- https://www.nuget.org/packages/shortid
- https://www.nuget.org/packages/System.Drawing.Common

Relevante Microsoft Docs
- https://docs.microsoft.com/en-us/dotnet/core/tools/dotnet-publish
  - https://docs.microsoft.com/en-us/dotnet/core/deploying/single-file#publish-a-single-file-app---cli
- https://docs.microsoft.com/en-us/dotnet/core/rid-catalog
- https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/ref

GetPixel som sådan gør projektet [langsommere](https://imgur.com/a/WfjY8Gj) - det kan gøres hurtigere ([se mere](http://csharpexamples.com/fast-image-processing-c/)).

Billedet, der medfølger i mappen er taget af Megan Bagshaw og kan findes på følgende link: https://unsplash.com/photos/zYDISXBOWmA.

Upload af billeder kan ske med
 - https://www.dropzone.dev/js/
 - https://pqina.nl/filepond/

 - https://github.com/dropzone/dropzone
 - https://github.com/pqina/filepond

Lige umiddelbart synes jeg bedst om dropzone

## Bygge instruktioner

### Go version
Hent go fra [go.dev](https://go.dev/doc/install).

Programmet kompileres med
```
go build
```
og køres med ``./thread-art-go`` (Linux) eller ``.\thread-art-go.exe`` (Windows).

Man kan også nøjes med
```
go run .
```
Det kører programmet, men man får ikke en eksekverbar fil.

### Dotnet version
På linux kræves ``libgdiplus``, som kan installeres med ``sudo dnf install libgdiplus``. I stedet for ``dnf`` kan ``apt`` eller (formentlig) andre package managers bruges.

Hent GitHub lageret og kør følgende i `threadArtApplication` mappen.
```
dotnet publish --self-contained true --runtime <RUNTIME_IDENTIFIER>
```
hvor `<RUNTIME_IDENTIFIER>` kommer an på styresystemet. Eksempler på sådanne er: `linux-x64`, `win-x64`, `osx-x64` osv. Mere information kan findes på .NET RID Catalog (https://docs.microsoft.com/en-us/dotnet/core/rid-catalog). 

Yderligere information om `dotnet publish` kan findes hos Microsoft Docs ved https://docs.microsoft.com/en-us/dotnet/core/tools/dotnet-publish.

## ToDo i forhold til ReadME
- [ ] I afsnittet bygge instruktioner - tilføj krav bl.a. .NET runtime
- [ ] Forklar, at man kan bruge `dotnet run`. Så slipper man for at kompilere programmet.

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

For at `System.Drawing` virker på Linux skal `libgdiplus` være installeret, men desværre virker det ikke på `Linux Arm` eller `Linux Arm 64`, da det giver følgende fejl:
```
Unhandled exception. System.ArgumentException: Parameter is not valid.
   at System.Drawing.SafeNativeMethods.Gdip.CheckStatus(Int32 status)
   at System.Drawing.Bitmap..ctor(String filename, Boolean useIcm)
   at threadArtApplication.Program.Main(String[] args) in C:\Users\nikol\Documents\GitHub\thread-art\threadArtApplication\Program.cs:line 336
```

Mere information om `Halide` kan findes ved https://halide-lang.org/, men det virker ikke som om, det kan anvendes til dette projekt.
