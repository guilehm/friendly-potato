import { Container } from "@chakra-ui/react"
import { AxiosResponse } from "axios"
import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import AlertComposition from "../../components/AlertComposition"
import Card from "../../components/Card/Card"
import { Spinner } from "../../components/Spinner/Spinner.styles"
import ApiService from "../../services/api-service"

type UrlParams = {
  slug: string
}

type GameDetailData = {
  name: string,
  slug: string,
  background_image: string
}

const Api = new ApiService()

const GameDetail = (): JSX.Element => {
  const { slug } = useParams<UrlParams>()
  const [gameData, setGameData] = useState<GameDetailData>()
  const [error, setError] = useState(true)


  useEffect(() => {
    const handleSuccess = (response: AxiosResponse<GameDetailData>) => {
      setGameData(response.data)
    }
    const handleError = () => setError(true)

    Api.getGameDetail(slug)
      .then(handleSuccess)
      .catch(handleError)
  }, [slug])

  return (
    gameData ?
      <Container maxW="container.xl">
        <Card
          height={"500px"}
          zoom={false}
          title={gameData.name}
          image={gameData.background_image}
          slug={gameData.slug}
        />
      </Container>
      : error ?
        <AlertComposition
          status={"error"}
          title={"There was an error processing your request"}
          description={"Please try again"} /> : <Spinner />
  )
}


export default GameDetail
