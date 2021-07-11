# thread-art

## Ressourcer
- http://artof01.com/vrellis/works/knit.html
- https://sim-on.github.io/2017/07/26/hula/
  - https://github.com/sim-on/aNewWayToKnit/blob/master/knit.py
  - https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
- https://www.youtube.com/watch?v=UsbBSttaJos
- https://www.nuget.org/packages/shortid
- https://www.nuget.org/packages/System.Drawing.Common

Relevante Microsoft Docs
- https://docs.microsoft.com/en-us/dotnet/core/tools/dotnet-publish
  - https://docs.microsoft.com/en-us/dotnet/core/deploying/single-file#publish-a-single-file-app---cli
- https://docs.microsoft.com/en-us/dotnet/core/rid-catalog
- https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/ref
- 

Billedet, der medfølger i mappen er taget af Megan Bagshaw og kan findes på følgende link: https://unsplash.com/photos/zYDISXBOWmA.

## Bygge instruktioner
Hent GitHub lageret og kør følgende i `threadArtApplication` mappen.
```
dotnet publish --self-contained true --runtime <RUNTIME_IDENTIFIER>
```
hvor `<RUNTIME_IDENTIFIER>` kommer an på styresystemet. Eksempler på sådanne er: `linux-x64`, `win-x64`, `osx-x64` osv. Mere information kan findes på .NET RID Catalog (https://docs.microsoft.com/en-us/dotnet/core/rid-catalog). 

Yderligere information om `dotnet publish` kan findes hos Microsoft Docs ved https://docs.microsoft.com/en-us/dotnet/core/tools/dotnet-publish.

## ToDo
- [ ] I afsnittet bygge instruktioner - tilføj krav bl.a. .NET runtime
- [ ] Forklar, at man kan bruge `dotnet run`. Så slipper man for at kompilere programmet.
