import { render } from "@testing-library/react"
import { useEffect, useRef, useState } from "react"
import { WSMessage } from "../../types/ws-types"
import * as S from "./RPG.styles"

const CHARACTER_SIZE = 100
const CANVAS_WIDTH = 1000
const CANVAS_HEIGHT = 800


const RPG = (): JSX.Element => {

  const canvasRef = useRef<HTMLCanvasElement>()

  useEffect(() => {
    // TODO: remove hardcoded url
    const location = "ws://localhost:8080/ws/rpg/"
    const webSocket = new WebSocket(location)

    const message: WSMessage = {
      type: "game-join",
      // TODO: remove hardcoded username
      data: { "username": "guilehm" },
    }

    webSocket.onopen = () => {
      webSocket.send(JSON.stringify(message))
    }

    webSocket.onmessage = (event) => {
      const data = JSON.parse(event.data)
      if (data.type === "login-successful" || data.type === "update") {
        const data = JSON.parse(event.data).data

        const canvas = canvasRef.current
        if (!canvas) return
        const ctx = canvas.getContext("2d")
        if (!ctx) return

        const image = new Image()
        image.onload = () => ctx.drawImage(
          image, data["position_x"], data["position_y"], CHARACTER_SIZE, CHARACTER_SIZE
        )
        image.src = `${window.location.origin}/img/assets/characters/tile096.png`

        canvas.width = CANVAS_WIDTH
        canvas.height = CANVAS_HEIGHT

        ctx.font = "40px serif"
        ctx.fillText(data["username"], data["position_x"], data["position_y"] - 15)
        ctx.imageSmoothingEnabled = false
      }
    }


  }, [])

  return (
    <S.Canvas id="clock" width="150" height="150" ref={canvasRef}>
    </S.Canvas>
  )

}


export default RPG
