import Button from '../../components/Button/Button'
import * as S from './Login.styles'


const Login = (): JSX.Element => (
  <S.Container>
    <S.Section>
      <header>
        <S.Title>{'Let\'s sign you in'}</S.Title>
        <S.Paragraph>Welcome back!</S.Paragraph>
      </header>
      <main>

      </main>
      <footer>
        <Button>Register</Button>
        <Button>Sign In</Button>
      </footer>
    </S.Section>
  </S.Container>
)


export default Login
