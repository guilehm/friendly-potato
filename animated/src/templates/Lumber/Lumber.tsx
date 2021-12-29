import { useEffect } from "react"
import { WSMessage } from "../../types/ws-types"
import * as S from "./Lumber.styles"


const Lumber = (): JSX.Element => {

  useEffect(() => {
    const location = "ws://localhost:8080/socket/"
    const webSocket = new WebSocket(location)

    const message: WSMessage = {
      type: "login",
      data: { "refresh_token": "whatever" },
    }

    webSocket.onopen = () => {
      webSocket.send(JSON.stringify(message))
    }

    webSocket.onmessage = (event) => {
      console.log("message:", event.data)
    }

  }, [])

  return (
    <S.LumberContainer>
      <h1>Hello from Lumber</h1>
    </S.LumberContainer>
  )
}


export default Lumber
