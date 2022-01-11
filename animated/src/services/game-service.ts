import { Player } from "../types/rogue-types"

export const handleAnimations = (player: Player, now: number) => {

  let sprite
  if (player.lastAnimationTime === undefined) player.lastAnimationTime = now
  if (player.animation) {
    sprite = player.sprite.animation
  } else {
    sprite = player.sprite
  }
  if (now > (player.lastAnimationTime + player.sprite.animationPeriod)) {
    player.animation = !player.animation
    player.lastAnimationTime = now
  }
  return sprite
}
