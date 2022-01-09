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

  if (
    player.lastInteraction === true &&
    (player.lastPosition.positionX !== player.positionX) ||
    (player.lastPosition.positionY !== player.positionY)
  ) {

    if (!player.movingPosition) player.movingPosition = {
      positionX: player.lastPosition.positionX,
      positionY: player.lastPosition.positionY,
    }

    if (
      (player.movingPosition.positionX !== player.positionX) ||
      (player.movingPosition.positionY !== player.positionY)
    ) {
      if (!player.lastMovingTime) player.lastMovingTime = now
      player.moving = true
      if (now > (player.lastMovingTime + (player.sprite.animationPeriod / 40))) {
        player.lastMovingTime = now

        if (player.movingPosition.positionX < player.positionX) player.movingPosition.positionX += 1
        if (player.movingPosition.positionX > player.positionX) player.movingPosition.positionX -= 1
        if (player.movingPosition.positionY < player.positionY) player.movingPosition.positionY += 1
        if (player.movingPosition.positionY > player.positionY) player.movingPosition.positionY -= 1
      }
    } else {
      player.moving = false
    }
  }
  return sprite
}
