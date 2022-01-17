type CharacterType = string
type TileSet = string

export const Warrior: CharacterType = "warrior"
export const Background: TileSet = "background"
export const Characters: TileSet = "characters"

type Animation = {
  spriteX: number
  spriteY: number
  spriteWidth: number
  spriteHeight: number
  xOffset: number
  yOffset: number
}

export type Sprite = {
  name: CharacterType
  tileSet: TileSet
  spriteX: number
  spriteY: number
  spriteWidth: number
  spriteHeight: number
  hp: number
  // moveRange: number
  damage: number
  attackRange: number
  xOffset: number
  yOffset: number
  animationPeriod: number
  animation: Animation
}

export interface Positions {
  spriteX: number
  spriteY: number
  spriteWidth: number
  spriteHeight: number
  xOffset: number
  yOffset: number
}

type Position = {
  positionX: number
  positionY: number
}

export type Player = {
  id: number
  health: number
  positionX: number
  positionY: number
  sprite: Sprite
  lastPosition: Position
  dead: boolean

  // frontend only
  animation: boolean
  lastAnimationTime: number
  moving: boolean
  movingPosition: Position
  lastMovingTime: number
}
