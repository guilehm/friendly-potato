import { AxiosError, AxiosResponse } from 'axios'
import { useEffect, useState } from 'react'
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

      const handleSuccess = (response: AxiosResponse) =>
        setGameList(response.data.results)

      const handleError = (error: AxiosError) => console.log(error)

      Api.getGameList()
        .then(handleSuccess)
        .catch(handleError)
    }
    fetchGameList()
  }, [])

  return (
    <HomeContainer>
      {gameList.map(game => <div key={game.slug}>{game.name}</div>)}
    </HomeContainer>
  )
}


export default Home
