export const getCharacterSprite = (character: string, direction: string) => {
  let d = ""
  if (direction === "ArrowLeft") {
    d = "_b"
  }
  return `/img/assets/characters/${character}${d}.png`
}


export const ARROW_LEFT = "ArrowLeft"
export const ARROW_UP = "ArrowUp"
export const ARROW_RIGHT = "ArrowRight"
export const ARROW_DOWN = "ArrowDown"
