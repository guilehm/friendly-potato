import { Player } from "../types/rogue-types"

interface AnimationTime {
  [key: string]: [number, boolean]
}
const lastAnimationMap: AnimationTime = {}


export const handleAnimations = (player: Player, now: number) => {
  // TODO: improve performance
  let sprite
  const animationData = lastAnimationMap[player.id.toString()]
  if (!animationData) lastAnimationMap[player.id.toString()] = [now, false]
  const [lastAnimation, animated] = animationData
  if (animated) {
    sprite = player.sprite.animation
  } else {
    sprite = player.sprite
  }
  if (now > (lastAnimation + player.sprite.animationPeriod)) {
    lastAnimationMap[player.id.toString()] = [now, !animated]
  }
  return sprite
}
