import Button from '../../components/Button/Button'
import * as S from './Login.styles'


const Login = (): JSX.Element => (
  <S.Container>
    <S.Title>{'Let\'s sign you in'}</S.Title>
    <S.Paragraph>Welcome back!</S.Paragraph>
    <Button>Sign In</Button>
  </S.Container>
)


export default Login
