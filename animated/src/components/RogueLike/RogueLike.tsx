import { MutableRefObject, useCallback, useEffect, useRef, useState } from "react"
import { Warrior } from "../../types/rogue-types"
import * as S from "./RogueLike.styles"


const RogueLike = (): JSX.Element => {

  const [ws, setWs] = useState<WebSocket | null>(null)
  const canvasRef = useRef() as MutableRefObject<HTMLCanvasElement>

  function drawBackground(
    canvas: HTMLCanvasElement,
    ctx: CanvasRenderingContext2D,
    dx: number,
    dy: number,
    dw: number,
    dh: number,
  ) {

    ctx.clearRect(0, 0, canvas.width, canvas.height)
    const background = new Image()
    background.src = `${window.location.origin}/img/assets/rogue/sprites/background.png`
    ctx.drawImage(background, dx, dy, dw, dh)

  }

  const animate = useCallback(() => {
    const canvas = canvasRef.current
    if (!canvas) return
    const ctx = canvas.getContext("2d")
    if (!ctx) return
    ctx.imageSmoothingEnabled = false

    requestAnimationFrame(animate)

    const dx = 0
    const dy = 0
    drawBackground(canvas, ctx, dx, dy, canvas.width, canvas.height)

  }, [])

  useEffect(() => {
    const location = process.env.REACT_APP_ROGUE_WS_LOCATION || "ws://localhost:8080/ws/rogue/"
    const webSocket = new WebSocket(location)
    setWs(webSocket)

    webSocket.onopen = () => {
      webSocket.send(JSON.stringify({
        type: "user-joins",
        // TODO: sprite should not be hardcoded
        data: { "sprite": Warrior },
      }))
    }
    animate()
  }, [animate])


  return (
    <S.Container>
      <S.Canvas width={720} height={480} ref={canvasRef}>
      </S.Canvas>
    </S.Container>
  )
}


export default RogueLike
