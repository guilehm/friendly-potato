import { ArrowBackIcon, ArrowDownIcon, ArrowForwardIcon, ArrowUpIcon, SmallCloseIcon } from "@chakra-ui/icons"
import { MutableRefObject, useEffect, useRef, useState } from "react"
import { ARROW_DOWN, ARROW_LEFT, ARROW_RIGHT, ARROW_UP, INTERACTION_COOLDOWN, KEY_A, KEY_D, KEY_S, KEY_SPACE, KEY_W } from "../../constants"
import { drawHealthbar } from "../../helpers"
import { handleAnimations } from "../../services/game-service"
import { Drop, DropSprite, Player, Positions, Projectile, ProjectileSprite, Warrior } from "../../types/rogue-types"
import { WSMessage } from "../../types/ws-types"
import * as S from "./RogueLike.styles"

const CANVAS_WIDTH = 8 * 16
const CANVAS_HEIGHT = 8 * 10
const FPS = 60

const RogueLike = (): JSX.Element => {

  const [ws, setWs] = useState<WebSocket | null>(null)
  const canvasRef = useRef() as MutableRefObject<HTMLCanvasElement>
  const [gameState, setGameState] = useState("waiting")

  const [playerLevel, setPlayerLevel] = useState(0)
  const playerLevelRef = useRef(playerLevel)

  useEffect(() => { return }, [playerLevel])

  let PLAYERS_DATA: Array<Player> = []
  let ENEMIES_DATA: Array<Player> = []
  let DROPS_DATA: Array<Drop> = []
  let PROJECTILES_DATA: Array<Projectile> = []
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
        PROJECTILES_DATA = data.projectiles
      }
    }

    const fpsInterval = 1000 / FPS
    const then = Date.now()
    const startTime = then
    animate(userId, fpsInterval, startTime, then)
    setGameState("started")
  }

  const drawBackground = (
    canvas: HTMLCanvasElement,
    ctx: CanvasRenderingContext2D,
    dw: number,
    dh: number,
  ) => {
    ctx.clearRect(0, 0, canvas.width, canvas.height)
    const background = new Image()
    background.src = `${window.location.origin}/img/assets/rogue/sprites/background.png`

    ctx.drawImage(
      background,
      // TODO: set limit for x+ and y+
      dw >= CANVAS_WIDTH / 2 ? dw - CANVAS_WIDTH / 2 : 0,
      dh >= CANVAS_HEIGHT / 2 ? dh - CANVAS_HEIGHT / 2 : 0,
      canvas.width,
      canvas.height,
      0,
      0,
      canvas.width,
      canvas.height,
    )

  }

  const drawProjectile = (
    ctx: CanvasRenderingContext2D,
    projectile: Projectile,
    sprite: ProjectileSprite,
    dw: number,
    dh: number,
  ) => {
    const image = new Image()
    image.src = `${window.location.origin}/img/assets/rogue/sprites/${projectile.sprite.tileSet}.png`

    const posX = projectile.positionX + CANVAS_WIDTH / 2
    const posY = projectile.positionY + CANVAS_HEIGHT / 2

    // TODO: consider end of map
    const px = (dw <= CANVAS_WIDTH / 2 ? posX - CANVAS_WIDTH / 2 : posX - dw) + projectile.sprite.xOffset
    const py = (dh <= CANVAS_HEIGHT / 2 ? posY - CANVAS_HEIGHT / 2 : posY - dh) + projectile.sprite.yOffset

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

  const drawDrop = (
    ctx: CanvasRenderingContext2D,
    drop: Drop,
    sprite: DropSprite,
    dw: number,
    dh: number,
  ) => {
    const image = new Image()
    image.src = `${window.location.origin}/img/assets/rogue/sprites/${drop.sprite.tileSet}.png`

    const posX = drop.positionX + CANVAS_WIDTH / 2
    const posY = drop.positionY + CANVAS_HEIGHT / 2

    // TODO: consider end of map
    const px = (dw <= CANVAS_WIDTH / 2 ? posX - CANVAS_WIDTH / 2 : posX - dw) + drop.sprite.xOffset
    const py = (dh <= CANVAS_HEIGHT / 2 ? posY - CANVAS_HEIGHT / 2 : posY - dh) + drop.sprite.yOffset

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
    const posX = (player.positionX + sprite.xOffset || 0) + CANVAS_WIDTH / 2
    const posY = (player.positionY + sprite.yOffset || 0) + CANVAS_HEIGHT / 2

    // TODO: consider end of map
    const px = (dw <= CANVAS_WIDTH / 2 ? posX - CANVAS_WIDTH / 2 : posX - dw)
    const py = (dh <= CANVAS_HEIGHT / 2 ? posY - CANVAS_HEIGHT / 2 : posY - dh)
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

  const animate = (userId: number, fpsInterval: number, startTime: number, then: number) => {
    const canvas = canvasRef.current
    if (!canvas) return
    const ctx = canvas.getContext("2d")
    if (!ctx) return
    ctx.imageSmoothingEnabled = false


    const now = Date.now()
    requestAnimationFrame(() => animate(userId, fpsInterval, startTime, then))
    const elapsed = now - then
    if (elapsed < fpsInterval) return
    then = now - (elapsed % fpsInterval)

    const p1 = PLAYERS_DATA.find(p => p.id === userId)
    if (!p1) return
    if (playerLevelRef.current !== p1.level) setPlayerLevel(p1.level)

    drawBackground(canvas, ctx, p1.positionX, p1.positionY)

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
    PROJECTILES_DATA?.forEach((projectile) => {
      drawProjectile(ctx, projectile, projectile.sprite, p1.positionX, p1.positionY)
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
      KEY_SPACE,
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
      {gameState === "waiting" && <button onClick={connect}>start</button>}
      <span>level: {playerLevel}</span>
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
          <SmallCloseIcon onClick={() => handleKeyDown(KEY_SPACE)} w={16} h={16} />
          <ArrowForwardIcon onClick={() => handleKeyDown(ARROW_RIGHT)} w={16} h={16} />
          <br />
          <ArrowDownIcon onClick={() => handleKeyDown(ARROW_DOWN)} w={16} h={16} />
        </S.ArrowContainer>
      }
    </S.Container >
  )
}


export default RogueLike
