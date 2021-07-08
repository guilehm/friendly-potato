import * as S from './Card.styles'


type CardProps = {
  title: string
  image: string
}


const Card: React.FC<CardProps> = ({ title, image }) => {
  return (
    <S.Section>
      <S.Image src={image}></S.Image>
      <S.Title>{title}</S.Title>
    </S.Section>
  )
}

export default Card
