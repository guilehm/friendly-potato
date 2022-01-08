import { MutableRefObject, useEffect, useRef, useState } from "react"
import { Warrior } from "../../types/rogue-types"
import * as S from "./RogueLike.styles"


const RogueLike = (): JSX.Element => {

  const [ws, setWs] = useState<WebSocket | null>(null)
  const canvasRef = useRef() as MutableRefObject<HTMLCanvasElement>

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
  }, [])


  return (
    <S.Container>
      <S.Canvas width={720} height={480}>
      </S.Canvas>
    </S.Container>
  )
}


export default RogueLike
