import { AxiosResponse } from "axios"
import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import Card from "../../components/Card/Card"
import { Spinner } from "../../components/Spinner/Spinner.styles"
import ApiService from "../../services/api-service"

type UrlParams = {
  slug: string
}

type GameDetailData = {
  name: string,
  background_image: string
}

const Api = new ApiService()

const GameDetail = (): JSX.Element => {
  const { slug } = useParams<UrlParams>()
  const [gameData, setGameData] = useState<GameDetailData>()


  useEffect(() => {
    const handleSuccess = (response: AxiosResponse<GameDetailData>) => {
      setGameData(response.data)
    }

    Api.getGameDetail(slug)
      .then(handleSuccess)
  }, [slug])

  return (
    gameData ?
      <Card title={gameData.name} image={gameData.background_image} />
      : <Spinner />
  )
}


export default GameDetail
