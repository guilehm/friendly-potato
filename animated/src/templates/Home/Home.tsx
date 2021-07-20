import { AxiosResponse } from "axios"
import { useEffect, useState } from "react"
import AlertComposition from "../../components/AlertComposition"
import Card from "../../components/Card/Card"
import Spinner from "../../components/Spinner"
import ApiService from "../../services/api-service"
import * as S from "./Home.styles"

type GameListResult = {
  id: number
  slug: string
  name: string
  background_image: string
}

const Api = new ApiService()

const Home = (): JSX.Element => {
  const [error, setError] = useState(false)
  const [gameList, setGameList] = useState<GameListResult[]>([])

  useEffect(() => {
    const fetchGameList = async () => {
      const handleError = () => setError(true)
      const handleSuccess = (response: AxiosResponse) => {
        setGameList(response.data.results)
      }

      Api.getGameList()
        .then(handleSuccess)
        .catch(handleError)
    }
    fetchGameList()
  }, [])

  return (
    <S.HomeContainer>
      <S.HomeSection>
        {gameList.length ? (
          gameList.map((game) => (
            <Card
              key={game.slug}
              title={game.name}
              slug={game.slug}
              image={game.background_image}
            />
          ))
        ) : error ?
          <AlertComposition
            status={"error"}
            title={"There was an error processing your request"}
            description={"Please try again"} /> :
          <Spinner />
        }
      </S.HomeSection>
    </S.HomeContainer>
  )
}


export default Home
