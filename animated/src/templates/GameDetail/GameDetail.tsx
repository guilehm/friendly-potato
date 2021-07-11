import { AxiosResponse } from 'axios'
import { useEffect, useState } from 'react'
import Card from '../../components/Card/Card'
import ApiService from '../../services/api-service'


type GameDetailData = {
  name: string,
  background_image: string
}

const Api = new ApiService()

const GameDetail = (): JSX.Element => {
  const [gameData, setGameData] = useState<GameDetailData>()


  useEffect(() => {
    const handleSuccess = (response: AxiosResponse<GameDetailData>) => {
      setGameData(response.data)
    }

    Api.getGameDetail('overkill')
      .then(handleSuccess)
  }, [])

  return (
    gameData ? <Card title={gameData.name} image={gameData.background_image} /> : <></>
  )
}


export default GameDetail
