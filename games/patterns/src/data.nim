# ONLY THE MAIN FILE HAS ACCESS TO THIS DATA

var
    TILENUM* = 3
    COLORNUM* = 1

type AllColors* = object
    orange*: string
    yellow*: string
    red*: string
    pink*: string
    nocolor*: string

const COLORS* = AllColors(
    orange: "#f09473",
    yellow: "#f5b82b",
    red: "#eb5952",
    pink: "#e95990",
    nocolor: "#f7b6b1",
)
