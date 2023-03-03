type Color* = object
    name*: string
    hex*: string

type ColorsArray* = array[11, Color]

const COLORS*: ColorsArray = [
        Color(name: "white", hex: "#FFFFFF"),
        Color(name: "yellow", hex: "#FFdd00"),
        Color(name: "orange", hex: "#FF9100"),
        Color(name: "red", hex: "#FF4b1F"),
        Color(name: "brown", hex: "#9e5b0e"),
        Color(name: "green", hex: "#6e9900"),
        Color(name: "blue", hex: "#146de0"),
        Color(name: "purple", hex: "#a34aa3"),
        Color(name: "pink", hex: "#feafc9"),
        Color(name: "grey", hex: "#b6b6b6"),
        Color(name: "black", hex: "#1e1e1e"),
    ]
