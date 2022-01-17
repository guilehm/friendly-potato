import { Player, Sprite } from "../types/rogue-types"

interface AnimationTime {
  [key: string]: [number, boolean]
}
const lastAnimationMap: AnimationTime = {}

const deadSprite: Sprite = {
  name: "Ghost",
  tileSet: "sprites",
  spriteX: 37,
  spriteY: 56,
  spriteWidth: 6,
  spriteHeight: 7,
  hp: 0,
  xOffset: 1,
  yOffset: 1,
  animationPeriod: 0,
  damage: 0,
  attackRange: 0,
  animation: {
    spriteX: 0,
    spriteY: 0,
    spriteWidth: 0,
    spriteHeight: 0,
    xOffset: 0,
    yOffset: 0,
  }
}

export const handleAnimations = (player: Player, now: number) => {
  if (player.dead) return deadSprite

  // TODO: improve performance
  let sprite
  let animationData = lastAnimationMap[player.id.toString()]
  if (!animationData) {
    lastAnimationMap[player.id.toString()] = [now, false]
    animationData = [now, false]
  }
  const [lastAnimation, animated] = animationData

  if (animated) sprite = player.sprite.animation
  else sprite = player.sprite

  if (now > (lastAnimation + player.sprite.animationPeriod)) {
    lastAnimationMap[player.id.toString()] = [now, !animated]
  }
  return sprite
}
