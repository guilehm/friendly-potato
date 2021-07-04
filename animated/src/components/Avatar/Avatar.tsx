import * as S from './Avatar.styles'

const Avatar = (): JSX.Element => (
  <S.Image src={`${window.location.origin}/img/default-avatar.png`} />
)

export default Avatar
