import * as S from './Layout.styles'


type LayoutProps = {
  children: JSX.Element
}

const Layout: React.FC<LayoutProps> = ({ children }): JSX.Element => (
  <S.Container>
    {children}
  </S.Container>
)


export default Layout
