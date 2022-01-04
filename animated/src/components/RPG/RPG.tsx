import { ArrowBackIcon, ArrowDownIcon, ArrowForwardIcon, ArrowUpIcon } from "@chakra-ui/icons"
import {
  Button,
  FormControl,
  FormErrorMessage,
  FormHelperText,
  FormLabel,
  HStack,
  Input,
  Stack
} from "@chakra-ui/react"
import { yupResolver } from "@hookform/resolvers/yup"
import { MutableRefObject, useRef, useState } from "react"
import { useForm } from "react-hook-form"
import * as yup from "yup"
import { ARROW_DOWN, ARROW_LEFT, ARROW_RIGHT, ARROW_UP } from "../../constants"
import { getMovingCharacterSprite } from "../../helpers"
import { Player } from "../../types/rpg-types"
import { WSMessage } from "../../types/ws-types"
import * as S from "./RPG.styles"


const CHARACTER_SIZE = 100 / 2
const CANVAS_WIDTH = 1000 / 2
const CANVAS_HEIGHT = 800 / 2
const FPS = 40


type GameStartInputs = {
  username: string
}

const schema = yup.object().shape({
  username: yup.string().min(3).max(15).required(),
})

const RPG = (): JSX.Element => {
  const [username, setUsername] = useState("")
  const [playerCount, setPlayerCount] = useState(0)
  const [ws, setWs] = useState<WebSocket | null>(null)
  const canvasRef = useRef() as MutableRefObject<HTMLCanvasElement>

  const { register, handleSubmit, formState: { errors } } = useForm<GameStartInputs>({
    mode: "onBlur",
    resolver: yupResolver(schema),
  })

  function startAnimating(fps: number) {
    const fpsInterval = 1000 / fps
    const then = Date.now()
    const startTime = then
    animate(fpsInterval, startTime, then)
  }

  let PLAYERS_DATA: Array<Player> = []

  const movingTimeRate = 0.5
  let movingTime = Date.now()
  let move = false

  function animate(fpsInterval: number, startTime: number, then: number) {
    const canvas = canvasRef.current
    if (!canvas) return
    const ctx = canvas.getContext("2d")
    if (!ctx) return

    const now = Date.now()
    requestAnimationFrame(() => animate(fpsInterval, startTime, then))
    const elapsed = now - then
    if (elapsed < fpsInterval) return
    then = now - (elapsed % fpsInterval)
    ctx.clearRect(0, 0, CANVAS_WIDTH, CANVAS_HEIGHT)
    const background = new Image()
    background.src = `${window.location.origin}/img/assets/backgrounds/mountains.png`
    ctx.drawImage(background, 0, 0, canvas.width, canvas.height)

    if ((now % movingTime) / 1000 >= movingTimeRate) {
      console.log("changing")
      move = !move
      movingTime = Date.now()
    }

    PLAYERS_DATA.forEach((player) => {
      const image = new Image()
      image.src = `${window.location.origin}${getMovingCharacterSprite(
        player["type"], player["last_direction"], move,
      )}`
      ctx.drawImage(
        image, player["position_x"], player["position_y"], CHARACTER_SIZE, CHARACTER_SIZE
      )
      ctx.fillText(player["username"] + `  ⚔️ (${player.wins.length})`, player["position_x"], player["position_y"] - 10)
    })
  }

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

    startAnimating(FPS)

    webSocket.onmessage = (event) => {
      const data = JSON.parse(event.data)
      if (data.type === "broadcast") {
        PLAYERS_DATA = data.players
        setPlayerCount(data.players.length)
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
      {!username ?
        <S.Container maxW="xl" centerContent>
          <section>

            <header>
              <S.Title>{"Let's play the game"}</S.Title>
              <S.Paragraph>Welcome!</S.Paragraph>
            </header>

            <main>
              <FormControl isInvalid={!!errors.username} onSubmit={handleSubmit(onStart)}>
                <form>
                  <Stack spacing={2}>
                    <FormLabel display="none">Username</FormLabel>
                    <FormHelperText>{"Choose your username here:"}</FormHelperText>
                    <Input
                      id="username"
                      placeholder="username"
                      isInvalid={!!errors.username}
                      {...register("username")}
                    />
                    {errors.username && <FormErrorMessage>{errors.username.message}</FormErrorMessage>}

                    <HStack>
                      <Button onClick={handleSubmit(onStart)} size="sm">Start</Button>
                    </HStack>
                  </Stack>
                </form>
              </FormControl>
            </main>

          </section>
        </S.Container > : (
          <>
            <S.GameContainer tabIndex={0} onKeyDown={(e) => handleKeyDown(e.key)}>
            </S.GameContainer>
            <S.Canvas id="rpg" ref={canvasRef}>
            </S.Canvas>
            <span>{playerCount} player{!playerCount || playerCount > 1 ? "s" : ""} online</span>
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
      <S.Credits>
        {"developed by "}
        <a target="_blank" rel="noopener noreferrer" href="https://github.com/Guilehm">guilherme luis</a>
      </S.Credits>
    </>
  )
}


export default RPG
