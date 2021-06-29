/* eslint-disable react/prop-types */
import * as S from './Card.styles'


type CardProps = {
  title: string
}


const Card: React.FC<CardProps> = ({ title }) => {
  return (
    <S.Title>{title}</S.Title>
  )
}

export default Card
