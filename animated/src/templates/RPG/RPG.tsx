import { useEffect, useRef, useState } from "react"
import { CharacterSpriteMap } from "../../constants"
import { Player } from "../../types/rpg-types"
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


    const canvas = canvasRef.current
    if (!canvas) return
    const ctx = canvas.getContext("2d")
    if (!ctx) return

    canvas.width = CANVAS_WIDTH
    canvas.height = CANVAS_HEIGHT

    ctx.font = "40px serif"
    ctx.imageSmoothingEnabled = false

    webSocket.onmessage = (event) => {

      const data = JSON.parse(event.data)

      if (data.type === "broadcast") {

        ctx.clearRect(0, 0, canvas.width, canvas.height)
        data.players.forEach((player: Player) => {
          const image = new Image()
          image.onload = () => ctx.drawImage(
            image, player["position_x"], player["position_y"], CHARACTER_SIZE, CHARACTER_SIZE
          )
          // TODO: remove hardcoded image

          image.src = `${window.location.origin}${CharacterSpriteMap.get(player.type)}`
          ctx.fillText(player["username"], player["position_x"], player["position_y"] - 15)
        })

      }
    }


  }, [])

  return (
    <S.Canvas id="clock" width="150" height="150" ref={canvasRef}>
    </S.Canvas>
  )

}


export default RPG
