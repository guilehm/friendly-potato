import { AxiosError, AxiosResponse } from 'axios'
import { useEffect, useState } from 'react'
import Card from '../../components/Card/Card'
import NavBar from '../../components/NavBar'
import Spinner from '../../components/Spinner'
import ApiService from '../../services/api-service'
import * as S from './Home.styles'


type GameListResult = {
  id: number
  slug: string
  name: string
  background_image: string
}

const Api = new ApiService()

const Home = (): JSX.Element => {
  const [loading, setLoading] = useState(true)
  const [gameList, setGameList] = useState<GameListResult[]>([])

  useEffect(() => {
    const fetchGameList = async () => {

      const handleError = (error: AxiosError) => console.log(error)
      const handleSuccess = (response: AxiosResponse) => {
        setGameList(response.data.results)
        setLoading(false)
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
        {loading ? <Spinner /> :
          gameList.map(game =>
            <Card
              key={game.slug}
              title={game.name}
              image={game.background_image} />)
        }
      </S.HomeSection>
    </S.HomeContainer>
  )
}


export default Home
