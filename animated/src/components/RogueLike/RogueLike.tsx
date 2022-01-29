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

  const [playerXP, setPlayerXP] = useState(0)
  const playerXPRef = useRef(playerXP)

  const [playerXPNextLevel, setPlayerXPNextLevel] = useState(0)
  const playerXPNextLevelRef = useRef(playerXPNextLevel)

  const [playerXPCurrentLevel, setPlayerXPCurrentLevel] = useState(0)
  const playerXPCurrentLevelRef = useRef(playerXPCurrentLevel)

  const [playerHP, setPlayerHP] = useState(0)
  const playerHPRef = useRef(playerHP)

  const [playerMaxHP, setPlayerMaxHP] = useState(0)
  const playerMaxHPRef = useRef(playerMaxHP)

  useEffect(() => { return }, [playerLevel, playerXP, playerHP, playerMaxHP, playerXPNextLevel, playerXPCurrentLevel])

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

  const drawObject = (
    ctx: CanvasRenderingContext2D,
    sprite: Positions,
    tileset: string,
    px: number,
    py: number,
  ) => {
    const image = new Image()
    image.src = `${window.location.origin}/img/assets/rogue/sprites/${tileset}.png`

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
    if (playerHPRef.current !== p1.health) setPlayerHP(p1.health)
    if (playerMaxHPRef.current !== p1.maxHP) setPlayerMaxHP(p1.maxHP)
    if (playerHPRef.current !== p1.level) setPlayerHP(p1.health)
    if (
      playerXPRef.current !== p1.xp ||
      playerXPNextLevelRef.current !== p1.xpNextLevel ||
      playerXPCurrentLevelRef.current !== p1.xpCurrentLevel
    ) {
      setPlayerXP(p1.xp)
      setPlayerXPNextLevel(p1.xpNextLevel)
      setPlayerXPCurrentLevel(p1.xpCurrentLevel)
    }

    drawBackground(canvas, ctx, p1.positionX, p1.positionY)

    PLAYERS_DATA.forEach((player) => {
      const sprite = handleAnimations(player, now)
      const posX = (player.positionX + sprite.xOffset || 0) + CANVAS_WIDTH / 2
      const posY = (player.positionY + sprite.yOffset || 0) + CANVAS_HEIGHT / 2

      const px = (p1.positionX <= CANVAS_WIDTH / 2 ? posX - CANVAS_WIDTH / 2 : posX - p1.positionX)
      const py = (p1.positionY <= CANVAS_HEIGHT / 2 ? posY - CANVAS_HEIGHT / 2 : posY - p1.positionY)

      drawObject(ctx, sprite, player.sprite.tileSet, px, py)
      drawHealthbar(ctx, px, py - 1, (player.health / (player.maxHP ? player.maxHP : player.sprite.hp)) * 100, player.sprite.spriteWidth, 1)
    })
    ENEMIES_DATA?.forEach((enemy) => {
      const sprite = handleAnimations(enemy, now)
      const posX = (enemy.positionX + sprite.xOffset || 0) + CANVAS_WIDTH / 2
      const posY = (enemy.positionY + sprite.yOffset || 0) + CANVAS_HEIGHT / 2

      const px = (p1.positionX <= CANVAS_WIDTH / 2 ? posX - CANVAS_WIDTH / 2 : posX - p1.positionX)
      const py = (p1.positionY <= CANVAS_HEIGHT / 2 ? posY - CANVAS_HEIGHT / 2 : posY - p1.positionY)

      drawObject(ctx, sprite, enemy.sprite.tileSet, px, py)
      drawHealthbar(ctx, px, py - 1, (enemy.health / (enemy.maxHP ? enemy.maxHP : enemy.sprite.hp)) * 100, enemy.sprite.spriteWidth, 1)
    })
    DROPS_DATA?.forEach((drop) => {
      const posX = drop.positionX + CANVAS_WIDTH / 2
      const posY = drop.positionY + CANVAS_HEIGHT / 2
      const px = (p1.positionX <= CANVAS_WIDTH / 2 ? posX - CANVAS_WIDTH / 2 : posX - p1.positionX) + drop.sprite.xOffset
      const py = (p1.positionY <= CANVAS_HEIGHT / 2 ? posY - CANVAS_HEIGHT / 2 : posY - p1.positionY) + drop.sprite.yOffset
      drawObject(ctx, drop.sprite, drop.sprite.tileSet, px, py)
    })
    PROJECTILES_DATA?.forEach((projectile) => {
      const posX = projectile.positionX + CANVAS_WIDTH / 2
      const posY = projectile.positionY + CANVAS_HEIGHT / 2

      // TODO: consider end of map
      const px = (p1.positionX <= CANVAS_WIDTH / 2 ? posX - CANVAS_WIDTH / 2 : posX - p1.positionX) + projectile.sprite.xOffset
      const py = (p1.positionY <= CANVAS_HEIGHT / 2 ? posY - CANVAS_HEIGHT / 2 : posY - p1.positionY) + projectile.sprite.yOffset

      drawObject(ctx, projectile.sprite, projectile.sprite.tileSet, px, py)
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
      {gameState == "started" && <>
        <S.StatsList>
          {gameState === "started" && <li>level: {playerLevel}</li>}
        </S.StatsList>
        <S.ProgressContainer>
          <S.LifeLabel htmlFor="health">{playerHP}/{playerMaxHP}</S.LifeLabel>
          <S.Progress id="health" value={playerHP} max={playerMaxHP}></S.Progress>
        </S.ProgressContainer>
        <S.ProgressXPContainer>
          <S.ProgressXP value={playerXP - playerXPCurrentLevel} max={playerXPNextLevel}></S.ProgressXP>
        </S.ProgressXPContainer>
      </>}
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
