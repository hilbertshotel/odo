import jsconsole
import dom
import strutils

# PRINT
proc print*(input: any) =
    console.log(input)

# GET ID
proc getId*(id: string): Node =
    document.getElementById(id)

# GET CLASS
proc getClass*(name: string, index: int): Node =
    document.getElementsByClass(name)[index]

# INSERT
proc insert*(parent: Node, elements: varargs[Node]) =
    for element in elements:
        parent.appendChild(element)

# MAKE
proc make*(kind: string, properties: varargs[string]): Node =
    let element = document.createElement(kind)
    for property in properties:
        let p = split(property, '=')
        case p[0]:
        of "class": element.className = p[1]
        of "id": element.id = p[1]
        of "text": element.innerHTML = p[1]
    element

# GET AUDIO
proc getAudio*(id: string): EmbedElement =
    return getId(id).EmbedElement
