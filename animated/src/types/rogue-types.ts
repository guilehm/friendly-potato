type CharacterType = string
type TileSet = string

export const Warrior: CharacterType = "warrior"
export const Background: TileSet = "background"
export const Characters: TileSet = "characters"


type Sprite = {
  name: CharacterType
  tileSet: TileSet
  spriteX: 0
  spriteY: 0
  spriteWidth: number
  spriteHeight: number
  hp: number
  moveRange: number
  attackRange: number
  xOffset: number
  yOffset: number
}

export type Player = {
  health: number,
  positionX: number
  positionY: number
  sprite: Sprite
}

