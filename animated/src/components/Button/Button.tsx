import * as S from './Button.styles'

type ButtonProps = {
  children: JSX.Element | string
  color?: string
  bgColor?: string
}


const Button: React.FC<ButtonProps> = ({ children, color, bgColor }): JSX.Element => (
  <S.Button color={color} bgColor={bgColor}>{children}</S.Button>
)


export default Button
