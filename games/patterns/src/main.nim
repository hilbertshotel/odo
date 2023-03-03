import dom
import strformat

import data
import procs

# BUTTONS
# ================================================================================

proc play() {.exportc.} =
    playSound("music")
    var tiles = sampleTiles(TILENUM, COLORNUM, COLORS)
    assembleTiles(tiles, TILENUM, COLORNUM, COLORS)


proc changeTiles(tile: Element) {.exportc.} =
    case $tile.innerHTML:
    of "3t": TILENUM = 6
    of "6t": TILENUM = 9
    of "9t": TILENUM = 12
    of "12t": TILENUM = 3
    tile.innerHTML = &"{TILENUM}t"


proc changeColors(tile: Element) {.exportc.} =
    case $tile.innerHTML:
    of "1c": COLORNUM = 2
    of "2c": COLORNUM = 3
    of "3c": COLORNUM = 4
    of "4c": COLORNUM = 1
    tile.innerHTML = &"{COLORNUM}c"


proc exit() {.exportc.} =
    window.location.href = "/"


# ON PAGE LOAD
# ================================================================================

proc onPageLoad() = loadEmptyBoxes()


onPageLoad()