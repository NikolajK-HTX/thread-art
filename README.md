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

Til webserver del af projekt
- https://www.npmjs.com/package/sqlite3
- https://docs.microsoft.com/en-us/dotnet/standard/data/sqlite/?tabs=netcore-cli
- https://www.npmjs.com/package/nanoid
- https://docs.microsoft.com/en-us/aspnet/core/host-and-deploy/linux-nginx?view=aspnetcore-5.0
- https://certbot.eff.org/
- https://www.raspberrypi.org/documentation/

Billedet, der medfølger i mappen er taget af Megan Bagshaw og kan findes på følgende link: https://unsplash.com/photos/zYDISXBOWmA.

## Bygge instruktioner
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
