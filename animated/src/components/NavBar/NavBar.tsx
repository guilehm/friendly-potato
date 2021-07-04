import Avatar from '../Avatar'
import * as S from './NavBar.styles'


const NavBar = (): JSX.Element => (
  <S.Nav>
    <S.Container>
      <S.SearchButton />
      <Avatar />
    </S.Container>
  </S.Nav>
)

export default NavBar
