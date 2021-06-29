import { AxiosError, AxiosResponse } from 'axios'
import { useEffect, useState } from 'react'
import Card from '../../components/Card/Card'
import ApiService from '../../services/api-service'
import { HomeContainer } from './Home.styles'


type GameListResult = {
  id: number
  slug: string
  name: string
  background_image: string
}

const Api = new ApiService()

const Home = (): JSX.Element => {
  const [gameList, setGameList] = useState<GameListResult[]>([])

  useEffect(() => {
    const fetchGameList = async () => {

      const handleError = (error: AxiosError) => console.log(error)
      const handleSuccess = (response: AxiosResponse) =>
        setGameList(response.data.results)

      Api.getGameList()
        .then(handleSuccess)
        .catch(handleError)
    }
    fetchGameList()
  }, [])

  return (
    <HomeContainer>
      {gameList.map(game => <Card key={game.slug} title={game.name} />)}
    </HomeContainer>
  )
}


export default Home
