import { useEffect, useRef } from "react"
import * as S from "./RPG.styles"

const RPG = (): JSX.Element => {

  const canvasRef = useRef<HTMLCanvasElement>()

  useEffect(() => {


    const canvas = canvasRef.current
    const ctx = canvas.getContext("2d")

    const image = new Image()
    image.onload = () => ctx.drawImage(image, 50, 50, 8, 8)
    image.src = `${window.location.origin}/img/assets/characters/tile096.png`
    ctx.imageSmoothingEnabled = false
  }, [])


  return (
    <S.Canvas id="clock" width="150" height="150" ref={canvasRef}>
    </S.Canvas>
  )

}


export default RPG
