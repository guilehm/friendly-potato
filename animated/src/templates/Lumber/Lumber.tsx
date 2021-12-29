import { useEffect } from "react"
import Cookies from "universal-cookie"
import { WSMessage } from "../../types/ws-types"
import * as S from "./Lumber.styles"

const cookies = new Cookies()

const Lumber = (): JSX.Element => {

  useEffect(() => {
    const location = "ws://localhost:8080/socket/"
    const webSocket = new WebSocket(location)

    const message: WSMessage = {
      type: "login",
      data: { "refresh_token": cookies.get("refresh") },
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
      <S.List>
        <S.ListItem>
          <span>1</span>
          <S.Image src={`${window.location.origin}/img/lumber/a.png`} />
        </S.ListItem>
        <S.ListItem>
          <span>1</span>
          <S.Image src={`${window.location.origin}/img/lumber/a.png`} />
        </S.ListItem>
      </S.List>
    </S.LumberContainer>
  )
}


export default Lumber
