import * as S from "./Card.styles"
import { Link } from "react-router-dom"

type CardProps = {
  title: string
  image: string
  slug: string
  zoom?: boolean
}


const Card: React.FC<CardProps> = ({ title, image, slug, zoom = true }) => {
  return (
    <S.Section zoom={zoom}>
      <S.Image src={image} ></S.Image>
      <S.Title>
        <Link to={`/games/${slug}/`}>
          {title}
        </Link>
      </S.Title>
    </S.Section>
  )
}

export default Card
export type { CardProps }
