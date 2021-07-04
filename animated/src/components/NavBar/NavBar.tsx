import { Link } from 'react-router-dom'
import Avatar from '../Avatar'
import * as S from './NavBar.styles'


const NavBar = (): JSX.Element => (
  <S.Nav>
    <S.Container>
      <Link to="/login">
        <Avatar />
      </Link>
    </S.Container>
  </S.Nav>
)

export default NavBar
