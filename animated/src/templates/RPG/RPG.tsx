import { ArrowBackIcon, ArrowDownIcon, ArrowForwardIcon, ArrowUpIcon } from "@chakra-ui/icons"
import { MutableRefObject, useEffect, useRef, useState } from "react"
import { ARROW_DOWN, ARROW_LEFT, ARROW_RIGHT, ARROW_UP, getCharacterSprite } from "../../constants"
import { Player } from "../../types/rpg-types"
import { WSMessage } from "../../types/ws-types"
import * as S from "./RPG.styles"


const CHARACTER_SIZE = 100 / 2
const CANVAS_WIDTH = 1000 / 2
const CANVAS_HEIGHT = 800 / 2


type GameStartInputs = {
  username: string
}

const schema = yup.object().shape({
  username: yup.string().min(3).max(15).required(),
})


const RPG = (): JSX.Element => {
  const [username, setUsername] = useState("")
  const [ws, setWs] = useState<WebSocket | null>(null)
  const canvasRef = useRef() as MutableRefObject<HTMLCanvasElement>


  const { register, handleSubmit, formState: { errors } } = useForm<GameStartInputs>({
    mode: "onBlur",
    resolver: yupResolver(schema),
  })

  const onStart = async (values: GameStartInputs) => {
    setUsername(values.username)

    const location = process.env.REACT_APP_WS_LOCATION || "ws://localhost:8080/ws/rpg/"
    const webSocket = new WebSocket(location)
    setWs(webSocket)

    const message: WSMessage = {
      type: "game-join",
      data: { "username": values.username },
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

    ctx.font = "20px serif"
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
          image.src = `${window.location.origin}${getCharacterSprite(player["type"], player["last_direction"])}`
          ctx.fillText(player["username"], player["position_x"], player["position_y"] - 10)
        })

      }
    }
  }


  const handleKeyDown = (key: string) => {
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
    <>
      <S.Container tabIndex={0} onKeyDown={(e) => handleKeyDown(e.key)}>
      </S.Container>
      <S.Canvas id="rpg" width="150" height="150" ref={canvasRef}>
      </S.Canvas>
      <S.KeysContainer>
        <ArrowUpIcon w={16} h={16} onClick={() => handleKeyDown(ARROW_UP)} />
        <br />
        <ArrowBackIcon w={16} h={16} onClick={() => handleKeyDown(ARROW_LEFT)} />
        <ArrowDownIcon w={16} h={16} onClick={() => handleKeyDown(ARROW_DOWN)} />
        <ArrowForwardIcon w={16} h={16} onClick={() => handleKeyDown(ARROW_RIGHT)} />
      </S.KeysContainer>
    </>
  )

}


export default RPG
