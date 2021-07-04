import * as S from './Button.styles'

type ButtonProps = {
  children: JSX.Element | string
}


const Button: React.FC<ButtonProps> = ({ children }): JSX.Element => (
  <S.Button>{children}</S.Button>
)


export default Button
