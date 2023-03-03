import random, dom, system, strformat
import lib, data


# GAME PROCEDURES
# ================================================================================


proc playNextColor(nextColor: cstring) =
    for color in COLORS:
        if color.hex == nextColor:
            getAudio(color.name).play()



proc action(tile, headerBoard: Node) =
    let
        currentHeader = headerBoard.children[0]
        currentColor = currentHeader.id

    if currentColor == tile.id:
        tile.style.visibility = "hidden"
        getAudio("grab").play()
        headerBoard.removeChild(currentHeader)

        let
            nextHeader = headerBoard.children[0]
            nextColor = nextHeader.id

        nextHeader.style.visibility = "visible"

        if nextColor != currentColor:
            playNextColor(nextColor)
            
    else:
        getAudio("fail").play()



proc sampleTileColors(): seq[Color] =
    randomize()
    for i in 1..36:
        add(result, sample(COLORS)) 



func sortByNextOccurence(tileColors: seq[Color]): seq[string] =
    let colors = tileColors[18..^1] & tileColors[0..17]
    var currentColor: string

    for c in colors:
        currentColor = c.hex
        if currentColor notin result:
            for color in colors:
                if color.hex == currentColor:
                    result.add(color.hex)



proc stackHeaders(headerBoard: Node, headerColors: seq[string]) =
    let 
        firstColor = headerColors[0]
        firstHeader = make("h1", "class=header", &"id={firstColor}", "text=GRABATILE")

    firstHeader.style.backgroundColor = firstColor
    headerBoard.insert(firstHeader)

    for color in headerColors[1..^1]:
        let header = make("h1", "class=header", &"id={color}", "text=GRABATILE")
        header.style.visibility = "hidden"
        header.style.backgroundColor = color
        headerBoard.insert(header)

    let lastHeader = make("h1", "class=header", "text=GRABATILE")
    lastHeader.style.visibility = "hidden"
    lastHeader.style.color = "#bebebe"
    headerBoard.insert(lastHeader)



proc fillTileMap(tileMap, headerBoard: Node, tileColors: seq[Color]) =
    for color in tileColors:
        
        closureScope:
            let tile = make("div", "class=tile", &"id={color.hex}")
            tile.style.backgroundImage = &"url(\"images/tiles/{color.name}_tile.png\")"
            tile.onclick = (proc (_: Event) = action(tile, headerBoard))
            tileMap.insert(tile)



# START
# ================================================================================

proc startButton() {.exportc.} =
    getAudio("music").play()

    let
        tileMap = getId("tileMap")
        headerBoard = getId("headerBoard")

    tileMap.innerHTML = ""
    headerBoard.innerHTML = ""

    let
        tileColors = sampleTileColors()
        headerColors = sortByNextOccurence(tileColors)

    stackHeaders(headerBoard, headerColors)
    fillTileMap(tileMap, headerBoard, tileColors)
    playNextColor(headerBoard.children[0].id)
