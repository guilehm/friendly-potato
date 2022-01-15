import { MutableRefObject, useRef, useState } from "react"
import { ARROW_DOWN, ARROW_LEFT, ARROW_RIGHT, ARROW_UP, INTERACTION_COOLDOWN } from "../../constants"
import { handleAnimations } from "../../services/game-service"
import { Player, Positions, Sprite, Warrior } from "../../types/rogue-types"
import { WSMessage } from "../../types/ws-types"
import * as S from "./RogueLike.styles"

const CANVAS_WIDTH = 8 * 16
const CANVAS_HEIGHT = 8 * 10

const RogueLike = (): JSX.Element => {

  const [ws, setWs] = useState<WebSocket | null>(null)
  const canvasRef = useRef() as MutableRefObject<HTMLCanvasElement>

  let PLAYERS_DATA: Array<Player> = []
  let ENEMIES_DATA: Array<Player> = []
  let LAST_INTERACTION = Date.now()

  const connect = () => {
    const location = process.env.REACT_APP_ROGUE_WS_LOCATION || "ws://localhost:8080/ws/rogue/"
    const webSocket = new WebSocket(location)
    setWs(webSocket)

    webSocket.onopen = () => {
      webSocket.send(JSON.stringify({
        type: "user-joins",
        data: { "sprite": Warrior },
      }))
    }

    webSocket.onmessage = (event) => {
      const data = JSON.parse(event.data)
      if (data.type === "broadcast") {
        PLAYERS_DATA = data.players
        ENEMIES_DATA = data.enemies
      }
    }
    animate()
  }

  const drawBackground = (
    canvas: HTMLCanvasElement,
    ctx: CanvasRenderingContext2D,
    dx: number,
    dy: number,
    dw: number,
    dh: number,
  ) => {
    ctx.clearRect(0, 0, canvas.width, canvas.height)
    const background = new Image()
    background.src = `${window.location.origin}/img/assets/rogue/sprites/background.png`

    ctx.drawImage(
      background,
      // TODO: set limit wor x+ and y+
      (dx + dw) >= 0 ? dx + dw : 0,
      (dy + dh) >= 0 ? dy + dh : 0,
      canvas.width,
      canvas.height,
      dx,
      dy,
      canvas.width,
      canvas.height,
    )

  }

  const drawPlayer = (
    ctx: CanvasRenderingContext2D,
    player: Player,
    sprite: Positions,
    dw: number,
    dh: number,
  ) => {
    const image = new Image()
    image.src = `${window.location.origin}/img/assets/rogue/sprites/${player.sprite.tileSet}.png`
    ctx.drawImage(
      image,
      sprite.spriteX,
      sprite.spriteY,
      sprite.spriteWidth,
      sprite.spriteHeight,
      // TODO: simplify this condition
      dw >= 0 ? ((player.positionX + sprite.xOffset || 0) - dw - player.sprite.spriteWidth - player.sprite.xOffset) + CANVAS_WIDTH / 2 : ((player.positionX + sprite.xOffset || 0) - dw - player.sprite.spriteWidth - player.sprite.xOffset) + CANVAS_WIDTH / 2 + dw,
      dh >= 0 ? ((player.positionY + sprite.yOffset || 0) - dh - player.sprite.spriteHeight - player.sprite.yOffset) + CANVAS_HEIGHT / 2 : ((player.positionY + sprite.yOffset || 0) - dh - player.sprite.spriteHeight - player.sprite.yOffset) + CANVAS_HEIGHT / 2 + dh,
      sprite.spriteWidth,
      sprite.spriteHeight,
    )

  }

  const animate = () => {
    const canvas = canvasRef.current
    if (!canvas) return
    const ctx = canvas.getContext("2d")
    if (!ctx) return
    ctx.imageSmoothingEnabled = false

    requestAnimationFrame(animate)

    // TODO: hardcoded from now. Send and retrieve ID from server.
    const p1 = PLAYERS_DATA.find(p => p.sprite.name === Warrior)
    if (!p1) return

    const dx = 0
    const dy = 0
    drawBackground(canvas, ctx, dx, dy, p1.positionX, p1.positionY)

    const now = Date.now()
    PLAYERS_DATA.forEach((player) => {
      const sprite = handleAnimations(player, now)
      drawPlayer(ctx, player, sprite, p1.positionX, p1.positionY)
    })
    ENEMIES_DATA.forEach((enemy) => {
      const sprite = handleAnimations(enemy, now)
      drawPlayer(ctx, enemy, sprite, p1.positionX, p1.positionY)
    })

  }

  const handleKeyDown = (key: string) => {
    const now = Date.now()
    if (!(now > LAST_INTERACTION + INTERACTION_COOLDOWN)) return
    LAST_INTERACTION = now
    const validKeys = [
      ARROW_LEFT,
      ARROW_UP,
      ARROW_RIGHT,
      ARROW_DOWN,
    ]

    if (!validKeys.includes(key)) return
    const msg: WSMessage = {
      type: "key-down",
      data: key,
    }
    ws && ws.send(JSON.stringify(msg))
  }


  return (
    <S.Container>
      <button onClick={connect}>start</button>
      <S.Canvas
        tabIndex={0}
        width={CANVAS_WIDTH}
        height={CANVAS_HEIGHT}
        ref={canvasRef}
        onKeyDown={(e) => handleKeyDown(e.key)}>
      </S.Canvas>
    </S.Container >
  )
}


export default RogueLike
