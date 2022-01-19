import { ArrowBackIcon, ArrowDownIcon, ArrowForwardIcon, ArrowUpIcon } from "@chakra-ui/icons"
import { MutableRefObject, useRef, useState } from "react"
import { ARROW_DOWN, ARROW_LEFT, ARROW_RIGHT, ARROW_UP, INTERACTION_COOLDOWN, KEY_A, KEY_D, KEY_S, KEY_W } from "../../constants"
import { drawHealthbar } from "../../helpers"
import { handleAnimations } from "../../services/game-service"
import { Drop, DropSprite, Player, Positions, Warrior } from "../../types/rogue-types"
import { WSMessage } from "../../types/ws-types"
import * as S from "./RogueLike.styles"
const CANVAS_WIDTH = 8 * 16
const CANVAS_HEIGHT = 8 * 10

const RogueLike = (): JSX.Element => {

  const [ws, setWs] = useState<WebSocket | null>(null)
  const canvasRef = useRef() as MutableRefObject<HTMLCanvasElement>
  const [gameState, setGameState] = useState("waiting")

  let PLAYERS_DATA: Array<Player> = []
  let ENEMIES_DATA: Array<Player> = []
  let DROPS_DATA: Array<Drop> = []
  let LAST_INTERACTION = Date.now()

  const connect = () => {
    const location = process.env.REACT_APP_ROGUE_WS_LOCATION || "ws://localhost:8080/ws/rogue/"
    const webSocket = new WebSocket(location)
    setWs(webSocket)
    // TODO: generate a better ID
    const userId = Date.now()

    webSocket.onopen = () => {
      webSocket.send(JSON.stringify({
        type: "user-joins",
        data: { "sprite": Warrior, "id": userId },
      }))
    }

    webSocket.onmessage = (event) => {
      const data = JSON.parse(event.data)
      if (data.type === "broadcast") {
        PLAYERS_DATA = data.players
        ENEMIES_DATA = data.enemies
        DROPS_DATA = data.drops
      }
    }
    animate(userId)
    setGameState("started")
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
      // TODO: set limit for x+ and y+
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

  const drawDrop = (
    ctx: CanvasRenderingContext2D,
    drop: Drop,
    sprite: DropSprite,
    dw: number,
    dh: number,
  ) => {
    const image = new Image()
    image.src = `${window.location.origin}/img/assets/rogue/sprites/${drop.sprite.tileSet}.png`

    const posX = drop.positionX - drop.sprite.spriteWidth - drop.sprite.xOffset + CANVAS_WIDTH / 2
    const posY = drop.positionY - drop.sprite.spriteHeight - drop.sprite.yOffset + CANVAS_HEIGHT / 2

    // TODO: consider end of map
    const px = dw >= 0 ? posX - dw : posX
    const py = dh >= 0 ? posY - dh : posY

    ctx.drawImage(
      image,
      sprite.spriteX,
      sprite.spriteY,
      sprite.spriteWidth,
      sprite.spriteHeight,
      px,
      py,
      sprite.spriteWidth,
      sprite.spriteHeight,
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

    const posX = (
      (player.positionX + sprite.xOffset || 0) - player.sprite.spriteWidth - player.sprite.xOffset
    ) + CANVAS_WIDTH / 2
    const posY = (
      (player.positionY + sprite.yOffset || 0) - player.sprite.spriteHeight - player.sprite.yOffset
    ) + CANVAS_HEIGHT / 2

    // TODO: consider end of map
    const px = dw >= 0 ? posX - dw : posX
    const py = dh >= 0 ? posY - dh : posY
    ctx.drawImage(
      image,
      sprite.spriteX,
      sprite.spriteY,
      sprite.spriteWidth,
      sprite.spriteHeight,
      px,
      py,
      sprite.spriteWidth,
      sprite.spriteHeight,
    )

    drawHealthbar(ctx, px, py - 1, (player.health / player.sprite.hp) * 100, player.sprite.spriteWidth, 1)

  }

  const animate = (userId: number) => {
    const canvas = canvasRef.current
    if (!canvas) return
    const ctx = canvas.getContext("2d")
    if (!ctx) return
    ctx.imageSmoothingEnabled = false

    requestAnimationFrame(() => animate(userId))

    const p1 = PLAYERS_DATA.find(p => p.id === userId)
    if (!p1) return

    const dx = 0
    const dy = 0
    drawBackground(canvas, ctx, dx, dy, p1.positionX, p1.positionY)

    const now = Date.now()
    PLAYERS_DATA.forEach((player) => {
      const sprite = handleAnimations(player, now)
      drawPlayer(ctx, player, sprite, p1.positionX, p1.positionY)
    })
    ENEMIES_DATA?.forEach((enemy) => {
      const sprite = handleAnimations(enemy, now)
      drawPlayer(ctx, enemy, sprite, p1.positionX, p1.positionY)
    })
    DROPS_DATA?.forEach((drop) => {
      drawDrop(ctx, drop, drop.sprite, p1.positionX, p1.positionY)
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
      KEY_A,
      KEY_W,
      KEY_D,
      KEY_S,
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
      {
        gameState == "started" &&
        <S.ArrowContainer>
          <ArrowUpIcon onClick={() => handleKeyDown(ARROW_UP)} w={16} h={16} />
          <br />
          <ArrowBackIcon onClick={() => handleKeyDown(ARROW_LEFT)} w={16} h={16} />
          <ArrowDownIcon onClick={() => handleKeyDown(ARROW_DOWN)} w={16} h={16} />
          <ArrowForwardIcon onClick={() => handleKeyDown(ARROW_RIGHT)} w={16} h={16} />
        </S.ArrowContainer>
      }
    </S.Container >
  )
}


export default RogueLike
