import { useEffect, useState } from "react"
import { RouteComponentProps } from "react-router-dom"
import Cookies from "universal-cookie"
import ApiService from "../../services/api-service"
import { WSMessage } from "../../types/ws-types"
import * as S from "./Lumber.styles"

const cookies = new Cookies()
const Api = new ApiService()


const Lumber = ({ history }: RouteComponentProps): JSX.Element => {

  const [lumberCount, setLumberCount] = useState<number>(0)
  const [goldCount, setGoldCount] = useState<number>(0)

  useEffect(() => {
    const cookies = new Cookies(["access", "refresh"])
    const accessToken = cookies.get("access")
    const refreshToken = cookies.get("refresh")

    if (!accessToken || !refreshToken) {
      history.push("/login")
    } else {
      Api.validateToken(refreshToken)
        .catch(() => history.push("/login"))
    }

  }, [history])

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
      const data = JSON.parse(event.data)
      if (data.type === "update") {
        const woods = data["player_data"]["woods"]
        woods && setLumberCount(woods.length)
        const gold = data["player_data"]["gold"]
        gold && setLumberCount(gold.length)
      }
    }

  }, [])

  return (
    <S.LumberContainer>
      <h1>Hello from Lumber</h1>
      <S.List>
        <S.ListItem>
          <S.Number>{lumberCount}</S.Number>
          <S.Image src={`${window.location.origin}/img/lumber/a.png`} />
        </S.ListItem>
        <S.ListItem>
          <S.Number>{goldCount}</S.Number>
          <S.Image src={`${window.location.origin}/img/lumber/gold-chest.png`} />
        </S.ListItem>
      </S.List>
    </S.LumberContainer>
  )
}


export default Lumber
