import { useEffect, useRef } from "react"
import * as S from "./RPG.styles"

const RPG = (): JSX.Element => {

  const canvasRef = useRef<HTMLCanvasElement>()

  useEffect(() => {

    const canvas = canvasRef.current
    if (!canvas) return
    const ctx = canvas.getContext("2d")
    if (!ctx) return

    const image = new Image()
    image.onload = () => ctx.drawImage(image, 50, 50, 80, 80)
    image.src = `${window.location.origin}/img/assets/characters/tile096.png`

    canvas.width = 1000
    canvas.height = 500

    ctx.font = "40px serif"
    ctx.fillText("gui", 50, 40)

    ctx.imageSmoothingEnabled = false
  }, [])


  return (
    <S.Canvas id="clock" width="150" height="150" ref={canvasRef}>
    </S.Canvas>
  )

}


export default RPG
