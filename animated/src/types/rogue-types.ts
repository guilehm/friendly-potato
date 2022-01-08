type CharacterType = string
type TileSet = string

export const Warrior: CharacterType = "warrior"
export const Background: TileSet = "background"
export const Characters: TileSet = "characters"


type Sprite = {
  name: CharacterType
  attackRange: number
  hp: number
  moveRange: number
  spriteX: 0
  spriteY: 0
  spriteHeight: number
  spriteWidth: number
  tileSet: TileSet
}

export type Player = {
  health: number,
  positionX: number
  positionY: number
  sprite: Sprite
}

