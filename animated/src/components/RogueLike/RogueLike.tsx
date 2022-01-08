import { MutableRefObject, useCallback, useEffect, useRef, useState } from "react"
import { Player, Warrior } from "../../types/rogue-types"
import * as S from "./RogueLike.styles"


const RogueLike = (): JSX.Element => {

  const [start, setStart] = useState(false)
  const [ws, setWs] = useState<WebSocket | null>(null)
  const canvasRef = useRef() as MutableRefObject<HTMLCanvasElement>

  let PLAYERS_DATA: Array<Player> = []

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
      }
    }
    animate()
  }

  const onStart = () => {
    setStart(true)
    connect()
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
    ctx.drawImage(background, dx, dy, dw, dh)
  }

  const drawPlayer = (
    ctx: CanvasRenderingContext2D,
    player: Player,
  ) => {
    const image = new Image()
    image.src = `${window.location.origin}/img/assets/rogue/sprites/${player.sprite.tileSet}.png`
    const sprite = player.sprite
    ctx.drawImage(
      image, sprite.spriteX, sprite.spriteY, sprite.spriteWidth, sprite.spriteHeight
    )

  }

  const animate = () => {
    const canvas = canvasRef.current
    if (!canvas) return
    const ctx = canvas.getContext("2d")
    if (!ctx) return
    ctx.imageSmoothingEnabled = false

    requestAnimationFrame(animate)

    const dx = 0
    const dy = 0
    drawBackground(canvas, ctx, dx, dy, canvas.width, canvas.height)
    PLAYERS_DATA.forEach((player) => drawPlayer(ctx, player))

  }


  return (
    <S.Container>
      <button onClick={onStart}>start</button>
      <S.Canvas width={8 * 15} height={8 * 10} ref={canvasRef}>
      </S.Canvas>
    </S.Container>
  )
}


export default RogueLike
